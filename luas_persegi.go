
package main
 
import "fmt"
 
func main() {
	var (
		sisi, hasil float64
	)
 
	fmt.Println("Masukan nilai sisi persegi:")
	fmt.Scanf("%f", &sisi)
	hasil = sisi * sisi
	fmt.Println("Luas persegi adalah :", hasil, "m2")
}