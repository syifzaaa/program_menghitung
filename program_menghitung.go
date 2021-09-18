package main

import (
	"fmt"
	// "strings"
)

func showMenu() {
	fmt.Println("======================")
	fmt.Println("1. Luas segitiga")
	fmt.Println("2. Keliling segitiga")
	fmt.Println("3. Luas persegi")
	fmt.Println("4. Keliling persegi")
	fmt.Println("0. Keluar")
	fmt.Println("======================")
}

func menuOne() {
	var alas, tinggi float64
	fmt.Print("Masukan alas segitiga : ")
	fmt.Scanln(&alas)
	fmt.Print("Masukan tinggi segitiga : ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * alas * tinggi
	fmt.Println("Luas segitiga adalah", luas)
}

func menuTwo() {
	var sisiA, sisiB, sisiC float64
	fmt.Print("Masukan sisi A segitiga : ")
	fmt.Scanln(&sisiA)
	fmt.Print("Masukan sisi B segitiga : ")
	fmt.Scanln(&sisiB)
	fmt.Print("Masukan sisi C segitiga : ")
	fmt.Scanln(&sisiC)

	keliling := sisiA + sisiB + sisiC
	fmt.Println("Keliling segitiga adalah", keliling)
}

func menuThree() {
	var panjang float64
	fmt.Print("Masukan panjang sisi persegi : ")
	fmt.Scanln(&panjang)

	luas := float64(panjang) * float64(panjang)
	fmt.Println("Luas persegi adalah", luas)
}

func menuFour() {
	var panjang float64
	fmt.Print("Masukan panjang sisi persegi : ")
	fmt.Scanln(&panjang)

	keliling := 4 * panjang
	fmt.Println("Keliling persegi adalah", keliling)
}

func getInput() {
	var choice string
	for ok := true; ok; ok = choice != "0" {
		showMenu()
		fmt.Print("Masukan pilihanmu : ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			menuOne()
		case "2":
			menuTwo()
		case "3":
			menuThree()
		case "4":
			menuFour()
		}
	}
}
func main() {
	getInput()
}
