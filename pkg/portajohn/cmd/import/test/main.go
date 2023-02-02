package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"portablejohn.com/pkg/portajohn"
)

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
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	query := `
		SELECT TOP(10)
			cocode,
			custmast,
			custnum,
			startdate,
			taxpcnt,
			taxpcnt2,
			taxexempt,
			penalty,
			po_num,
			d_waiver,
			sitename,
			siteaddr,
			siteaddr2,
			sitecity,
			sitestate,
			sitezip,
			sitephone,    
			sitefax,
			sitezip4,
			sitecntry,
			super,
			
			bllmast,
			blladdr,
			blladdr2,
			bllcity,
			bllemail1,
			bllemail2,    
			bllfax,
			bllname,
			bllphone,
			bllstate,
			bllzip,
			bllcocode,
			bllcontact,    
			bllcountry,
										
			acctstatus,
			accttype,
			county,
			custcode1,
			terms,    
			email,
			email2,
			emaillst,
			nomail,
			nomailing,

			FKjcustmast

		FROM customersites
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var cocode, custmast, taxexempt, po_num,
			d_waiver, sitename, siteaddr, siteaddr2,
			sitecity, sitestate, sitezip, sitephone,
			sitefax, sitezip4, sitecntry, super,
			billmast, billaddr, billaddr2, billcity,
			billemail1, billemail2, billfax,
			billname, billphone, billstate,
			billzip, billcocode, billcontact,
			billcountry, acctstatus, accttype,
			county, custcode1, terms, email, email2 string
		var custnum, taxpcnt, taxpcnt2, penalty float64
		var emaillst, nomail bool
		var fkjcustmast, nomailing int
		var startdate time.Time

		err = rows.Scan(
			&cocode,
			&custmast,
			&custnum,
			&startdate,
			&taxpcnt,
			&taxpcnt2,
			&taxexempt,
			&penalty,
			&po_num,
			&d_waiver,
			&sitename,
			&siteaddr,
			&siteaddr2,
			&sitecity,
			&sitestate,
			&sitezip,
			&sitephone,
			&sitefax,
			&sitezip4,
			&sitecntry,
			&super,
			&billmast,
			&billaddr,
			&billaddr2,
			&billcity,
			&billemail1,
			&billemail2,
			&billfax,
			&billname,
			&billphone,
			&billstate,
			&billzip,
			&billcocode,
			&billcontact,
			&billcountry,
			&acctstatus,
			&accttype,
			&county,
			&custcode1,
			&terms,
			&email,
			&email2,
			&emaillst,
			&nomail,
			&nomailing,
			&fkjcustmast,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(
			custmast,
			po_num,
			sitename,
			emaillst,
			nomail,
			nomailing,
			fkjcustmast,
		)
		sitephone = portajohn.FormatPhoneNumber(sitephone)
		billphone = portajohn.FormatPhoneNumber(billphone)
	}
}
