package main

import (
	"fmt"
	"unsafe"
)

// Exercise 16.2:
// Use unsafe.Sizeof and unsafe.Offsetof to print size and offsets for
// OrderInfo struct. Create SmallOrderInfo with same fields reordered for
// minimal memory usage.

// TODO: Define OrderInfo struct with mixed field types (bool, int64, string, etc.)
// TODO: Print unsafe.Sizeof(OrderInfo{}) and unsafe.Offsetof for each field
// TODO: Define SmallOrderInfo with fields reordered to minimize padding
// TODO: Print Sizeof and Offsetof for SmallOrderInfo, compare

type OrderInfo struct {
	// TODO: Define fields here
}

type SmallOrderInfo struct {
	// TODO: Define fields reordered here
}

func main() {
	fmt.Println("OrderInfo size:", unsafe.Sizeof(OrderInfo{}))
	fmt.Println("SmallOrderInfo size:", unsafe.Sizeof(SmallOrderInfo{}))
	fmt.Println("TODO: implement ex2-sizeof-offsetof fully")
}
