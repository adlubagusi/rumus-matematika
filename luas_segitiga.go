package main

import "fmt"
 
func main() {
	var (
		alas, tinggi float64
	)
 
	fmt.Println("Masukan alas :")
	fmt.Scanf("%f", &alas)
	fmt.Println("Masukan tinggi : ")
	fmt.Scanf("%f", &tinggi)
 
	hasil :=  alas * tinggi / 2
	fmt.Println("Hasil = ", hasil)
}