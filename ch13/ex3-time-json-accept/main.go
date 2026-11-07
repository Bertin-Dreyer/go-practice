package main

import "fmt"

// Exercise 13.3:
// Add the ability to return the time as JSON. Use the Accept header to control
// whether JSON or text is returned (default to text).
// JSON structure:
//   {"day_of_week": "Monday", "day_of_month": 10, "month": "April", "year": 2023, "hour": 20, "minute": 15, "second": 20}

// TODO: Write a timeResponse struct with json tags matching the structure above
// TODO: Write handler checking r.Header.Get("Accept") for "application/json"
// TODO: If JSON, encode timeResponse; otherwise write RFC3339 text
// TODO: Start server

func main() {
	// TODO: Implement time server with Accept-header content negotiation
	fmt.Println("TODO: implement ex3-time-json-accept")
}
