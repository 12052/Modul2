package main

//muhammad daffa bagus jumantoro(2311102222)
import "fmt"

func main() {
	var base, height float64
	//input nilai alas dan tinggi dari pengguna

	fmt.Print("memasukan nilai alas: ")
	fmt.Scan(&base)

	fmt.Print("memasuakan nilai tinggi: ")
	fmt.Scan(&height)

	//rumus luas segitiga = 1/2*alas*tinggi
	area := 0.5 * base * height

	//menampilkan hasil perhitungan Luas segitiga
	fmt.Printf("luas segitiga adalah: %.2f\n", area)
}
