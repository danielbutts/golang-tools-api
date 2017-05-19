package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/danielbutts/toolexchange"
	_ "github.com/lib/pq"
)

const (
	DB_NAME = "tools-rest-api-dev"
)

func main() {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("Inserting Values")
	newTime, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0800 PST 2017")

	tools := toolexchange.Tools{
		toolexchange.Tool{
			Id:         1,
			Name:       "hammer",
			ImageURL:   "http://lghttp.18445.nexcesscdn.net/808F9E/mage/media/catalog/product/cache/1/thumbnail/550x/9df78eab33525d08d6e5fb8d27136e95/v/5/v508_hammer.jpg",
			IsBorrowed: true,
			BorrowedOn: newTime,
		},
		toolexchange.Tool{
			Id:         2,
			Name:       "hand saw",
			ImageURL:   "https://db1736767dbd5e7094bb-d61bbc5d0b342a54145a236e2d5d1ebf.ssl.cf4.rackcdn.com/Product-800x800/3588e25d-8421-4c7f-8402-1ecef1daf256.jpg",
			IsBorrowed: false,
			BorrowedOn: time.Time{},
		},
		toolexchange.Tool{
			Id:         3,
			Name:       "screwdriver",
			ImageURL:   "http://c.shld.net/rpx/i/s/i/spin/image/spin_prod_1198400812?hei=1000&wid=1000&op_sharpen=1",
			IsBorrowed: true,
			BorrowedOn: newTime,
		},
	}

	db.QueryRow("DELETE FROM tools;")

	for _, el := range tools {
		db.QueryRow("INSERT INTO tools(name, image_url, is_borrowed, borrowed_on) VALUES($1, $2, $3, $4) returning id;", el.Name, el.ImageURL, el.IsBorrowed, el.BorrowedOn)
	}
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
