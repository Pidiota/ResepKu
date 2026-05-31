package main
import "fmt"

type resep struct{
	nama string
	bahan string
	langkah string
	durasi int

}

type tabResep [100]resep

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
		fmt.Println("99 - Keluar")
		opsi = 99

		for opsi != 0 {
			fmt.Scan(&opsi)
			if opsi == 1 {
				cariResep(T, target)
			} else if opsi == 2 {

			} else if opsi == 3 {

			} else if opsi == 4 {

			} else{
				fmt.Println("Input Tidak Valid")
				opsi = 0
			}
		}
		

	}
	
	
}

func sortMenu(T tabResep) {
	var pass, k, n int
	var temp resep
	n = 99
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
	n := 99

	for i := 0; i < n; i++{
		if T[i].nama == target {
			fmt.Println(T[i].nama, T[i].bahan, T[i].langkah, T[i].durasi )
		}
	}
	fmt.Println("Entahlah Isi Apa")
}


//func tambahMenu(T tabResep){

//}

//func ubahMenu() {

//}

//func hapusMenu() {

//}

