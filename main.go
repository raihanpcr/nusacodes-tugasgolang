package main

import (
	"errors"
	"fmt"
)

func main() {

	jumlah := Add(1,4)
	bagi, err := Bagi(3,1)
	variadic := KurangVariadic(9,8,7,6,5,4,3,2,1)

	// Hasil Penjumlahan
	fmt.Println("Hasil penjumlahan adalah : ", jumlah)

	// Hasil Pembagian
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Hasil pembagian adalah : ", bagi)
	}

	//output variadic
	fmt.Println(variadic)
	
}

func Add(a, b int) int {
	return a+b
}

func Bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak dapat membagi dengan 0")
	}
	return a / b, nil
}

func KurangVariadic(angkas ...int) int {
	
	result := len(angkas)

	for _, angka := range angkas{
		result -= angka
	}

	return result
	

}