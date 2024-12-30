package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// kapasitas maksimum data
const MaxData = 100

// tipe data TempatWisata berisi Nama, Fasilitas, Jarak, dan Biaya
type TempatWisata struct {
	Nama      string
	Fasilitas string
	Jarak     int
	Wahana    string
	Biaya     int
}

// daftar tempat wisata dengan kapasitas maksimum 100
var daftarTempat [MaxData]TempatWisata

// menyimpan jumlah tempat wisata
var jumlahData int

// fungsi inisialisasi utk menambahkan data awal
func init() {
	dataAwal := []TempatWisata{
		{"Baturraden", "Tempat Parkir, Toilet, Warung Makan, Penginapan", 15, "Kolam Renang", 20000},
		{"Curug Cipendok", "Tempat Parkir, Toilet, Warung Makan", 25, "Trekking", 10000},
		{"Telaga Sunyi", "Tempat Parkir, Toilet", 17, "Perahu", 15000},
		{"Limpak Kuwus", "Tempat Parkir, Toilet, Warung Makan", 20, "Jembatan Gantung", 20000},
		{"Taman Balai Kemambang", "Tempat Parkir, Toilet", 5, "Permainan Anak", 10000},
	}

	// menyalin data awal ke dalam array daftarTempat
	for _, tempat := range dataAwal {
		daftarTempat[jumlahData] = tempat
		jumlahData++
	}
}

// fungsi menambahkan tempat wisata baru ke dalam array daftarTempat menggunakan metode sequential search
func tambahTempat(nama, fasilitas string, jarak int, wahana string, biaya int) {
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(daftarTempat[i].Nama) == strings.ToLower(nama) {
			fmt.Println("Tempat dengan nama tersebut sudah ada.")
			return
		}
	}
	if jumlahData < MaxData {
		daftarTempat[jumlahData] = TempatWisata{nama, fasilitas, jarak, wahana, biaya}
		jumlahData++
		fmt.Println("Data telah ditambah.")
	} else {
		fmt.Println("Data telah penuh, tidak dapat menambahkan data tempat baru.")
	}
}

// fungsi mengubah tempat wisata menggunakan metode sequential search
func ubahTempat(namaAsli string, tempatBaru TempatWisata) {
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(daftarTempat[i].Nama) == strings.ToLower(namaAsli) {
			daftarTempat[i] = tempatBaru
			fmt.Println("Data telah diubah.")
			return
		}
	}
	fmt.Println("Tempat tidak ditemukan.")
}

// fungsi menghapus tempat wisata menggunakan metode sequential search
func hapusTempat(nama string) {
	for i := 0; i < jumlahData; i++ {
		if strings.ToLower(daftarTempat[i].Nama) == strings.ToLower(nama) {
			for j := i; j < jumlahData-1; j++ {
				daftarTempat[j] = daftarTempat[j+1]
			}
			jumlahData--
			fmt.Println("Data telah dihapus.")
			return
		}
	}
	fmt.Println("Tempat tidak ditemukan.")
}

// Fungsi Selection Sort
func selectionSort(data []TempatWisata, banding func(a, b TempatWisata) bool) {
	for i := 0; i < len(data)-1; i++ {
		indeksMin := i
		for j := i + 1; j < len(data); j++ {
			if banding(data[j], data[indeksMin]) {
				indeksMin = j
			}
		}
		data[i], data[indeksMin] = data[indeksMin], data[i]
	}
}

// Urutkan Jarak menggunakan Selection Sort
func urutkanJarak(urut string) {
	// Untuk menampilkan data sebelum diurutkan
	fmt.Println("Data sebelum diurutkan berdasarkan jarak:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}

	// Mengurutkan data dengan membandingkan jarak
	selectionSort(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
		if urut == "ascending" {
			return a.Jarak < b.Jarak // Ascending
		}
		return a.Jarak > b.Jarak // Descending
	})

	// Untuk menampilkan data setelah diurutkan
	fmt.Println("\nData setelah diurutkan berdasarkan jarak:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}
}

// Fungsi Insertion Sort
func insertionSort(data []TempatWisata, banding func(a, b TempatWisata) bool) {
	for i := 1; i < len(data); i++ {
		kunci := data[i]
		j := i - 1
		for j >= 0 && banding(kunci, data[j]) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = kunci
	}
}

// Urutkan Biaya menggunakan Insertion Sort
func urutkanBiaya(urut string) {
	// Untuk menampilkan data sebelum diurutkan
	fmt.Println("Data sebelum diurutkan berdasarkan biaya:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}

	// Mengurutkan data dengan membandingkan biaya
	insertionSort(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
		if urut == "ascending" {
			return a.Biaya < b.Biaya // Ascending
		}
		return a.Biaya > b.Biaya // Descending
	})

	// Untuk menampilkan data setelah diurutkan
	fmt.Println("\nData setelah diurutkan berdasarkan biaya:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}
}

// Urutkan Fasilitas menggunakan Selection Sort
func urutkanFasilitas(urut string) {
	// Untuk menampilkan data sebelum diurutkan
	fmt.Println("Data sebelum diurutkan berdasarkan fasilitas:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}

	// Mengurutkan data dengan membandingkan banyaknya jumlah fasilitas
	selectionSort(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
		countA := len(strings.Split(a.Fasilitas, ","))
		countB := len(strings.Split(b.Fasilitas, ","))
		if urut == "ascending" {
			return countA < countB // Ascending
		}
		return countA > countB // Descending
	})

	// Untuk menampilkan data setelah diurutkan
	fmt.Println("\nData setelah diurutkan berdasarkan fasilitas:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s - Wahana: %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas, daftarTempat[i].Wahana)
	}
}

func isSorted(data []TempatWisata) bool {
	for i := 1; i < len(data); i++ {
		if strings.ToLower(data[i-1].Nama) > strings.ToLower(data[i].Nama) {
			return false
		}
	}
	return true
}

// Fungsi pencarian binary
func cariBinarySearch(kataKunci string) {
	// Pastikan data sudah terurut, jika belum, urutkan sekali
	if !isSorted(daftarTempat[:jumlahData]) {
		// Menggunakan sorting bawaan Go, yang menggunakan algoritma QuickSort secara internal
		sort.SliceStable(daftarTempat[:jumlahData], func(i, j int) bool {
			return strings.ToLower(daftarTempat[i].Nama) < strings.ToLower(daftarTempat[j].Nama)
		})
	}

	// Binary search
	rendah, tinggi := 0, jumlahData-1
	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2
		if strings.ToLower(daftarTempat[tengah].Nama) == strings.ToLower(kataKunci) {
			// Menampilkan hasil pencarian
			fmt.Printf("%s (%d km, Rp%d) - %s - Wahana: %s\n", daftarTempat[tengah].Nama, daftarTempat[tengah].Jarak, daftarTempat[tengah].Biaya, daftarTempat[tengah].Fasilitas, daftarTempat[tengah].Wahana)
			return
		} else if strings.ToLower(daftarTempat[tengah].Nama) < strings.ToLower(kataKunci) {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}
	fmt.Println("Tempat wisata tidak ditemukan")
}

// Fungsi pencarian berdasarkan nama dengan memanggil fungsi cariBinarySearch (binary search)
func cariNama(kataKunci string) {
	cariBinarySearch(kataKunci)
}

// Fungsi pencarian berdasarkan jarak metode sequential search
func cariJarak(kataKunci string) {
	jarak, _ := strconv.Atoi(kataKunci)
	found := false
	for _, tempat := range daftarTempat[:jumlahData] {
		if tempat.Jarak == jarak {
			fmt.Printf("%s (%d km, Rp%d) - Fasilitas: %s - Wahana: %s\n", tempat.Nama, tempat.Jarak, tempat.Biaya, tempat.Fasilitas, tempat.Wahana)
			found = true
		}
	}
	if !found {
		fmt.Println("Tempat wisata tidak ditemukan")
	}
}

// Fungsi pencarian berdasarkan biaya metode sequential search
func cariBiaya(kataKunci string) {
	biaya, _ := strconv.Atoi(kataKunci)
	found := false
	for _, tempat := range daftarTempat[:jumlahData] {
		if tempat.Biaya == biaya {
			fmt.Printf("%s (%d km, Rp%d) - Fasilitas: %s - Wahana: %s\n", tempat.Nama, tempat.Jarak, tempat.Biaya, tempat.Fasilitas, tempat.Wahana)
			found = true
		}
	}
	if !found {
		fmt.Println("Tempat wisata tidak ditemukan")
	}
}

// Fungsi pencarian berdasarkan fasilitas metode sequential search
func cariFasilitas(kataKunci string) {
	found := false
	for _, tempat := range daftarTempat[:jumlahData] {
		if strings.Contains(strings.ToLower(tempat.Fasilitas), strings.ToLower(kataKunci)) {
			fmt.Printf("%s (%d km, Rp%d) - Fasilitas: %s - Wahana: %s\n", tempat.Nama, tempat.Jarak, tempat.Biaya, tempat.Fasilitas, tempat.Wahana)
			found = true
		}
	}
	if !found {
		fmt.Println("Tempat wisata tidak ditemukan")
	}
}

// Fungsi menu login
func menuLogin() int {
	var peran int
	reader := bufio.NewReader(os.Stdin) // Inisialisasi reader

	// menu login
	for {
		fmt.Println("\n--- Menu Login ---")
		fmt.Println("1. Admin")
		fmt.Println("2. Pengguna")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1-3) : ")
		peranStr, _ := reader.ReadString('\n')
		peranStr = strings.TrimSpace(peranStr)
		peran, _ = strconv.Atoi(peranStr)

		switch peran {
		case 1, 2:
			return peran // Kembalikan peran yang dipilih
		case 3:
			fmt.Println("Thank you telah menggunakan Aplikasi Pariwisata.")
			return 0 // Kembali ke menu utama
		default:
			fmt.Println("Input pilih menu anda tidak sesuai. Coba lagi.")
		}
	}
}

// Fungsi utama
func main() {
	var pilihan int
	var peran int
	reader := bufio.NewReader(os.Stdin) // Inisialisasi reader

	for {
		peran = menuLogin() // Panggil fungsi menuLogin untuk mendapatkan peran

		if peran == 0 {
			return // Keluar dari aplikasi jika memilih untuk keluar
		}

		// Menu utama
		for {
			fmt.Println("\nAplikasi Pariwisata")
			fmt.Println("1. Tambah tempat wisata")
			fmt.Println("2. Ubah tempat wisata")
			fmt.Println("3. Hapus tempat wisata")
			fmt.Println("4. Tampilkan tempat wisata terurut")
			fmt.Println("5. Cari tempat wisata")
			fmt.Println("0. Kembali Menu Login")

			if peran == 2 {
				// Pengguna hanya bisa melihat atau mencari tempat wisata
				fmt.Println("Pengguna hanya dapat memilih opsi 4 dan 5.")
			}

			fmt.Print("Pilih menu : ")
			pilihanStr, _ := reader.ReadString('\n')
			pilihanStr = strings.TrimSpace(pilihanStr)
			pilihan, _ = strconv.Atoi(pilihanStr)

			if peran == 2 && (pilihan == 1 || pilihan == 2 || pilihan == 3) { // jika pengguna memilih opsi 1, 2, atau 3 maka tampilkan pesan pengguna tidak memiliki akses
				fmt.Println("Anda tidak memiliki akses untuk opsi ini.")
			} else {
				switch pilihan { // pilihan menu
				case 1: // tambah tempat wisata
					var nama, fasilitas, wahana string
					var jarak, biaya int
					fmt.Print("Nama : ")
					nama, _ = reader.ReadString('\n')
					nama = strings.TrimSpace(nama) // Menghapus spasi di awal dan akhir
					fmt.Print("Fasilitas : ")
					fasilitas, _ = reader.ReadString('\n')
					fasilitas = strings.TrimSpace(fasilitas)

					// Membaca jarak dan biaya sebagai string, kemudian mengonversinya ke int
					fmt.Print("Jarak (km) : ")
					jarakStr, _ := reader.ReadString('\n')
					jarakStr = strings.TrimSpace(jarakStr)
					jarak, _ = strconv.Atoi(jarakStr)

					fmt.Print("Wahana : ")
					wahana, _ = reader.ReadString('\n')
					wahana = strings.TrimSpace(wahana)

					fmt.Print("Biaya (Rp) : ")
					biayaStr, _ := reader.ReadString('\n')
					biayaStr = strings.TrimSpace(biayaStr)
					biaya, _ = strconv.Atoi(biayaStr)

					tambahTempat(nama, fasilitas, jarak, wahana, biaya)

				case 2: // ubah tempat wisata
					var namaAsli, fasilitas, nama, wahana string
					var jarak, biaya int
					fmt.Print("Nama tempat yang akan diubah : ")
					namaAsli, _ = reader.ReadString('\n')
					namaAsli = strings.TrimSpace(namaAsli)

					fmt.Print("Nama baru : ")
					nama, _ = reader.ReadString('\n')
					nama = strings.TrimSpace(nama)

					fmt.Print("Fasilitas baru : ")
					fasilitas, _ = reader.ReadString('\n')
					fasilitas = strings.TrimSpace(fasilitas)

					// Membaca jarak dan biaya sebagai string, kemudian mengonversinya ke int
					fmt.Print("Jarak baru (km) : ")
					jarakStr, _ := reader.ReadString('\n')
					jarakStr = strings.TrimSpace(jarakStr)
					jarak, _ = strconv.Atoi(jarakStr)

					fmt.Print("Wahana baru : ")
					wahana, _ = reader.ReadString('\n')
					wahana = strings.TrimSpace(wahana)

					fmt.Print("Biaya baru (Rp) : ")
					biayaStr, _ := reader.ReadString('\n')
					biayaStr = strings.TrimSpace(biayaStr)
					biaya, _ = strconv.Atoi(biayaStr)

					// Panggil fungsi ubahTempat untuk memperbarui data
					ubahTempat(namaAsli, TempatWisata{nama, fasilitas, jarak, wahana, biaya})

				case 3: // hapus tempat wisata
					var nama string
					fmt.Print("Nama tempat yang akan dihapus : ")
					nama, _ = reader.ReadString('\n')
					nama = strings.TrimSpace(nama) // Menghapus spasi di awal dan akhir
					hapusTempat(nama)

				case 4: // tampilkan tempat wisata terurut
					var kategori, urut string
					fmt.Print("Kategori (jarak/biaya/fasilitas) : ")
					kategori, _ = reader.ReadString('\n')
					kategori = strings.TrimSpace(kategori)

					fmt.Print("Urutkan (ascending/descending) : ")
					urut, _ = reader.ReadString('\n')
					urut = strings.TrimSpace(urut)

					if kategori == "jarak" {
						urutkanJarak(urut)
					} else if kategori == "biaya" {
						urutkanBiaya(urut)
					} else if kategori == "fasilitas" {
						urutkanFasilitas(urut)
					} else {
						fmt.Println("Kategori tidak sesuai.")
					}

				case 5: // cari tempat wisata
					var kategori, kataKunci string

					fmt.Print("Kategori (nama/jarak/biaya/fasilitas) : ")
					kategori, _ = reader.ReadString('\n')
					kategori = strings.TrimSpace(kategori)

					fmt.Print("Kata kunci : ")
					kataKunci, _ = reader.ReadString('\n')
					kataKunci = strings.TrimSpace(kataKunci)

					if kategori == "nama" {
						cariNama(kataKunci)
					} else if kategori == "jarak" {
						cariJarak(kataKunci)
					} else if kategori == "biaya" {
						cariBiaya(kataKunci)
					} else if kategori == "fasilitas" {
						cariFasilitas(kataKunci)
					} else {
						fmt.Println("Kategori tidak sesuai.")
					}

				case 0: // kembali ke menu login
					goto repeat

				default: // pilihan tidak sesuai
					fmt.Println("Menu tidak sesuai.")
				}
			}
		}
	repeat:
	}
}
