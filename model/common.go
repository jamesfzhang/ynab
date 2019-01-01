package model

import (
  "fmt"
  "net/http"
)

type DateFormat struct {
  Format string `json:"format"`
}

type CurrencyFormat struct {
  IsoCode       string `json:"iso_code"`
  ExampleFormat string `json:"example_format"`

  DecimalDigits    int    `json:"decimal_digits"`
  DecimalSeparator string `json:"decimal_separator"`
  GroupSeparator   string `json:"group_separator"`

  CurrencySymbol string `json:"currency_symbol"`
  DisplaySymbol  bool   `json:"display_symbol"`
  SymbolFirst    bool   `json:"symbol_first"`
}

type ApiError struct {
  Response *http.Response
  Detail   ErrorDetail `json:"error"`
}

func (err ApiError) Error() string {
  return fmt.Sprintf("Error (status %v)\n%v\n%+v",
    err.Response.StatusCode,
    err.Response.Request.URL,
    err.Detail,
  )
}

type ErrorDetail struct {
  Id     string `json:"id"`
  Detail string `json:"detail"`
  Name   string `json:"name"`
}
