package main

import (
	"fmt"
	"strings"
)

const MaxData = 100 // kapasitas maksimum data

type TempatWisata struct { // tipe data TempatWisata berisi Nama, Fasilitas, Jarak, dan Biaya
	Nama       string
	Fasilitas  string
	Jarak      int
	Biaya      int
}

// daftar tempat wisata dengan kapasitas maksimum 100
var daftarTempat [MaxData]TempatWisata
var jumlahData int // menyimpan jumlah tempat wisata yg ada

// fungsi inisialisasi utk menambahkan data awal
func init() {
	dataAwal := []TempatWisata{
		{"Baturraden", "Tempat Parkir, Toilet, Warung Makan, Penginapan", 15, 20000},
		{"Curug Cipendok", "Tempat Parkir, Toilet, Warung Makan", 25, 10000},
		{"Telaga Sunyi", "Tempat Parkir, Toilet, Spot Foto", 17, 15000},
		{"Limpak Kuwus", "Tempat Parkir, Spot Foto, Warung Makan", 20, 20000},
		{"Taman Balai Kemambang", "Tempat Parkir, Toilet, Area Bermain Anak", 5, 10000},
	}

	// menyalin data awal ke dalam array daftarTempat
	for _, tempat := range dataAwal {
		daftarTempat[jumlahData] = tempat
		jumlahData++
	}
}

func tambahTempat(nama, fasilitas string, jarak, biaya int) { // menambahkan tempat wisata baru ke dalam array daftarTempat
	if jumlahData < MaxData {
		daftarTempat[jumlahData] = TempatWisata{nama, fasilitas, jarak, biaya}
		jumlahData++
		fmt.Println("Data telah ditambah.")
	} else {
		fmt.Println("Data telah penuh, tidak dapat menambahkan data tempat baru.")
	}
}

func ubahTempat(nama string, tempatBaru TempatWisata) {
    // Urutkan data
    urutkanPilih(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
        return strings.ToLower(a.Nama) < strings.ToLower(b.Nama)
    })

    fmt.Println("Data setelah diurutkan:")
    for _, tempat := range daftarTempat[:jumlahData] {
        fmt.Println(tempat.Nama)
    }

    // Cari data
    idx := cariBiner(daftarTempat[:jumlahData], nama)
    if idx != -1 {
        daftarTempat[idx] = tempatBaru
        fmt.Println("Data telah diubah.")
    } else {
        fmt.Println("Tempat tidak ditemukan.")
    }
}

func hapusTempat(nama string) {
	// Urutkan data berdasarkan nama
	urutkanPilih(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
		return a.Nama < b.Nama
	})

	idx := cariBiner(daftarTempat[:jumlahData], nama)
	if idx != -1 {
		for i := idx; i < jumlahData-1; i++ {
			daftarTempat[i] = daftarTempat[i+1]
		}
		jumlahData--
		fmt.Println("Data telah dihapus.")
	} else {
		fmt.Println("Tempat tidak ditemukan.")
	}
}

func tampilkanTerurut(kategori string, urut bool) { // menampilkan data tempat wisata yang sudah diurutkan berdasarkan kategori dan urutan tertentu
	switch kategori {
	case "jarak": // jika kategori jarak, maka urutkan berdasarkan jarak tempat wisata
		urutkanPilih(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
			if urut {
				return a.Jarak < b.Jarak
			}
			return a.Jarak > b.Jarak
		})
	case "biaya": // jika kategori biaya, maka urutkan berdasarkan biaya tempat wisata
		urutkanSisip(daftarTempat[:jumlahData], func(a, b TempatWisata) bool {
			if urut {
				return a.Biaya < b.Biaya
			}
			return a.Biaya > b.Biaya
		})
	default:
		fmt.Println("Kategori tidak sesuai.")
		return
	}
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d km, Rp%d) - %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas)
	}
}

func cariTempat(kataKunci string) { // mencari data tempat wisata berdasarkan kata kunci yang dimasukkan pengguna
	fmt.Println("Hasil pencarian:")
	for i := 0; i < jumlahData; i++ {
		if strings.Contains(strings.ToLower(daftarTempat[i].Nama), strings.ToLower(kataKunci)) {
			fmt.Printf("%d. %s (%d km, Rp%d) - %s\n", i+1, daftarTempat[i].Nama, daftarTempat[i].Jarak, daftarTempat[i].Biaya, daftarTempat[i].Fasilitas)
		}
	}
}

func cariUrut(data []TempatWisata, kunci string) int { // mencari data tempat wisata berdasarkan nama tempat dengan metode sequential search
	for i, tempat := range data {
		if strings.ToLower(tempat.Nama) == strings.ToLower(kunci) {
			return i
		}
	}
	return -1
}

func cariBiner(data []TempatWisata, kunci string) int { // mencari data tempat wisata berdasarkan nama tempat dengan metode binary search
	rendah, tinggi := 0, len(data)-1
	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2
		if strings.ToLower(data[tengah].Nama) == strings.ToLower(kunci) {
			return tengah
		} else if strings.ToLower(data[tengah].Nama) < strings.ToLower(kunci) {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}
	return -1
}

func urutkanPilih(data []TempatWisata, banding func(a, b TempatWisata) bool) { // mengurutkan data tempat wisata dengan metode selection sort
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

func urutkanSisip(data []TempatWisata, banding func(a, b TempatWisata) bool) { // mengurutkan data tempat wisata dengan metode insertion sort
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

func main() {
	var pilihan int
	var peran int
	var selesai bool 

	// menu login
	for !selesai {
		fmt.Println("\n--- Menu Login ---")
		fmt.Println("1. Admin")
		fmt.Println("2. Pengguna")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1-3) : ")
		fmt.Scan(&peran)

		if peran == 3 { 
			fmt.Println("Thank you telah menggunakan Aplikasi Pariwisata.")
			return
		}

		if peran == 1 || peran == 2 { 
			selesai = true
		} else {
			fmt.Println("Input pilih menu anda tidak sesuai. Coba lagi.")
		}
	}

	selesai = false 
	for !selesai { // menu utama
		fmt.Println("\nAplikasi Pariwisata")
		fmt.Println("1. Tambah tempat wisata")
		fmt.Println("2. Ubah tempat wisata")
		fmt.Println("3. Hapus tempat wisata")
		fmt.Println("4. Tampilkan tempat wisata terurut")
		fmt.Println("5. Cari tempat wisata")
		fmt.Println("0. Keluar")

		if peran == 2 {
			// Pengguna hanya bisa melihat atau mencari tempat wisata
			fmt.Println("Pengguna hanya dapat memilih opsi 4 dan 5.")
		}

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if peran == 2 && (pilihan == 1 || pilihan == 2 || pilihan == 3) { // jika pengguna memilih opsi 1, 2, atau 3 maka tampilkan pesan pengguna tidak memiliki akses
			fmt.Println("Anda tidak memiliki akses untuk opsi ini.")
		} else {
			switch pilihan { // pilihan menu
			case 1: // tambah tempat wisata
				var nama, fasilitas string
				var jarak, biaya int
				fmt.Print("Nama : ")
				fmt.Scan(&nama)
				fmt.Print("Fasilitas : ")
				fmt.Scan(&fasilitas)
				fmt.Print("Jarak (km) : ")
				fmt.Scan(&jarak)
				fmt.Print("Biaya (Rp) : ")
				fmt.Scan(&biaya)
				tambahTempat(nama, fasilitas, jarak, biaya)
			case 2: // ubah tempat wisata
				var nama, fasilitas string
				var jarak, biaya int
				fmt.Print("Nama tempat yang akan diubah : ")
				fmt.Scan(&nama)
				fmt.Print("Nama baru : ")
				fmt.Scan(&nama)
				fmt.Print("Fasilitas baru : ")
				fmt.Scan(&fasilitas)
				fmt.Print("Jarak baru (km) : ")
				fmt.Scan(&jarak)
				fmt.Print("Biaya baru (Rp) : ")
				fmt.Scan(&biaya)
				ubahTempat(nama, TempatWisata{nama, fasilitas, jarak, biaya})
			case 3: // hapus tempat wisata
				var nama string
				fmt.Print("Nama tempat yang akan dihapus : ")
				fmt.Scan(&nama)
				hapusTempat(nama)
			case 4: // tampilkan tempat wisata terurut
				var kategori string
				var urut bool
				fmt.Print("Kategori (jarak/biaya) : ")
				fmt.Scan(&kategori)
				fmt.Print("Urutkan (true/false) : ")
				fmt.Scan(&urut)
				tampilkanTerurut(kategori, urut)
			case 5: // cari tempat wisata
				var kataKunci string
				fmt.Print("Kata kunci : ")
				fmt.Scan(&kataKunci)
				cariTempat(kataKunci)
			case 0: // keluar
				fmt.Println("Thank you telah menggunakan Aplikasi Pariwisata.")
				selesai = true
			default: // pilihan tidak sesuai
				fmt.Println("Menu tidak sesuai.")
			}
		}
	}
}