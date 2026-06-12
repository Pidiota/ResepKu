package main
import "fmt"
const NMAX int = 100

type resep struct {
	nama     string
	bahan    string
	langkah  [NMAX]string
	nLangkah int
	durasi   int
}
type tabResep [NMAX]resep

func main() {
	var menu, n int
	var R tabResep
	gudangResep(&R, &n)
	fmt.Println()
	fmt.Println("^_^ Selamat Datang di ResepKu ^_^")
	fmt.Println()
	fmt.Println("------- Menu Utama -------")
	fmt.Println("Silahkan pilih menu di bawah ini:")
	fmt.Println("1 | Cari resep")
	fmt.Println("2 | Tambah resep")
	fmt.Println("3 | Ubah resep")
	fmt.Println("4 | Hapus resep")
	fmt.Println("0 | Keluar")
	fmt.Print("Menu : ")
	fmt.Scan(&menu)
	fmt.Println()
	for menu != 0 {
		if menu == 1 {
			cariResep(R, n)
		} else if menu == 2 {
			tambahMenu(&R, &n)
		} else if menu == 3 {
			ubahMenu(&R, n)
		} else if menu == 4 {
			hapusMenu(&R, &n)
		} else {
			fmt.Println("!! Input tidak valid !!")
		}
		sortResep(&R, n)
		fmt.Println()
		fmt.Println("------- Menu Utama -------")
		fmt.Println("1 | Cari resep")
		fmt.Println("2 | Tambah resep")
		fmt.Println("3 | Ubah resep")
		fmt.Println("4 | Hapus resep")
		fmt.Println("0 | Keluar")
		fmt.Print("Menu : ")
		fmt.Scan(&menu)
		fmt.Println()
	}
	fmt.Println("^_^ Terima kasih telah menggunakan ResepKu ^_^")
	fmt.Println()
}

func cariResep(R tabResep, n int) {
	// {I.S. terdefinisi array R yang berisi n data resep
	//  F.S. menampilkan hasil pencarian atau daftar resep sesuai pilihan pengguna}
	var menu, resep int
	fmt.Println("----- Cari Resep -----")
	fmt.Println("Silahkan pilih menu di bawah ini")
	fmt.Println("1 | Tampilkan seluruh resep")
	fmt.Println("2 | Carikan resep tertentu")
	fmt.Println("0 | Kembali")
	fmt.Print("Menu : ")
	fmt.Scan(&menu)
	fmt.Println()
	for menu != 0 {
		if menu == 1 {
			if n == 0 {
				fmt.Println("Belum ada resep yang tersedia.")
			} else {
				showResep(R, n)
				fmt.Println()
				fmt.Println("Silahkan pilih resep yang ingin anda lihat")
				fmt.Print("Resep : ")
				fmt.Scan(&resep)
				if resep >= 1 && resep <= n {
					showResepDetail(R, resep)
				} else {
					fmt.Println("Nomor resep tidak valid.")
				}
			}
		} else if menu == 2 {
			searchResep(R, n)
		} else {
			fmt.Println("!! Input tidak valid !!")
		}
		fmt.Println()
		fmt.Println("Silahkan pilih menu di bawah ini")
		fmt.Println("1 | Tampilkan seluruh resep")
		fmt.Println("2 | Carikan resep tertentu")
		fmt.Println("0 | Kembali")
		fmt.Print("Menu : ")
		fmt.Scan(&menu)
		fmt.Println()
	}
}

func showResep(R tabResep, n int) {
	// {I.S. terdefinisi array R berisi n data resep
	//  F.S. menampilan semua data resep}
	var i int
	sortResep(&R, n)
	fmt.Println("No | Nama Resep           | Bahan      | Durasi")
	fmt.Println("---+----------------------+------------+--------")
	for i = 0; i < n; i++ {
		fmt.Printf("%-2d | %-20s | %-10s | %d menit\n", i+1, R[i].nama, R[i].bahan, R[i].durasi)
	}
}

func showResepDetail(R tabResep, resep int) {
	// {I.S. terdefinisi array R dan resep sebagai nomor resep yang dicari
	//  F.S. detail resep yang dipilih ditampilkan ke layar}
	var i, j int
	i = resep - 1
	fmt.Println()
	fmt.Println("Nama Resep        : ", R[i].nama)
	fmt.Println("Bahan Utama resep : ", R[i].bahan)
	fmt.Println("Durasi (menit)    : ", R[i].durasi)
	fmt.Println("Langkah-langkah:")
	for j = 0; j < R[i].nLangkah; j++ {
		fmt.Println(j+1, ".", R[i].langkah[j])
	}
}

func searchResep(R tabResep, n int) {
	// {I.S. terdefinisi array R berisi n data resep
	//  F.S. menampilkan data resep berdasarkan bahan utama}
	var pilihan int
	var i, j, resep int
	var bahan string
	var daftar [NMAX]int
	fmt.Println("Pilih nomor bahan utama resep")
	fmt.Println("1 | Ayam")
	fmt.Println("2 | Ikan")
	fmt.Println("3 | Sapi")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		bahan = "Ayam"
	} else if pilihan == 2 {
		bahan = "Ikan"
	} else if pilihan == 3 {
		bahan = "Sapi"
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	if pilihan == 1 || pilihan == 2 || pilihan == 3 {
		fmt.Println()
		fmt.Println("Resep dengan bahan utama", bahan)
		j = 0
		for i = 0; i < n; i++ {
			if R[i].bahan == bahan {
				fmt.Println(j+1, "|", R[i].nama)
				daftar[j] = i
				j++
			}
		}
		if j == 0 {
			fmt.Println("Tidak ada resep ditemukan.")
		} else {
			fmt.Println()
			fmt.Print("Pilih resep : ")
			fmt.Scan(&resep)

			if resep >= 1 && resep <= j {
				showResepDetail(R, daftar[resep-1]+1)
			} else {
				fmt.Println("Nomor resep tidak valid.")
			}
		}
	}
}

func sortResep(R *tabResep, n int) {
	// {I.S. terdefinisi array R yang berisi n data resep
	//  F.S. array R terurut secara ASCENDING/naik berdasarkan nama menggunakan algoritma insertion sort}
	var pass, k int
	var temp resep
	for pass = 1; pass < n; pass++ {
		k = pass
		temp = R[k]
		for k > 0 && temp.nama < R[k-1].nama {
			R[k] = R[k-1]
			k--
		}
		R[k] = temp
	}
}

func tambahMenu(R *tabResep, n *int) {
	// {I.S. terdefinisi array R yang berisi n data resep
	//  F.S. data resep baru ditambahkan ke array R dan n bertambah satu}
	var i, j int
	var langkah string
	i = *n
	fmt.Println("----- Tambah Resep -----")
	fmt.Println("Silahkan isi data resep baru anda")
	fmt.Print("Nama resep        : ")
	fmt.Scan(&R[i].nama)
	fmt.Print("Bahan utama resep : ")
	fmt.Scan(&R[i].bahan)
	fmt.Print("Durasi (menit)    : ")
	fmt.Scan(&R[i].durasi)
	fmt.Println()
	fmt.Println("Masukkan langkah-langkah resep")
	fmt.Println("Ketik END untuk selesai")
	j = 0
	fmt.Print(j+1, ". ")
	fmt.Scan(&langkah)
	for langkah != "END" && j < NMAX {
		R[i].langkah[j] = langkah
		j++

		fmt.Print(j+1, ". ")
		fmt.Scan(&langkah)
	}
	R[i].nLangkah = j
	*n = *n + 1
	fmt.Println("Resep berhasil ditambahkan.")
}

func ubahMenu(R *tabResep, n int) {
	// {I.S. terdefinisi array R yang berisi n data resep
	//  F.S. mengubah salah satu data resep tertentu di array R}
	var idx int
	var j int
	var langkah string
	if n == 0 {
		fmt.Println("Belum ada resep.")
	} else {
		showResep(*R, n)
		fmt.Print("Pilih nomor resep yang ingin diubah : ")
		fmt.Scan(&idx)
		if idx < 1 || idx > n {
			fmt.Println("Nomor resep tidak valid.")
		} else {
			idx = idx - 1
			fmt.Print("Nama resep baru        : ")
			fmt.Scan(&R[idx].nama)
			fmt.Print("Bahan utama resep baru : ")
			fmt.Scan(&R[idx].bahan)
			fmt.Print("Durasi baru (menit)    : ")
			fmt.Scan(&R[idx].durasi)
			fmt.Println("Masukkan langkah-langkah baru")
			fmt.Println("Ketik END untuk selesai")
			j = 0
			fmt.Print(j+1, ". ")
			fmt.Scan(&langkah)
			for langkah != "END" && j < NMAX {
				R[idx].langkah[j] = langkah
				j++

				fmt.Print(j+1, ". ")
				fmt.Scan(&langkah)
			}
			R[idx].nLangkah = j
			fmt.Println("Resep berhasil diubah.")
		}
	}
}

func hapusMenu(R *tabResep, n *int) {
	// {I.S. terdefinisi array R yang berisi n data resep
	//  F.S. satu data resep dihapus dari array R dan n berkurang satu}
	var idx int
	var i int
	if *n == 0 {
		fmt.Println("Belum ada resep.")
	} else {
		showResep(*R, *n)
		fmt.Print("Pilih nomor resep yang ingin dihapus : ")
		fmt.Scan(&idx)
		if idx < 1 || idx > *n {
			fmt.Println("Nomor resep tidak valid.")
		} else {
			idx = idx - 1
			for i = idx; i < *n-1; i++ {
				R[i] = R[i+1]
			}
			*n = *n - 1
			fmt.Println("Resep berhasil dihapus.")
		}
	}
}

func gudangResep(R *tabResep, n *int) {
	// {I.S. terdefinisi array R kosong dan n sembarang
	//  F.S. array R terisi data resep awal dan n berisi jumlah resep}
	var i int
	i = 0

	// Ayam Goreng
	R[i].nama = "Ayam_goreng"
	R[i].bahan = "Ayam"
	R[i].durasi = 30
	R[i].langkah[0] = "Cuci_ayam"
	R[i].langkah[1] = "Lumuri_bumbu"
	R[i].langkah[2] = "Diamkan_15_menit"
	R[i].langkah[3] = "Goreng_hingga_matang"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Ayam Bakar
	R[i].nama = "Ayam_bakar"
	R[i].bahan = "Ayam"
	R[i].durasi = 40
	R[i].langkah[0] = "Cuci_ayam"
	R[i].langkah[1] = "Lumuri_bumbu"
	R[i].langkah[2] = "Diamkan_15_menit"
	R[i].langkah[3] = "Bakar_hingga_matang"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Ayam Kecap
	R[i].nama = "Ayam_kecap"
	R[i].bahan = "Ayam"
	R[i].durasi = 35
	R[i].langkah[0] = "Cuci_ayam"
	R[i].langkah[1] = "Tumis_bumbu"
	R[i].langkah[2] = "Masukkan_ayam"
	R[i].langkah[3] = "Tambahkan_kecap"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Ikan Goreng
	R[i].nama = "Ikan_goreng"
	R[i].bahan = "Ikan"
	R[i].durasi = 25
	R[i].langkah[0] = "Bersihkan_ikan"
	R[i].langkah[1] = "Lumuri_bumbu"
	R[i].langkah[2] = "Diamkan_10_menit"
	R[i].langkah[3] = "Goreng_hingga_matang"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Ikan Bakar
	R[i].nama = "Ikan_bakar"
	R[i].bahan = "Ikan"
	R[i].durasi = 35
	R[i].langkah[0] = "Bersihkan_ikan"
	R[i].langkah[1] = "Lumuri_bumbu"
	R[i].langkah[2] = "Diamkan_10_menit"
	R[i].langkah[3] = "Bakar_hingga_matang"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Ikan Asam Manis
	R[i].nama = "Ikan_asam_manis"
	R[i].bahan = "Ikan"
	R[i].durasi = 40
	R[i].langkah[0] = "Bersihkan_ikan"
	R[i].langkah[1] = "Goreng_ikan"
	R[i].langkah[2] = "Masak_saos"
	R[i].langkah[3] = "Campur_dengan_ikan"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Rendang
	R[i].nama = "Rendang"
	R[i].bahan = "Sapi"
	R[i].durasi = 120
	R[i].langkah[0] = "Potong_daging"
	R[i].langkah[1] = "Masak_bumbu"
	R[i].langkah[2] = "Masukkan_daging"
	R[i].langkah[3] = "Masak_hingga_empuk"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Semur Daging
	R[i].nama = "Semur_daging"
	R[i].bahan = "Sapi"
	R[i].durasi = 90
	R[i].langkah[0] = "Potong_daging"
	R[i].langkah[1] = "Tumis_bumbu"
	R[i].langkah[2] = "Masukkan_daging"
	R[i].langkah[3] = "Tambahkan_kecap"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	// Sate Daging
	R[i].nama = "Sate_daging"
	R[i].bahan = "Sapi"
	R[i].durasi = 45
	R[i].langkah[0] = "Potong_daging"
	R[i].langkah[1] = "Tusuk_dengan_sate"
	R[i].langkah[2] = "Lumuri_bumbu"
	R[i].langkah[3] = "Bakar_hingga_matang"
	R[i].langkah[4] = "Sajikan"
	R[i].nLangkah = 5
	i++

	*n = i
}