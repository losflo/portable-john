package main

import (
	"encoding/csv"
	"io"
	"os"

	"portablejohn.com/pkg/portajohn"
)

func main() {
	f, err := os.OpenFile("customercontactlist.csv", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fClean, err := os.OpenFile("customer_contact_list.csv", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fClean.Close()

	w := csv.NewReader(f)
	wc := csv.NewWriter(fClean)

	i := 0
	for {
		/*
			Cust Code: 0
			Company: 1
			Contact: 2
			Phone: 3
			Fax: 4
			Email: 5
		*/
		rows, err := w.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if i == 0 {
			wc.Write(rows)
			wc.Flush()
			i++
			continue
		}
		if i > 1000 {
			break
		}
		wc.Write(
			[]string{
				rows[0],
				rows[1],
				rows[2],
				portajohn.FormatPhoneNumber(rows[3]),
				portajohn.FormatPhoneNumber(rows[4]),
				rows[5],
			},
		)
		wc.Flush()
	}
}
