package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pahsantana/dio-expert-session-finance/model/transaction"
	"github.com/pahsantana/dio-expert-session-finance/util"
)

func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	// var layout = "2006-01-02T15:04:05"
	// salaryReceived, _ := time.Parse(layout, "2021-03-16T01:57:00")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2021-03-16T01:57:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}
