package actuator

import (
	"encoding/json"
	"net/http"
)

type HealthBody struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	// var layout = "2006-01-02T15:04:05"
	// salaryReceived, _ := time.Parse(layout, "2021-03-16T01:57:00")

	var status = HealthBody{"alive"}

	// transaction.Transactions{
	// 	transaction.Transaction{
	// 		Title:     "Salário",
	// 		Amount:    1200.0,
	// 		Type:      0,
	// 		CreatedAt: salaryReceived,
	// 	},
	// }

	_ = json.NewEncoder(w).Encode(status)

	// var var = ""
	// var2 := ""
	// var var3 string
}
