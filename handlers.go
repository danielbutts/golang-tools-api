package toolexchange

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ToolIndex(w http.ResponseWriter, r *http.Request) {
	newTime, e := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0800 PST 2017")
	fmt.Println(e)

	tools := Tools{
		Tool{
			Id:         1,
			Name:       "hammer",
			ImageURL:   "http://lghttp.18445.nexcesscdn.net/808F9E/mage/media/catalog/product/cache/1/thumbnail/550x/9df78eab33525d08d6e5fb8d27136e95/v/5/v508_hammer.jpg",
			IsBorrowed: true,
			BorrowedOn: newTime,
		},
		Tool{
			Id:         2,
			Name:       "hand saw",
			ImageURL:   "https://db1736767dbd5e7094bb-d61bbc5d0b342a54145a236e2d5d1ebf.ssl.cf4.rackcdn.com/Product-800x800/3588e25d-8421-4c7f-8402-1ecef1daf256.jpg",
			IsBorrowed: false,
			BorrowedOn: time.Time{},
		},
		Tool{
			Id:         3,
			Name:       "hammer",
			ImageURL:   "http://lghttp.18445.nexcesscdn.net/808F9E/mage/media/catalog/product/cache/1/thumbnail/550x/9df78eab33525d08d6e5fb8d27136e95/v/5/v508_hammer.jpg",
			IsBorrowed: true,
			BorrowedOn: newTime,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func ToolShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toolId := vars["toolId"]
	fmt.Fprintln(w, "Tool show:", toolId)
}
