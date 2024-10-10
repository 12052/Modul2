*2
package main
import "fmt"
func main(){
	var angka int
	total:=0
	fmt.Println("masukan angka postif untuk dijumlahkan.masukan angka negatif untuk keluar. ")
	//simulasi while loop
	for{
		fmt.Print("masukan angka : ")
		fmt.Scanln(&angka)

		if angka<0{////kelaur dari loop jika angka negatif}
			break
	}
	total += angka ////tambahkan angka ke total
	fmt.Printf("total jumlah dari angka yang dimasukan adalah: %d\n",total)
}