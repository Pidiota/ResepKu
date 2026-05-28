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
		opsi = 0

		for opsi != 99 {
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


func cariResep(T tabResep, target string) {
	fmt.Println("Silahkan Masukkan Nama Menu Yang Anda Cari")
	fmt.Scan(&target)

	for i := 0; i < n; i++{ // Bagian looping untuk scannya bingung kondisinya gmana
		if T[i].nama == target {
			fmt.Println(T[i].nama, T[i].bahan, T[i].langkah, T[i].durasi )
		}
	}
	fmt.Println("Entahlah Isi Apa")
}


//func tambahMenu(){

//}

//func ubahMenu() {

//}

//func hapusMenu() {

//}

