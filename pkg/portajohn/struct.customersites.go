package portajohn

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerSite struct {
	Cocode      string
	Custmast    string
	Custnum     int
	Startdate   time.Time
	Taxpcnt     float64
	Taxpcnt2    float64
	Taxexempt   string
	Penalty     float64
	Po_num      string
	D_waiver    string
	Sitename    string
	Siteaddr    string
	Siteaddr2   string
	Sitecity    string
	Sitestate   string
	Sitezip     string
	Sitephone   string
	Sitefax     string
	Sitezip4    string
	Sitecntry   string
	Super       string
	Bllmast     string
	Blladdr     string
	Blladdr2    string
	Bllcity     string
	Bllemail1   sql.NullString
	Bllemail2   sql.NullString
	Bllfax      string
	Bllname     string
	Bllphone    string
	Bllstate    string
	Bllzip      string
	Bllzip4     string
	Bllcocode   string
	Bllcontact  string
	Bllcountry  string
	Acctstatus  string
	Accttype    string
	County      string
	Custcode1   string
	Terms       string
	Email       sql.NullString
	Email2      sql.NullString
	Emaillst    bool
	Nomail      bool
	Nomailing   float64
	FKjcustmast sql.NullInt64
}

func (cs *CustomerSite) Format() {
	cs.Custmast = FormatString(cs.Custmast)
	cs.Sitename = FormatString(cs.Sitename)
	cs.Sitephone = FormatPhoneNumber(cs.Sitephone)
	cs.Sitefax = FormatPhoneNumber(cs.Sitefax)
	cs.Siteaddr = FormatString(cs.Siteaddr)
	cs.Siteaddr2 = FormatString(cs.Siteaddr2)
	cs.Sitecity = FormatString(cs.Sitecity)
	cs.Sitestate = FormatString(cs.Sitestate)
	cs.Sitezip = FormatString(cs.Sitezip)

	// regex match for zip and zip4
	ziprgx := regexp.MustCompile(`(?P<zip>[0-9]{5})(.*)?(?P<zip4>[0-9]{4})?`)
	match := ziprgx.FindStringSubmatch(cs.Sitezip)
	pm := make(map[string]string)
	for i, name := range ziprgx.SubexpNames() {
		if i > 0 && i <= len(match) {
			pm[name] = match[i]
		}
	}
	cs.Sitezip = pm["zip"]
	cs.Sitezip4 = pm["zip4"]

	cs.Sitecntry = FormatString(cs.Sitecntry)
	cs.Po_num = FormatString(cs.Po_num)
	cs.Super = FormatString(cs.Super)
	cs.Terms = FormatString(cs.Terms)

	cs.Bllname = FormatString(cs.Bllname)
	cs.Bllfax = FormatPhoneNumber(cs.Bllfax)
	cs.Bllphone = FormatPhoneNumber(cs.Bllphone)
	cs.Blladdr = FormatString(cs.Blladdr)
	cs.Blladdr2 = FormatString(cs.Blladdr2)
	cs.Bllcity = FormatString(cs.Bllcity)
	cs.Bllstate = FormatString(cs.Bllstate)
	cs.Bllzip = FormatString(cs.Bllzip)

	// regex match for zip and zip4
	match = ziprgx.FindStringSubmatch(cs.Bllzip)
	pm = make(map[string]string)
	for i, name := range ziprgx.SubexpNames() {
		if i > 0 && i <= len(match) {
			pm[name] = match[i]
		}
	}
	cs.Bllzip = pm["zip"]
	cs.Bllzip4 = pm["zip4"]

	cs.Bllcountry = FormatString(cs.Bllcountry)
	cs.Bllcontact = FormatString(cs.Bllcontact)
} // ./Format

func (cs *CustomerSite) Scan(i interface{}) error {
	c2 := CustomerSite{}
	switch v := i.(type) {
	case *sql.Row:
		err := v.Scan(
			c2.Cocode,
			c2.Custmast,
			c2.Custnum,
			c2.Startdate,
			c2.Taxpcnt,
			c2.Taxpcnt2,
			c2.Taxexempt,
			c2.Penalty,
			c2.Po_num,
			c2.D_waiver,
			c2.Sitename,
			c2.Siteaddr,
			c2.Siteaddr2,
			c2.Sitecity,
			c2.Sitestate,
			c2.Sitezip,
			c2.Sitephone,
			c2.Sitefax,
			c2.Sitezip4,
			c2.Sitecntry,
			c2.Super,
			c2.Bllmast,
			c2.Blladdr,
			c2.Blladdr2,
			c2.Bllcity,
			c2.Bllemail1,
			c2.Bllemail2,
			c2.Bllfax,
			c2.Bllname,
			c2.Bllphone,
			c2.Bllstate,
			c2.Bllzip,
			c2.Bllcocode,
			c2.Bllcontact,
			c2.Bllcountry,
			c2.Acctstatus,
			c2.Accttype,
			c2.County,
			c2.Custcode1,
			c2.Terms,
			c2.Email,
			c2.Email2,
			c2.Emaillst,
			c2.Nomail,
			c2.Nomailing,
			c2.FKjcustmast,
		)
		*cs = c2
		cs.Format()
		return err
	case *sql.Rows:
		err := v.Scan(
			&c2.Cocode,
			&c2.Custmast,
			&c2.Custnum,
			&c2.Startdate,
			&c2.Taxpcnt,
			&c2.Taxpcnt2,
			&c2.Taxexempt,
			&c2.Penalty,
			&c2.Po_num,
			&c2.D_waiver,
			&c2.Sitename,
			&c2.Siteaddr,
			&c2.Siteaddr2,
			&c2.Sitecity,
			&c2.Sitestate,
			&c2.Sitezip,
			&c2.Sitephone,
			&c2.Sitefax,
			&c2.Sitezip4,
			&c2.Sitecntry,
			&c2.Super,
			&c2.Bllmast,
			&c2.Blladdr,
			&c2.Blladdr2,
			&c2.Bllcity,
			&c2.Bllemail1,
			&c2.Bllemail2,
			&c2.Bllfax,
			&c2.Bllname,
			&c2.Bllphone,
			&c2.Bllstate,
			&c2.Bllzip,
			&c2.Bllcocode,
			&c2.Bllcontact,
			&c2.Bllcountry,
			&c2.Acctstatus,
			&c2.Accttype,
			&c2.County,
			&c2.Custcode1,
			&c2.Terms,
			&c2.Email,
			&c2.Email2,
			&c2.Emaillst,
			&c2.Nomail,
			&c2.Nomailing,
			&c2.FKjcustmast,
		)
		*cs = c2
		cs.Format()
		return err
	default:
		return fmt.Errorf("*sql.Row or *sql.Rows required")
	}
} // ./Scan

func (cs CustomerSite) Customer() Customer {
	return Customer{
		ID:             primitive.NewObjectID(),
		UID:            NewUID(),
		TacMaster:      cs.Custmast,
		TacMasterId:    int64(cs.Custnum),
		AccountType:    0,
		Status:         1,
		Email:          cs.Bllemail1.String,
		PhoneNumber:    cs.Bllphone,
		Name:           cs.Bllname,
		Zip:            cs.Bllzip,
		Zip4:           cs.Bllzip4,
		SecondaryEmail: cs.Bllemail1.String,
		BillingInfo: BillingInfo{
			Address:  cs.Blladdr,
			Address2: cs.Blladdr2,
			City:     cs.Bllcity,
			Zip:      cs.Bllzip,
			Zip4:     cs.Bllzip4,
		},
		CreatedAt: cs.Startdate,
	}
} // ./Customer

func (cs CustomerSite) Site() Site {
	return Site{
		ID:          primitive.NewObjectID(),
		AccountID:   primitive.NewObjectID(),
		Status:      1,
		Title:       cs.Sitename,
		PhoneNumber: cs.Sitephone,
		Address:     cs.Siteaddr,
		Address2:    cs.Siteaddr2,
		City:        cs.Sitecity,
		State:       cs.Sitestate,
		Zip:         cs.Sitezip,
		Zip4:        cs.Sitezip4,
		TacMaster:   cs.Custmast,
		TacMasterId: int64(cs.Custnum),

		CreatedAt: cs.Startdate,
	}
} // ./Site
