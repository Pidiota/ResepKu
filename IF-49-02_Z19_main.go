package main
import "fmt"

const NMAX int = 100
type resep struct{
	nama string
	bahan string
	langkah string
	durasi int

}

type tabResep [NMAX]resep

func main() {
	var x, opsi int
	var target string
	var T tabResep
	x = 99
	for x != 0 {
		fmt.Println("----- selamat datang di ResepKu -----")
		fmt.Println("Silahkan pilih opsi di bawah ini:")
		fmt.Println("1 - Cari resep")
		fmt.Println("2 - Tambah resep")
		fmt.Println("3 - Ubah resep")
		fmt.Println("4 - Hapus resep")
		fmt.Println("0 - Keluar")
		fmt.Scan(&opsi)
		for opsi != 0 {
			if opsi == 1 {
				cariResep(T, target)
				fmt.Println("Silahkan pilih opsi di bawah ini:")
				fmt.Println("1 - Cari resep")
				fmt.Println("2 - Tambah resep")
				fmt.Println("3 - Ubah resep")
				fmt.Println("4 - Hapus resep")
				fmt.Println("0 - Keluar")
				fmt.Scan(&opsi)
			} else if opsi == 2 {
				tambahMenu(&T)
				fmt.Println("Silahkan pilih opsi di bawah ini:")
				fmt.Println("1 - Cari resep")
				fmt.Println("2 - Tambah resep")
				fmt.Println("3 - Ubah resep")
				fmt.Println("4 - Hapus resep")
				fmt.Println("0 - Keluar")
				fmt.Scan(&opsi)
			} else if opsi == 3 {
				//
				fmt.Println("Silahkan pilih opsi di bawah ini:")
				fmt.Println("1 - Cari resep")
				fmt.Println("2 - Tambah resep")
				fmt.Println("3 - Ubah resep")
				fmt.Println("4 - Hapus resep")
				fmt.Println("0 - Keluar")
				fmt.Scan(&opsi)
			} else if opsi == 4 {
				//
				fmt.Println("Silahkan pilih opsi di bawah ini:")
				fmt.Println("1 - Cari resep")
				fmt.Println("2 - Tambah resep")
				fmt.Println("3 - Ubah resep")
				fmt.Println("4 - Hapus resep")
				fmt.Println("0 - Keluar")
				fmt.Scan(&opsi)
			} else {
				fmt.Println("input tidak valid")
				fmt.Println("Silahkan pilih opsi di bawah ini:")
				fmt.Println("1 - Cari resep")
				fmt.Println("2 - Tambah resep")
				fmt.Println("3 - Ubah resep")
				fmt.Println("4 - Hapus resep")
				fmt.Println("0 - Keluar")
				fmt.Scan(&opsi)
			}
		}
	}
	
	
}

func sortMenu(T tabResep) {
	var pass, k, n int
	var temp resep
	n ;= 100
	for pass = 1; pass < n; pass++{
		k = pass
		temp = T[k]
		for k > 0 && temp.nama > T[k-1].nama{
			T[k] = T[k-1]
			k = k-1
		}
		T[k] = temp	
	}
}


func cariResep(T tabResep, target string) {
	fmt.Println("Silahkan Masukkan Nama Menu Yang Anda Cari")
	fmt.Scan(&target)
	n := 100
	for i := 0; i < n; i++{
		if T[i].nama == target {
			fmt.Println(T[i].nama, T[i].bahan, T[i].langkah, T[i].durasi )
		}
	}
	fmt.Println("Entahlah Isi Apa")
}


func tambahMenu(T *tabResep){
	var n, i int	
	for i = 5 ; i < n; i++{
		fmt.Print("Nama Menu :")
		fmt.Scan(&T[i].nama)
		fmt.Print("Detail Bahan :")
		fmt.Scan(&T[i].bahan)
		fmt.Print("Langkah - langkah :")
		fmt.Scan(&T[i].langkah)
		fmt.Print("Durasi")
		fmt.Scan(&T[i].durasi)
	} 
}

//func ubahMenu() {

//}

//func hapusMenu() {

//}


