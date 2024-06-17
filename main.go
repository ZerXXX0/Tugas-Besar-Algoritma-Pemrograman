package main

import "fmt"

const NMAX int = 1000

//Data barang struct and array
type dataBarang [NMAX]struct {
	nama              string
	tanggal, id, stok int
}

//Log creation
type log [NMAX]string

//Global variable
var update string
var logActivity log

func main() {
	//variable
	var userInput, jumlah, actualJumlah int
	var barang dataBarang
	//var logActivity log
	//var update string

	//interface
	fmt.Println("============================================")
	fmt.Println("          Lembaga Sembako Present")
	fmt.Println("============================================")
	fmt.Println("   Welcome to Aplikasi Inventori Barang")
	fmt.Println("============================================")

	//Looping interface
	//jika i != 1, maka infinite loop
	//nilai i = 1 ditaruh di pilihan 0.Exit
	i := 0
	for i != 1 {
		//interface
		fmt.Println("============================================")
		fmt.Println("                    MENU")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Lihat Inventori")
		fmt.Println("4. Edit Barang")
		fmt.Println("5. Log Barang")
		fmt.Println("6. Cari Barang")
		fmt.Println("7. Urutkan Barang")
		fmt.Println("0. Exit")
		fmt.Println("============================================")

		//scan input dari user
		fmt.Println("Pilihan anda: ")
		fmt.Scan(&userInput)

		//switch-case untuk menu
		switch {
		case userInput == 1:
			//fmt.Scan(&jumlah)
			tambahBarang(&barang, &jumlah, &actualJumlah)
			fmt.Println("Barang berhasil ditambahkan")
		case userInput == 2:
			var idDelete int
			fmt.Println("Masukkan id barang yang ingin dihapus")
			fmt.Scan(&idDelete)
			hapusBarang(&barang, &actualJumlah, idDelete)
			fmt.Println("Barang berhasil dihapus")
		case userInput == 3:
			//header
			fmt.Println("              Inventori Anda")
			fmt.Println("============================================")
			fmt.Println("Id    Nama             Tanggal masuk    Stok")
			lihatInventori(barang, actualJumlah)
		case userInput == 4:
			var idEdit int
			edit(&barang, &actualJumlah, idEdit)
		case userInput == 5:
			printLog(logActivity)
		case userInput == 6:
			var idToSearch int
			fmt.Println("Masukkan Id Barang yang ingin dicari:")
			fmt.Scan(&idToSearch)
			linearSearch(&barang, idToSearch)
		case userInput == 7:
			var isAsc bool
			fmt.Println("Urutkan berdasarkan stok. Masukkan 1 untuk ascending atau 0 untuk descending:")
			fmt.Scan(&isAsc)
			insertionSort(&barang, &actualJumlah, isAsc)
			fmt.Println("Barang berhasil diurutkan")
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
func tambahBarang(barang *dataBarang, jumlah *int, actualJumlah *int) {
	//input jumlah
	fmt.Println("Masukkan jumlah barang: ")
	fmt.Scan(jumlah)

	//out of bound error handling
	if *jumlah > NMAX {
		*jumlah = NMAX
	}

	*actualJumlah = *jumlah
	for i := 0; i < *jumlah; i++ {
		if barang[i].nama != "" {
			*actualJumlah++
		}
	}

	//inputting data
	for i := 0; i < *actualJumlah; i++ {
		if barang[i].nama == "" {
			fmt.Println("Masukkan nama barang: ")
			fmt.Scan(&barang[i].nama)
			fmt.Println("Masukkan id barang: ")
			fmt.Scan(&barang[i].id)
			fmt.Println("Masukkan tanggal (MMDD): ")
			fmt.Scan(&barang[i].tanggal)
			fmt.Println("Masukkan stok: ")
			fmt.Scan(&barang[i].stok)

			//log
			update = "Tambah barang, banyak barang = " + fmt.Sprintf("%d", *jumlah) + " Id barang: " + fmt.Sprintf("%d", barang[i].id) + " Pada tanggal " + fmt.Sprintf("%d", barang[i].tanggal)
			logBarang(&logActivity, update)
		}
	}
}
func hapusBarang(barang *dataBarang, actualJumlah *int, idDelete int) {
	//searching index
	var searchedIndex int
	for i := 0; i < *actualJumlah; i++ {
		if barang[i].id == idDelete {
			searchedIndex = i
		}
	}

	//decrementing the index
	*actualJumlah -= 1
	//swapping the deleted index
	for searchedIndex < *actualJumlah {
		barang[searchedIndex].nama = barang[searchedIndex+1].nama
		barang[searchedIndex].tanggal = barang[searchedIndex+1].tanggal
		barang[searchedIndex].id = barang[searchedIndex+1].id
		searchedIndex++
	}

	//log
	update = "Hapus barang, Id barang dihapus: " + fmt.Sprintf("%d", barang[idDelete].id) + " Pada tanggal " + fmt.Sprintf("%d", barang[idDelete].tanggal)
	logBarang(&logActivity, update)
}
func lihatInventori(barang dataBarang, actualJumlah int) {
	//error handling for null value
	if actualJumlah == 0 {
		fmt.Println("Inventori masih kosong")
	} else {
		//out of bound error handling
		if actualJumlah > NMAX {
			actualJumlah = NMAX
		}

		//printing data
		for i := 0; i < actualJumlah; i++ {
			fmt.Printf("%-5d %-16s %-16d %-4d\n", barang[i].id, barang[i].nama, barang[i].tanggal, barang[i].stok)
		}
	}
}
func edit(barang *dataBarang, actualJumlah *int, idEdit int) {
	//input id
	fmt.Println("Masukkan id barang: ")
	fmt.Scan(&idEdit)

	//if id sama dengan yang dicari
	var found bool = false
	for i := 0; i < *actualJumlah; i++ {
		if barang[i].id == idEdit {
			found = true
		}
	}

	//if true
	if found == true {
		//searching index
		var searchedIndex int
		for i := 0; i < *actualJumlah; i++ {
			if barang[i].id == idEdit {
				searchedIndex = i
			}
		}
		//option editing
		var option int
		fmt.Println("Pilih edit: ")
		fmt.Println("1. Nama")
		fmt.Println("2. Stok")
		fmt.Println("3. Tanggal")
		fmt.Scan(&option)
		switch {
		case option == 1:
			var changeName string
			fmt.Println("Masukkan nama baru: ")
			fmt.Scan(&changeName)
			barang[searchedIndex].nama = changeName
			fmt.Println("Nama berhasil diubah")
			//log
			update = "Edit nama barang, Id barang diedit: " + fmt.Sprintf("%d", barang[searchedIndex].id)
			logBarang(&logActivity, update)
		case option == 2:
			var changeStok int
			fmt.Println("Masukkan stok baru: ")
			fmt.Scan(&changeStok)
			barang[searchedIndex].stok = changeStok
			fmt.Println("Stok berhasil diubah")
			//log
			update = "Edit stok barang, Id barang diedit: " + fmt.Sprintf("%d", barang[searchedIndex].id)
			logBarang(&logActivity, update)
		case option == 3:
			var changeTanggal string
			fmt.Println("Masukkan tanggal baru: ")
			fmt.Scan(&changeTanggal)
			barang[searchedIndex].nama = changeTanggal
			fmt.Println("Tanggal berhasil diubah")
			//log
			update = "Edit tanggal barang, Id barang diedit: " + fmt.Sprintf("%d", barang[searchedIndex].id)
			logBarang(&logActivity, update)
		}
	} else {
		fmt.Println("Id tidak ditemukan")
	}
}
func logBarang(logActivity *log, update string) {
	for i := 0; i < NMAX; i++ {
		if logActivity[i] == "" {
			logActivity[i] = update
			break
		}
	}
}
func printLog(logActivity log) {
	for i := 0; i < NMAX; i++ {
		if logActivity[i] == "" {
			break
		}
		fmt.Println(logActivity[i])
	}
}
func linearSearch(data *dataBarang, x int) {
	fmt.Println("Masukan Id Barang:")
	isFound := false
	idx := -1
	for i := 0; i < len(data); i++ {
		isFound = data[i].id == x
		if isFound {
			idx = i
			break
		}
	}
	if isFound {
		fmt.Println("Barang berhasil ditemukan. Hasilnya:")
		//fmt.Println(data[idx])
		fmt.Println("============================================")
		fmt.Println("Id    Nama             Tanggal masuk    Stok")
		fmt.Printf("%-5d %-16s %-16d %-4d\n", data[idx].id, data[idx].nama, data[idx].tanggal, data[idx].stok)
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}

func insertionSort(barang *dataBarang, actualJumlah *int, isAsc bool) {
	for i := 0; i < *actualJumlah; i++ {
		if isAsc {
			for j := i; j > 0 && barang[j-1].stok > barang[j].stok; j-- {
				barang[j], barang[j-1] = barang[j-1], barang[j]
			}
		} else {
			for j := i; j > 0 && barang[j-1].stok < barang[j].stok; j-- {
				barang[j], barang[j-1] = barang[j-1], barang[j]
			}
		}
	}
	fmt.Println("              Inventori Anda")
	fmt.Println("============================================")
	fmt.Println("Id    Nama             Tanggal masuk    Stok")
	lihatInventori(*barang, *actualJumlah)
}
