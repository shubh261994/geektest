package orders

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"geektest/app/common/services/postgre"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query()["limit"]
	offsetParam := r.URL.Query()["start"]
	limit := 5
	offset := 0
	if len(limitParam) > 0 {
		limit, _ = strconv.Atoi(limitParam[0])
	}

	if len(offsetParam) > 0 {
		offset, _ = strconv.Atoi(offsetParam[0])
	}
	fmt.Println(limit, offset)
	orders := postgre.GetOrdersList(offset, limit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}