// package main

// import "fmt"

// func pembagianLoop(n, m int) int {
// 	hasil := 0
// 	for n >= m {
// 		n -= m
// 		hasil++
// 	}
// 	return hasil
// }

//	func main() {
//		var n, m int
//		fmt.Scan(&n, &m)
//		fmt.Println(pembagianLoop(n, m))
//	}
package main

import "fmt"

func faktor(b1, b2 int) bool {
	/* fungsi mengembalikan nilai boolean true, jika bilangan a
	   adalah faktor dari bilangan b*/
	var result bool

	if (b1 % 2) != 0 {
		result = true
	} else {
		result = false
	}

	return result
}

func faktor(b1 int, b2 int) bool {
	// Mengembalikan false jika a adalah 0 (0 tidak dapat menjadi faktor)
	if b1 == 0 {
		return false
	}

	return b2%b1 == 0
}
func main() {
	var b1, b2 int
	var jenis bool
	fmt.Scan(&b1, &b2)
	jenis = faktor(b1, b2)
	fmt.Println(jenis)
}
