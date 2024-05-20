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
		fmt.Println("4. ")
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
			//header
			fmt.Println("              Inventori Anda")
			fmt.Println("============================================")
			lihatInventori(barang, jumlah)
		case userInput == 0:
			fmt.Println("Terima kasih, sampai jumpa lagi!!")
			fmt.Println("Thank you, see you again!!")
			fmt.Println("Danke, auf Wiedersehen!!")
			fmt.Println("Merci, au revoir!!")
			fmt.Println("Gracias, adiós!!")
			fmt.Println("Grazie, arrivederci!!")
			i = 1
		}
	}
}
func tambahBarang(barang *dataBarang, jumlah *int) {
	fmt.Scan(jumlah)

	//out of bound error handling
	if *jumlah > NMAX {
		*jumlah = NMAX
	}

	//inputting data
	for i := 0; i < *jumlah; i++ {
		fmt.Println("Masukkan nama barang: ")
		fmt.Scan(&barang[i].nama)
		fmt.Println("Masukkan id barang: ")
		fmt.Scan(&barang[i].id)
		fmt.Println("Masukkan tanggal (MMDD): ")
		fmt.Scan(&barang[i].tanggal)
	}
}
func hapusBarang(barang *dataBarang, jumlah *int, idDelete int) {

	for i := 0; i < *jumlah; i++ {
		barang[idDelete].nama = barang[idDelete+1].nama
		barang[idDelete].tanggal = barang[idDelete+1].tanggal
		barang[idDelete].id = barang[idDelete+1].id
	}
}
func lihatInventori(barang dataBarang, jumlah int) {
	//out of bound error handling
	if jumlah > NMAX {
		jumlah = NMAX
	}

	//inputting data
	for i := 0; i < jumlah; i++ {
		fmt.Print("Nama barang: ", &barang[i].nama, " ")
		fmt.Print("Id barang: ", &barang[i].id, " ")
		fmt.Print("Tanggal masuk: ", &barang[i].tanggal)
		fmt.Println()
	}
}
