package portajohn

import (
	"database/sql"
	"time"
)

type Jinvt struct {
	Serial        string
	Descrip       string
	Custnum       float64
	Inputdate     time.Time
	Billthru      time.Time
	Dailyrate     float64
	Wklyrate      float64
	Extradays     string
	Mthrate       float64
	Delchrg       float64
	Factory       string
	Prevcust      float64
	Delordnum     string
	Status        string
	Srvdrate      float64
	Srvwrate      float64
	Srvmrate      float64
	Damage        float64
	Wdamage       float64
	Mdamage       float64
	Ddisposal     float64
	Disposal      float64
	Mdisposal     float64
	Othdrate      float64
	Othwrate      float64
	Othmrate      float64
	Othdesc       string
	Uentdate      time.Time
	Uenttime      string
	Unitqty       float64
	Utaxrate1     float64
	Utaxrate2     float64
	Specqty       float64
	Spectax       float64
	Specamt       float64
	Specitem      string
	Specmemo      sql.NullString
	Specamttax    float64
	Speccode      string
	Ivtserv       string
	Yardname      string
	Condition     string
	Uload         string
	Spectrate     float64
	Specrate      float64
	Usize         float64
	Umthmin       float64
	Ivtmemo       sql.NullString
	Uinspect      time.Time
	Upaint        time.Time
	Umeterlast    float64
	Umetercur     float64
	R_version     float64
	Salescred     string
	Active        bool
	Rfid          string
	Serial2       string
	Ivtclerk      string
	Serviceday1   bool
	Serviceday2   bool
	Serviceday3   bool
	Serviceday4   bool
	Serviceday5   bool
	Serviceday6   bool
	Serviceday7   bool
	ChargeCode6D  float64
	ChargeCode6W  float64
	ChargeCode6M  float64
	ChargeCode7D  float64
	ChargeCode7W  float64
	ChargeCode7M  float64
	ChargeCode8D  float64
	ChargeCode8W  float64
	ChargeCode8M  float64
	ChargeCode9D  float64
	ChargeCode9W  float64
	ChargeCode9M  float64
	ChargeCode10D float64
	ChargeCode10W float64
	ChargeCode10M float64
	Latitude      string
	Longitude     string
	PriceBook     string
	QRCode        interface{}
	Barcode       string
}

func (j *Jinvt) Format() {
	j.Descrip = FormatString(j.Descrip)
	j.Extradays = FormatString(j.Extradays)
	j.Rfid = FormatString(j.Rfid)
	j.Serial2 = FormatString(j.Serial2)
	j.Ivtclerk = FormatString(j.Ivtclerk)
	j.Speccode = FormatString(j.Speccode)
	j.Ivtserv = FormatString(j.Ivtserv)
	j.Yardname = FormatString(j.Yardname)
	j.Condition = FormatString(j.Condition)
	j.Delordnum = FormatString(j.Delordnum)
	j.Status = FormatString(j.Status)
	j.Serial = FormatString(j.Serial)
	j.Factory = FormatString(j.Factory)
	j.Latitude = FormatString(j.Latitude)
	j.Longitude = FormatString(j.Longitude)
} // ./Format

func (j Jinvt) Scan(i interface{}) error {
	switch v := i.(type) {
	case *sql.Row:
		err := v.Scan(
			&j.Serial,
			&j.Descrip,
			&j.Custnum,
			&j.Inputdate,
			&j.Billthru,
			&j.Dailyrate,
			&j.Wklyrate,
			&j.Extradays,
			&j.Mthrate,
			&j.Delchrg,
			&j.Factory,
			&j.Prevcust,
			&j.Delordnum,
			&j.Status,
			&j.Srvdrate,
			&j.Srvwrate,
			&j.Srvmrate,
			&j.Damage,
			&j.Wdamage,
			&j.Mdamage,
			&j.Ddisposal,
			&j.Disposal,
			&j.Mdisposal,
			&j.Othdrate,
			&j.Othwrate,
			&j.Othmrate,
			&j.Othdesc,
			&j.Uentdate,
			&j.Uenttime,
			&j.Unitqty,
			&j.Utaxrate1,
			&j.Utaxrate2,
			&j.Specqty,
			&j.Spectax,
			&j.Specamt,
			&j.Specitem,
			&j.Specmemo,
			&j.Specamttax,
			&j.Speccode,
			&j.Ivtserv,
			&j.Yardname,
			&j.Condition,
			&j.Uload,
			&j.Spectrate,
			&j.Specrate,
			&j.Usize,
			&j.Umthmin,
			&j.Ivtmemo,
			&j.Uinspect,
			&j.Upaint,
			&j.Umeterlast,
			&j.Umetercur,
			&j.R_version,
			&j.Salescred,
			&j.Active,
			&j.Rfid,
			&j.Serial2,
			&j.Ivtclerk,
			&j.Serviceday1,
			&j.Serviceday2,
			&j.Serviceday3,
			&j.Serviceday4,
			&j.Serviceday5,
			&j.Serviceday6,
			&j.Serviceday7,
			&j.ChargeCode6D,
			&j.ChargeCode6W,
			&j.ChargeCode6M,
			&j.ChargeCode7D,
			&j.ChargeCode7W,
			&j.ChargeCode7M,
			&j.ChargeCode8D,
			&j.ChargeCode8W,
			&j.ChargeCode8M,
			&j.ChargeCode9D,
			&j.ChargeCode9W,
			&j.ChargeCode9M,
			&j.ChargeCode10D,
			&j.ChargeCode10W,
			&j.ChargeCode10M,
			&j.Latitude,
			&j.Longitude,
			&j.PriceBook,
			&j.QRCode,
			&j.Barcode,
		)
		j.Format()
		return err
	case *sql.Rows:
		err := v.Scan(
			&j.Serial,
			&j.Descrip,
			&j.Custnum,
			&j.Inputdate,
			&j.Billthru,
			&j.Dailyrate,
			&j.Wklyrate,
			&j.Extradays,
			&j.Mthrate,
			&j.Delchrg,
			&j.Factory,
			&j.Prevcust,
			&j.Delordnum,
			&j.Status,
			&j.Srvdrate,
			&j.Srvwrate,
			&j.Srvmrate,
			&j.Damage,
			&j.Wdamage,
			&j.Mdamage,
			&j.Ddisposal,
			&j.Disposal,
			&j.Mdisposal,
			&j.Othdrate,
			&j.Othwrate,
			&j.Othmrate,
			&j.Othdesc,
			&j.Uentdate,
			&j.Uenttime,
			&j.Unitqty,
			&j.Utaxrate1,
			&j.Utaxrate2,
			&j.Specqty,
			&j.Spectax,
			&j.Specamt,
			&j.Specitem,
			&j.Specmemo,
			&j.Specamttax,
			&j.Speccode,
			&j.Ivtserv,
			&j.Yardname,
			&j.Condition,
			&j.Uload,
			&j.Spectrate,
			&j.Specrate,
			&j.Usize,
			&j.Umthmin,
			&j.Ivtmemo,
			&j.Uinspect,
			&j.Upaint,
			&j.Umeterlast,
			&j.Umetercur,
			&j.R_version,
			&j.Salescred,
			&j.Active,
			&j.Rfid,
			&j.Serial2,
			&j.Ivtclerk,
			&j.Serviceday1,
			&j.Serviceday2,
			&j.Serviceday3,
			&j.Serviceday4,
			&j.Serviceday5,
			&j.Serviceday6,
			&j.Serviceday7,
			&j.ChargeCode6D,
			&j.ChargeCode6W,
			&j.ChargeCode6M,
			&j.ChargeCode7D,
			&j.ChargeCode7W,
			&j.ChargeCode7M,
			&j.ChargeCode8D,
			&j.ChargeCode8W,
			&j.ChargeCode8M,
			&j.ChargeCode9D,
			&j.ChargeCode9W,
			&j.ChargeCode9M,
			&j.ChargeCode10D,
			&j.ChargeCode10W,
			&j.ChargeCode10M,
			&j.Latitude,
			&j.Longitude,
			&j.PriceBook,
			&j.QRCode,
			&j.Barcode,
		)
		j.Format()
		return err
	}
	return nil
} // ./Scan
