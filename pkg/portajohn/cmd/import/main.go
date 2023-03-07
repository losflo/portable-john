package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	// _ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"portablejohn.com/pkg/portajohn"
)

var db *sql.DB
var mdb *mongo.Database

func init() {
	var err error
	qry := url.Values{}
	qry.Add("database", "portablejohn")
	username := os.Getenv("SQL_SERVER_USER")
	password := os.Getenv("SQL_SERVER_PASS")
	hostname := os.Getenv("ROSETEC_SERVER")
	port := os.Getenv("SQL_SERVER_PORT")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: qry.Encode(),
	}
	db, err = sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbctx, dbcancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbcancel()
	opts := options.ClientOptions{}
	opts.ApplyURI("mongodb://" + fmt.Sprintf("%s:%s", os.Getenv("ROSETEC_SERVER"), os.Getenv("MONGO_PORT")))
	mngo, err := mongo.Connect(dbctx, &opts)
	if err != nil {
		panic(err)
	}
	err = mngo.Ping(dbctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	mdb = mngo.Database("teal_test")
}

func main() {
	qry := url.Values{}
	qry.Add("database", "portablejohn")
	username := os.Getenv("SQL_SERVER_USER")
	password := os.Getenv("SQL_SERVER_PASS")
	hostname := os.Getenv("ROSETEC_SERVER")
	port := os.Getenv("SQL_SERVER_PORT")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(username, password),
		Host:   fmt.Sprintf("%s:%s", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: qry.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	_, err = db.Begin()
	if err != nil {
		panic(err)
	}

	// import distinct inventory and insert to mongo
	err = importInventory()
	if err != nil {
		panic(err)
	}

	// import customers/sites and inventory linked to customer
	err = importCustomerSites()
	if err != nil {
		panic(err)
	}
}

func importCustomerSites() error {
	query := `
		SELECT TOP(10)
			c1.cocode,
			c1.custmast,
			c1.custnum,
			c1.startdate,
			c1.taxpcnt,
			c1.taxpcnt2,
			c1.taxexempt,
			c1.penalty,
			c1.po_num,
			c1.d_waiver,
			c1.sitename,
			c1.siteaddr,
			c1.siteaddr2,
			c1.sitecity,
			c1.sitestate,
			c1.sitezip,
			c1.sitephone,    
			c1.sitefax,
			c1.sitezip4,
			c1.sitecntry,
			c1.super,
			
			c7.bllmast,
			c7.blladdr,
			c7.blladdr2,
			c7.bllcity,
			c7.bllemail1,
			c7.bllemail2,    
			c7.bllfax,
			c7.bllname,
			c7.bllphone,
			c7.bllstate,
			c7.bllzip,
			c7.bllcocode,
			c7.bllcontact,    
			c7.bllcountry,
										
			c9.acctstatus,
			c9.accttype,
			c9.county,
			c9.custcode1,
			c9.terms,    
			c9.email,
			c9.email2,
			c9.emaillst,
			c9.nomail,
			c9.nomailing,

			c1.FKjcustmast

		FROM PortableJohnData.dbo.jcusf01 as c1
		INNER JOIN PortableJohnData.dbo.jcusf07 as c7 ON c1.FKjcustmast = c7.FKjcustmast 
		INNER JOIN PortableJohnData.dbo.jcusf09 as c9 ON c1.custnum = c9.custnum
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return err
	}

	for rows.Next() {
		cs := portajohn.CustomerSite{}
		if err != nil {
			return err
		}
		err = cs.Scan(rows)
		if err != nil {
			return err
		}

		err = insertCustomerSiteInfo(cs)
		if err != nil {
			return err
		}
	}
	return nil
} // ./importCustomerSites

// importInventory queries for distinct inventory
// cleans description, inserts to mongo with unique id
func importInventory() error {
	query := `
		SELECT DISTINCT descrip 
		FROM PortableJohnData.dbo.jivtf01
		WHERE descrip <> ''
	`
	col := mdb.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	for rows.Next() {
		p := portajohn.Product{
			ID:     primitive.NewObjectID(),
			PID:    fmt.Sprintf("S%s", portajohn.NewUID()),
			Status: 1,
			Images: []string{
				"https://teal-staging.s3.amazonaws.com/files/1000X1000/products/6471678150486.jpg",
			},
			LocationToken: "S5iG65mwWJgSLD0Bt6TXJgSELP5qzx56",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		err := rows.Scan(&p.Title)
		if err != nil {
			return err
		}
		p.Title = portajohn.FormatString(p.Title)
		if p.Title == "" {
			continue
		}
		_, err = col.InsertOne(ctx, p)
		if err != nil {
			return err
		}
	}
	return nil
} // ./importInventory

// importCustomerInventory queries for customer inventory info
// takes custnum, joins customer table with inventory table on custnum
func importCustomerInventory(custnum int, customerId, siteId primitive.ObjectID) error {
	col := mdb.Collection("productlocations")
	query := `
		SELECT * FROM PortableJohnData.dbo.jivtf01 AS inv
		where custnum = @p1 AND descrip <> ''
	`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rows, err := db.QueryContext(ctx, query, custnum)
	if err != nil {
		return err
	}
	for rows.Next() {
		var inv portajohn.Jinvt
		err = inv.Scan(rows)
		if err != nil {
			return err
		}
		p, err := productByTitle(inv.Descrip)
		if err != nil {
			return err
		}
		if p == nil {
			continue
		}
		pl := portajohn.ProductLocation{
			ID:           primitive.NewObjectID(),
			ProductTitle: inv.Descrip,
			ProductID:    p.ID,
			CustomerID:   &customerId,
			SiteID:       &siteId,

			CreatedAt: inv.Inputdate,
			UpdatedAt: inv.Inputdate,
		}
		err = markAccountActive(customerId)
		if err != nil {
			return err
		}
		ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
		_, err = col.InsertOne(ctx2, pl)
		if err != nil {
			cancel2()
			return err
		}
		cancel2()
	}
	return nil
} // ./importCustomerInventory

func markAccountActive(custID primitive.ObjectID) error {
	col := mdb.Collection("useraccounts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": custID}
	update := bson.M{
		"$set": bson.M{
			"status": 1,
		},
	}
	_, err := col.UpdateOne(ctx, filter, update)
	return err
} // ./markAccountActive

// get product by title
func productByTitle(title string) (*portajohn.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"title": title,
	}
	col := mdb.Collection("products")
	var p portajohn.Product
	sr := col.FindOne(ctx, filter)
	if err := sr.Decode(&p); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("product '%s' not found", title)
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
} // ./productByTitle

// insertCustomerSiteInfo
func insertCustomerSiteInfo(cs portajohn.CustomerSite) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c := cs.Customer()
	s := cs.Site()

	col := mdb.Collection("useraccounts")
	colSite := mdb.Collection("sites")

	var cFound portajohn.Customer

	// try search by customer number
	filter := bson.M{
		"tacMaster": c.TacMaster,
	}
	found := true
	sr := col.FindOne(ctx, filter)
	if err := sr.Decode(&cFound); err != nil {
		if err == mongo.ErrNoDocuments {
			found = false
		} else {
			return err
		}
	}

	// no match on cust number
	// try match by billing address
	if !found {
		filter := bson.M{
			"billingInfo.address": c.BillingInfo.Address,
		}
		sr := col.FindOne(ctx, filter)
		if err := sr.Decode(&cFound); err != nil {
			if err == mongo.ErrNoDocuments {
				found = false
			} else {
				return err
			}
		}
	}

	// only insert site
	if found {
		s.AccountID = cFound.ID
		s.TacMaster = cFound.TacMaster
		_, err := colSite.InsertOne(ctx, s)
		if err != nil {
			return err
		}
		err = importCustomerInventory(cs.Custnum, cFound.ID, s.ID)
		return err
	}

	// customer not found
	// insert site and customer
	_, err := col.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	s.AccountID = c.ID
	s.TacMaster = c.TacMaster
	_, err = colSite.InsertOne(ctx, s)
	if err != nil {
		return err
	}

	// check inventory
	err = importCustomerInventory(cs.Custnum, c.ID, s.ID)
	return err
} // ./insertCustomerSiteInfo
