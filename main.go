package main

import (
	"fmt"
	"os"
	"strconv"
)

type biodata struct {
	nama    string
	alamat  string
	pekerjaan string
	alasan string
}

var dataList = []biodata{
	{
		nama:    "Freddy Oktoniyer S",
		alamat:  "Cikarang",
		pekerjaan: "Application Developer",
		alasan: "Ingin meningkatkan kemampuan pemrograman",
	},
	{
		nama:    "Arman Humanda",
		alamat:  "Dihatimu",
		pekerjaan: "Pengangguran Kece",
		alasan: "Pengen aja nyoba nyoba",
	},
	{
		nama:    "Atri",
		alamat:  "Jakarta",
		pekerjaan: "Data Science",
		alasan: "Ingin mengembangkan skill backend",
	},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go [nomor absen]")
		return
	}

	absen := args[1]
	if absen == "" {
		fmt.Println("Nomor absen tidak boleh kosong")
		return
	}

	// ubah nomor absen ke dalam bentuk integer
	absenInt, err := strconv.Atoi(absen)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	// ambil data berdasarkan nomor absen
	if absenInt < 1 || absenInt > len(dataList) {
		fmt.Printf("Tidak ada data dengan nomor absen %d\n", absenInt)
		return
	}

	selectedData := dataList[absenInt-1]

	// tampilkan Data yang dipilih
	fmt.Printf("Nama: %s\n", selectedData.nama)
	fmt.Printf("Alamat: %s\n", selectedData.alamat)
	fmt.Printf("Pekerjaan: %s\n", selectedData.pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", selectedData.alasan)
}
