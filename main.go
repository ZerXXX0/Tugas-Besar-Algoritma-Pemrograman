package main

import "fmt"

const NMAX int = 100

type dataBarang [NMAX]struct {
	nama        string
	tanggal, id int
}

func main() {
	//interface
	fmt.Println("============================================")
	fmt.Println("   Welcome to Aplikasi Inventori Barang")
	fmt.Println("============================================")

	i := 0
	for i != 1 {
		//interface
		fmt.Println("============================================")
		fmt.Println("                   MENU")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Lihat Inventori")
		fmt.Println("0. Exit")

		//variable
		var userInput, jumlah int
		var barang dataBarang

		//scan input dari user
		fmt.Scan(&userInput)
		//switch-case untuk menu
		switch {
		case userInput == 1:
			fmt.Scan(&jumlah)
			tambahBarang(&barang, &jumlah)
			fmt.Println("Barang berhasil ditambahkan")
		case userInput == 2:
			var idDelete int
			fmt.Scan(&idDelete)
			hapusBarang(&barang, &jumlah, idDelete)
			fmt.Println("Barang berhasil dihapus")
		case userInput == 3:
			lihatInventori(barang, jumlah)
		case userInput == 0:
			fmt.Println("Terima kasih, sampai jumpa lagi!!")
			fmt.Println("Thank you, see you again!!")
			fmt.Println("Danke, auf Wiedersehen!!")
			fmt.Println("merci, au revoir!!")
			fmt.Println("Gracias, adiós!!")
			fmt.Println("Grazie, arrivederci!!")
			i = 1
		}
	}
}
func tambahBarang(barang *dataBarang, jumlah *int) {

}
func hapusBarang(barang *dataBarang, jumlah *int, idDelete int) {

}
func lihatInventori(barang dataBarang, jumlah int) {

}
