package main

import (
	"fmt"
)

var deger int

func main() {
	const PI float32 = 3.142
	fmt.Printf("PI degeri: %v\n", PI)
	fmt.Printf("PI tipi: %T\n", PI)

	deger = 42
	fmt.Printf("deger: %v", deger)
}
