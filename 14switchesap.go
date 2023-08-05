package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("1-toplam \n" +
		"2 - cıkarma \n" +
		"3 bölme  \n" +
		"4 carpma \n " +
		"5 mod alma")

	fmt.Println("lütfen yapmak istediğiniz işlem sırasını giriniz  = ")
	scanner.Scan()
	secim := scanner.Text()

	fmt.Print("1. sayiyi giriniz = ")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("2. sayiyi giriniz = ")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	switch secim {
	case "1":
		fmt.Print("toplama = ", sayi1+sayi2)
	case "2":
		fmt.Println("cıkarma = ", sayi1-sayi2)
	case "3":
		fmt.Println("bölme = ", sayi1/sayi2)
	case "4":
		fmt.Println("çarpma = ", sayi1*sayi2)
	case "5":
		fmt.Println("mod alma = ", int(sayi1)%int(sayi2))
	default:
		fmt.Println("geçersiz işle m")
	}
}
