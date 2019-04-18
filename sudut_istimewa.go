package main

import (
"fmt"

"math"
)

var a float64

func main() {
	fmt.Println("masukkan nilai : ")
	fmt.Scanf("%f",&a)
	sinus := math.Sin(a)
	cosinus :=math.Cos(a)
	tangen:=math.Tan(a)
	fmt.Println("nilai sinusnya adalah : ", sinus)
	fmt.Println("nilai cosinus adalah : ", cosinus)
	fmt.Println("nilai tangen adalah : ", tangen)
}