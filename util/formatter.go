package util

import "fmt"

// FormatAmount returns a string formatted currency amount.
// https://api.youneedabudget.com/#formats
func FormatAmount(amount int64, curr string) string {
  val := float64(amount) / 1000.0
  return fmt.Sprintf("%v%.2f", curr, val)
}
