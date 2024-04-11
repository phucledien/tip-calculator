package total_amount

import (
  "fmt"
	"net/http"
	"strconv"

	"github.com/leapkit/core/render"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  amount, err := strconv.ParseFloat(r.Form.Get("Amount"), 64)
  if err != nil {
    http.Error(w, "Invalid Amount", http.StatusBadRequest)
    return
  }

  tipPercentage, err := strconv.ParseFloat(r.Form.Get("TipPercentage"), 64)
  if err != nil {
    http.Error(w, "Invalid Tip Percentage", http.StatusBadRequest)
    return
  }

  tip := amount * (tipPercentage / 100)
  total := amount + tip

  rw := render.FromCtx(r.Context())
  rw.Set("totalBillAmount", fmt.Sprintf("$%.2f", total))

  err = rw.RenderClean("total_amount/total_amount.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
