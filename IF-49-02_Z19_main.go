package main
import "fmt"

func main() {
	var x int
	fmt.Println("----- selamat datang di ResepKu -----")
	
}

func pilihMenu(opsi *int) {
	fmt.Println("Silahkan pilih opsi di bawah ini:")
	fmt.Println("1 - Cari resep")
	fmt.Println("2 - Tambah resep")
	fmt.Println("3 - Ubah resep")
	fmt.Println("4 - Hapus resep")
	fmt.Println("0 - Keluar")
	fmt.Scan(&opsi)
	if opsi == 1 {

	} else if opsi == 2 {

	} else if opsi == 3 {

	} else if opsi == 4 {

	} else{

	}
}

func cariResep() {
	for i := 0; i < n; i++{
		if x[i].nilai == target {
			fmt.Println(x[i].matkul, x[i].nilai)
		}
	}
	fmt.Println("Data Ditemukan!")
}


func tambahMenu(){

}

func ubahMenu() {

}

func hapusMenu() {

}

