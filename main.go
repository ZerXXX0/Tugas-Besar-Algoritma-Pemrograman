package main

import "fmt"

const NMAX int = 1000

type dataBarang [NMAX]struct {
	nama        string
	tanggal, id int
}

func main() {
	//variable
	var userInput, jumlah int
	var barang dataBarang

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
		fmt.Println("4. Cari barang")
		fmt.Println("0. Exit")
		fmt.Println("============================================")

		//scan input dari user
		fmt.Println("Pilihan anda: ")
		fmt.Scan(&userInput)

		//switch-case untuk menu
		switch {
		case userInput == 1:
			//fmt.Scan(&jumlah)
			tambahBarang(&barang, &jumlah)
			fmt.Println("Barang berhasil ditambahkan", jumlah)
			break
		case userInput == 2:
			var idDelete int
			fmt.Scan(&idDelete)
			hapusBarang(&barang, &jumlah, idDelete)
			fmt.Println("Barang berhasil dihapus")
			break
		case userInput == 3:
			//header
			fmt.Println("              Inventori Anda")
			fmt.Println("============================================")
			fmt.Println("Id    Nama                Tanggal masuk")
			lihatInventori(barang, jumlah)
			break
		case userInput == 0:
			fmt.Println("Terima kasih, sampai jumpa lagi!!")
			fmt.Println("Thank you, see you again!!")
			fmt.Println("Danke, auf Wiedersehen!!")
			fmt.Println("Merci, au revoir!!")
			fmt.Println("Gracias, adiós!!")
			fmt.Println("Grazie, arrivederci!!")
			i = 1
			break
		}
	}
}
func tambahBarang(barang *dataBarang, jumlah *int) {
	//input jumlah
	fmt.Println("Masukkan jumlah barang: ")
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
	//searching index
	var searchedIndex int
	for i := 0; i < *jumlah; i++ {
		if barang[i].id == idDelete {
			searchedIndex = i
		}
	}

	//decrementing the index
	*jumlah -= 1
	//swapping the deleted index
	for searchedIndex < *jumlah {
		barang[searchedIndex].nama = barang[searchedIndex+1].nama
		barang[searchedIndex].tanggal = barang[searchedIndex+1].tanggal
		barang[searchedIndex].id = barang[searchedIndex+1].id
		searchedIndex++
	}
}
func lihatInventori(barang dataBarang, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Inventori masih kosong")
	} else {
		//out of bound error handling
		if jumlah > NMAX {
			jumlah = NMAX
		}

		//inputting data
		for i := 0; i < jumlah; i++ {
			fmt.Print(barang[i].id, "    ")
			fmt.Print(barang[i].nama, "                 ")
			fmt.Print(barang[i].tanggal)
			fmt.Println()
		}
	}
}
