package main

import (
	"fmt"
	"math"
)

func deret(awal float64, i float64, jumlah float64) float64 {
	total := 0.0

	if i < jumlah {
		bagi := awal / i
		total += bagi
		fmt.Printf("%.2f + ", total)
		return bagi + deret(awal, i+1, jumlah)
	} else {
		bagi := awal / i
		total += bagi
		fmt.Printf("%.2f", bagi)
		return bagi
	}
}

func deretPangkat(i int, jumlah int, pecahan string) string {
	if i <= jumlah {
		pangkat := int(math.Pow(float64(i), float64(2)))
		if pecahan == "pembilang" {
			fmt.Printf("%dx/%d + ", pangkat, i)
		} else {
			fmt.Printf("%d/%dx + ", i, pangkat)
		}
		return deretPangkat(i+1, jumlah, pecahan)
	} else {
		return " "
	}
}

func main() {
	var jumlah int

	fmt.Print("Masukan jumlah : ")
	fmt.Scan(&jumlah)

	var jml float64 = float64(jumlah)

	fmt.Print("a. ")
	fmt.Println(" =", deret(20, 1, jml))
	fmt.Print("b. ")
	fmt.Println(" =", deret(100, 1, jml))
	fmt.Print("c. ")
	fmt.Println(" =", deret(1, 1, jml))
	fmt.Print("d. ")
	fmt.Println("=", deretPangkat(1, jumlah, "penyebut"))
	fmt.Print("e. ")
	fmt.Println("=", deretPangkat(1, jumlah, "pembilang"))

}
