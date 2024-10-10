// //*1
// package main
// import "fmt"

// func main(){
// 	var a, b, c float64
// 	var hipotenusa bool

// 	fmt.Scanln(&a, &b, &c)
// 	hipotenusa = (c*c) == (a*a + b*b)
// 	fmt.Println( "sisi c adalah  hipotenusa segitiga a,b,c: ", hipotenusa)
// }

// *2
// package main
// import "fmt"
// func main(){
// 	var angka int
// 	total:=0
// 	fmt.Println("masukan angka postif untuk dijumlahkan.masukan angka negatif untuk keluar. ")
// 	//simulasi while loop
// 	for{
// 		fmt.Print("masukan angka : ")
// 		fmt.Scanln(&angka)

// 		if angka<0{////kelaur dari loop jika angka negatif}
// 			break
// 	}
// 	total += angka ////tambahkan angka ke total
// 	fmt.Printf("total jumlah dari angka yang dimasukan adalah: %d\n",total)
// }

// *3
// package main

// import "fmt"

// func main() {

// 	var angka int
// 	total:=0
// 	fmt.Println("Masukkan angka positif untuk di jumlahkan. Masukkan angka negatif untuk keluar. ")

// 	//simulasi while loop
// 	for {
// 		fmt.Print("Masukkan angka: ")
// 		fmt.Scanln(&angka)
// 		if angka < 0 { //keluar dari loop jika angka negatif
// 			break
// 		}
// 		total += angka //tambahkan angka ke total
// 	}

// 	fmt.Printf("Total jumlah dari angka yang dimasukkan adalah : %d\n", total)
// }

// *4
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	maxNumber := 10 //batas maksimum
// 	number := 1     //angka awal
// 	//simulasi repeat-until
// 	for finished := false; !finished; {
// 		fmt.Println("Angka:", number) //cetak angka saat ini
// 		//cek kondisi keluar
// 		if number >= maxNumber {
// 			finished = true //ubah kondisi untuk keluar dari loop
// 		}
// 		number++ //increment angka
// 	}
// }


