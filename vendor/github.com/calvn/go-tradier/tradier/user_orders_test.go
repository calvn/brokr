package tradier

import (
	"net/http"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestActivityService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/orders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(userOrdersJSON)
	})

	got, _, err := client.User.Orders()
	if err != nil {
		t.Errorf("User.Orders returned error: %v", err)
	}

	if diff := pretty.Compare(wantUserOrders, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

var userOrdersJSON = []byte(`{
  "accounts": {
    "account": [
      {
        "account_number": "6YA05991",
        "orders": {
          "order": {
            "id": 182042,
            "type": "market",
            "symbol": "GOOGL",
            "side": "buy",
            "quantity": 1,
            "status": "pending",
            "duration": "gtc",
            "avg_fill_price": 0,
            "exec_quantity": 0,
            "last_fill_price": 0,
            "last_fill_quantity": 0,
            "remaining_quantity": 1,
            "create_date": "2016-08-23T05:17:37.617Z",
            "transaction_date": "2016-08-23T12:15:07.268Z",
            "class": "equity"
          }
        }
      },
      {
        "account_number": "6YA05708",
        "orders": "null"
      }
    ]
  }
}`)

var (
	createdDate    = time.Date(2016, 8, 23, 05, 17, 37, 617000000, time.UTC)
	transitionDate = time.Date(2016, 8, 23, 12, 15, 07, 268000000, time.UTC)
)

var wantUserOrders = &User{
	Accounts: &Accounts{
		{
			AccountNumber: String("6YA05991"),
			Orders: &Orders{
				{
					ID:                Int(182042),
					Type:              String("market"),
					Symbol:            String("GOOGL"),
					Side:              String("buy"),
					Quantity:          Float64(1),
					Status:            String("pending"),
					Duration:          String("gtc"),
					AvgFillPrice:      Float64(0),
					ExecQuantity:      Float64(0),
					LastFillPrice:     Float64(0),
					LastFillQuantity:  Float64(0),
					RemainingQuantity: Float64(1),
					CreateDate:        &createdDate,
					TransactionDate:   &transitionDate,
					Class:             String("equity"),
				},
			},
		},
		{
			AccountNumber: String("6YA05708"),
			Orders:        &Orders{},
		},
	},
}
