package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lowerCase(kata string, i int) int {
	lower := strings.ToLower(kata)

	if i < len(kata) {
		var filter bool
		if string(kata[i]) == string(lower[i]) {
			filter = true
		}
		if filter == true {
			return 1 + lowerCase(kata, i+1)
		} else {
			return 0 + lowerCase(kata, i+1)
		}
	} else {
		return 0
	}
}

func hitungAngka(kata string, i int) int {
	if i < len(kata) {
		_, err := strconv.Atoi(string(kata[i]))
		if err == nil {
			return 1 + hitungAngka(kata, i+1)
		} else {
			return 0 + hitungAngka(kata, i+1)
		}
	} else {
		return 0
	}
}

func main() {
	var kata string
	var menu int

	fmt.Println("============================")
	fmt.Println("1. Hitung jumlah lower case \n2. Hitung jumlah angka")
	fmt.Println("============================")
	fmt.Print("Pilih menu [1/2] : ")
	fmt.Scan(&menu)
	fmt.Print("Masukan kata : ")
	fmt.Scan(&kata)

	if menu == 1 {
		fmt.Printf("Jumlah huruf non kapital pada kata %s : ", kata)
		fmt.Println(lowerCase(kata, 0))
	} else {
		fmt.Printf("Jumlah angka pada kata %s : ", kata)
		fmt.Println(hitungAngka(kata, 0))
	}

}
