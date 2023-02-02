package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	// _ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
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
	_, err = db.Begin()
	if err != nil {
		panic(err)
	}

	query := `
		SELECT 
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

		FROM jcusf01 as c1

		INNER JOIN jcusf07 as c7 ON c1.FKjcustmast = c7.FKjcustmast 

		INNER JOIN jcusf09 as c9 ON c1.custnum = c9.custnum
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
		var custnum, taxpcnt, taxpcnt2, penalty, nomailing, fkjcustmast int
		var emaillst, nomail bool
		var startdate time.Time

		rows.Scan(
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
			billmast,
			billaddr,
			billaddr2,
			billcity,
			billemail1,
			billemail2,
			billfax,
			billname,
			billphone,
			billstate,
			billzip,
			billcocode,
			billcontact,
			billcountry,
			acctstatus,
			accttype,
			county,
			custcode1,
			terms,
			email, email2,
			emaillst,
			nomail,
			nomailing,
			fkjcustmast,
		)
		sitephone = portajohn.FormatPhoneNumber(sitephone)
		billphone = portajohn.FormatPhoneNumber(billphone)
	}
}