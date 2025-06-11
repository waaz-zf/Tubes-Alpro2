package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAKS_FILM = 100
const MAKS_SUTRADARA = 50
const MAKS_KATEGORI_GENRE = 30
type TipeFilm struct {
	ID           int
	Judul        string
	Sutradara    string
	Genre        string
	TahunRilis   int
	Rating       float64
	StatusTonton string
}

type TipeSutradara struct {
	IDSutradara int
	Nama        string
	JumlahFilm  int
}

type TipeKategoriGenre struct {
	IDGenre    int
	Nama       string
	JumlahFilm int
}

var DaftarFilm [MAKS_FILM]TipeFilm
var JumlahFilm int
var DaftarSutradara [MAKS_SUTRADARA]TipeSutradara
var JumlahSutradara int
var DaftarKategoriGenre [MAKS_KATEGORI_GENRE]TipeKategoriGenre
var JumlahKategoriGenre int
func bacaString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func bacaInt(prompt string) int {
	var inputStr string
	var nilaiInt int
	var err error
	for {
		fmt.Print(prompt)
		_, err = fmt.Scanln(&inputStr)
		if err == nil {
			nilaiInt, err = strconv.Atoi(strings.TrimSpace(inputStr))
			if err == nil {
				return nilaiInt
			}
		}
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat.")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

func bacaFloat64(prompt string) float64 {
	var inputStr string
	var nilaiFloat float64
	var err error
	for {
		fmt.Print(prompt)
		_, err = fmt.Scanln(&inputStr)
		if err == nil {
			nilaiFloat, err = strconv.ParseFloat(strings.TrimSpace(inputStr), 64)
			if err == nil {
				return nilaiFloat
			}
		}
		fmt.Println("Input tidak valid. Harap masukkan bilangan desimal (gunakan titik sebagai pemisah desimal).")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

func enterLanjut() {
	fmt.Println("\nTekan ENTER untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func garisPisah() {
	fmt.Println("----------------------------------------------------------------------------------------------------")
}

func spaceScreen() {
	fmt.Print("\033[H\033[2J")
}

func generateNewID() int {
	if JumlahFilm == 0 {
		return 1
	}
	return DaftarFilm[JumlahFilm-1].ID + 1
}

func generateNewIDSutradara() int {
	if JumlahSutradara == 0 {
		return 1
	}
	return DaftarSutradara[JumlahSutradara-1].IDSutradara + 1
}

func generateNewIDGenre() int {
	if JumlahKategoriGenre == 0 {
		return 1
	}
	return DaftarKategoriGenre[JumlahKategoriGenre-1].IDGenre + 1
}

func cariSutradara(nama string) int {
	namaLower := strings.ToLower(nama)
	for i := 0; i < JumlahSutradara; i++ {
		if strings.ToLower(DaftarSutradara[i].Nama) == namaLower {
			return i
		}
	}
	return -1
}

func cariGenre(nama string) int {
	namaLower := strings.ToLower(nama)
	for i := 0; i < JumlahKategoriGenre; i++ {
		if strings.ToLower(DaftarKategoriGenre[i].Nama) == namaLower {
			return i
		}
	}
	return -1
}

func tambahSutradaraBaru(namaSutradara string) {
	if JumlahSutradara >= MAKS_SUTRADARA {
		fmt.Println("Kapasitas penyimpanan sutradara penuh.")
		return
	}
	indeksSutradara := cariSutradara(namaSutradara)
	if indeksSutradara == -1 {
		var sutradaraBaru TipeSutradara
		sutradaraBaru.IDSutradara = generateNewIDSutradara()
		sutradaraBaru.Nama = namaSutradara
		sutradaraBaru.JumlahFilm = 1
		DaftarSutradara[JumlahSutradara] = sutradaraBaru
		JumlahSutradara++
	} else {
		DaftarSutradara[indeksSutradara].JumlahFilm++
	}
}

func tambahGenreBaru(namaGenre string) {
	if JumlahKategoriGenre >= MAKS_KATEGORI_GENRE {
		fmt.Println("Kapasitas penyimpanan kategori genre penuh.")
		return
	}
	indeksGenre := cariGenre(namaGenre)
	if indeksGenre == -1 {
		var genreBaru TipeKategoriGenre
		genreBaru.IDGenre = generateNewIDGenre()
		genreBaru.Nama = namaGenre
		genreBaru.JumlahFilm = 1
		DaftarKategoriGenre[JumlahKategoriGenre] = genreBaru
		JumlahKategoriGenre++
	} else {
		DaftarKategoriGenre[indeksGenre].JumlahFilm++
	}
}

func tambahFilm() {
	spaceScreen()
	fmt.Println("--- TAMBAH FILM BARU ---")
	garisPisah()
	if JumlahFilm >= MAKS_FILM {
		fmt.Println("Maaf, kapasitas penyimpanan film sudah penuh.")
		enterLanjut()
		return
	}

	var filmBaru TipeFilm
	filmBaru.ID = generateNewID()
	filmBaru.Judul = bacaString("Masukkan Judul Film: ")
	filmBaru.Sutradara = bacaString("Masukkan Sutradara: ")
	filmBaru.Genre = bacaString("Masukkan Genre: ")
	filmBaru.TahunRilis = bacaInt("Masukkan Tahun Rilis: ")
	for {
		rating := bacaFloat64("Masukkan Rating (1.0 - 5.0): ")
		if rating >= 1.0 && rating <= 5.0 {
			filmBaru.Rating = rating
			break
		}
		fmt.Println("Rating tidak valid. Harap masukkan angka antara 1.0 sampai 5.0.")
	}

	for {
		status := strings.ToLower(bacaString("Masukkan Status Tonton (belum ditonton/sedang ditonton/sudah ditonton): "))
		if status == "belum ditonton" || status == "sedang ditonton" || status == "sudah ditonton" {
			filmBaru.StatusTonton = status
			break
		}
		fmt.Println("Status tidak valid. Harap masukkan 'belum ditonton', 'sedang ditonton', atau 'sudah ditonton'.")
	}
	tambahSutradaraBaru(filmBaru.Sutradara)
	tambahGenreBaru(filmBaru.Genre)
	DaftarFilm[JumlahFilm] = filmBaru
	JumlahFilm++
	fmt.Println("\nFilm berhasil ditambahkan!")
	enterLanjut()
}

func lihatDaftarFilm(daftar [MAKS_FILM]TipeFilm, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada film dalam daftar.")
		return
	}
	garisPisah()
	fmt.Printf("%-5s | %-30s | %-20s | %-15s | %-12s | %-8s | %-20s\n",
		"ID", "Judul", "Sutradara", "Genre", "Tahun Rilis", "Rating", "Status Tonton")
	garisPisah()

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-5d | %-30s | %-20s | %-15s | %-12d | %-8.1f | %-20s\n",
			daftar[i].ID, daftar[i].Judul, daftar[i].Sutradara, daftar[i].Genre,
			daftar[i].TahunRilis, daftar[i].Rating, daftar[i].StatusTonton)
	}
	garisPisah()
}

func pencarianFilm(kataKunci string, cariBerdasarkan string) ([MAKS_FILM]TipeFilm, int) {
	var hasilPencarian [MAKS_FILM]TipeFilm
	jumlahHasil := 0
	kataKunciLower := strings.ToLower(kataKunci)
	for i := 0; i < JumlahFilm; i++ {
		cocok := false
		switch strings.ToLower(cariBerdasarkan) {
		case "judul":
			if strings.Contains(strings.ToLower(DaftarFilm[i].Judul), kataKunciLower) {
				cocok = true
			}
		case "sutradara":
			if strings.Contains(strings.ToLower(DaftarFilm[i].Sutradara), kataKunciLower) {
				cocok = true
			}
		case "genre":
			if strings.Contains(strings.ToLower(DaftarFilm[i].Genre), kataKunciLower) {
				cocok = true
			}
		case "statustonton":
			if strings.Contains(strings.ToLower(DaftarFilm[i].StatusTonton), kataKunciLower) {
				cocok = true
			}
		}
		if cocok {
			hasilPencarian[jumlahHasil] = DaftarFilm[i]
			jumlahHasil++
		}
	}
	return hasilPencarian, jumlahHasil
}

func cariBinerByID(id int) int {
	low := 0
	high := JumlahFilm - 1
	for low <= high {
		mid := (low + high) / 2
		if mid < 0 || mid >= JumlahFilm {
			return -1
		}
		if DaftarFilm[mid].ID == id {
			return mid
		} else if DaftarFilm[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func prosedurCariFilm() {
	spaceScreen()
	fmt.Println("--- CARI FILM ---")
	garisPisah()
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Judul")
	fmt.Println("2. Sutradara")
	fmt.Println("3. Genre")
	fmt.Println("4. Status Tonton")
	fmt.Println("5. ID Film")
	garisPisah()
	pilihan := bacaInt("Pilih kriteria pencarian (1-5): ")

	var kataKunci string
	var hasilPencarian [MAKS_FILM]TipeFilm
	var jumlahHasil int
	indeksDitemukan := -1
	switch pilihan {
	case 1:
		kataKunci = bacaString("Masukkan Judul Film: ")
		hasilPencarian, jumlahHasil = pencarianFilm(kataKunci, "judul")
	case 2:
		kataKunci = bacaString("Masukkan Nama Sutradara: ")
		hasilPencarian, jumlahHasil = pencarianFilm(kataKunci, "sutradara")
	case 3:
		kataKunci = bacaString("Masukkan Genre Film: ")
		hasilPencarian, jumlahHasil = pencarianFilm(kataKunci, "genre")
	case 4:
		kataKunci = bacaString("Masukkan Status Tonton Film (belum ditonton/sedang dititonton/sudah ditonton): ")
		hasilPencarian, jumlahHasil = pencarianFilm(kataKunci, "statustonton")
	case 5:
		id := bacaInt("Masukkan ID Film: ")
		indeksDitemukan = cariBinerByID(id)
		if indeksDitemukan != -1 {
			hasilPencarian[0] = DaftarFilm[indeksDitemukan]
			jumlahHasil = 1
		}
	default:
		fmt.Println("Pilihan tidak valid.")
		enterLanjut()
		return
	}

	if jumlahHasil == 0 {
		fmt.Println("Film tidak ditemukan.")
	} else {
		fmt.Println("\nHasil Pencarian:")
		lihatDaftarFilm(hasilPencarian, jumlahHasil)
	}
	enterLanjut()
}

func selectionSort(kriteria string, urutanNaik bool) {
	for i := 0; i < JumlahFilm-1; i++ {
		indeksMinAtauMaks := i
		for j := i + 1; j < JumlahFilm; j++ {
			perluTukar := false
			switch strings.ToLower(kriteria) {
			case "judul":
				if urutanNaik {
					if strings.ToLower(DaftarFilm[j].Judul) < strings.ToLower(DaftarFilm[indeksMinAtauMaks].Judul) {
						perluTukar = true
					}
				} else {
					if strings.ToLower(DaftarFilm[j].Judul) > strings.ToLower(DaftarFilm[indeksMinAtauMaks].Judul) {
						perluTukar = true
					}
				}
			case "tahunrilis":
				if urutanNaik {
					if DaftarFilm[j].TahunRilis < DaftarFilm[indeksMinAtauMaks].TahunRilis {
						perluTukar = true
					}
				} else {
					if DaftarFilm[j].TahunRilis > DaftarFilm[indeksMinAtauMaks].TahunRilis {
						perluTukar = true
					}
				}
			}
			if perluTukar {
				indeksMinAtauMaks = j
			}
		}
		DaftarFilm[i], DaftarFilm[indeksMinAtauMaks] = DaftarFilm[indeksMinAtauMaks], DaftarFilm[i]
	}
}

func insertionSort(kriteria string, urutanNaik bool) {
	for i := 1; i < JumlahFilm; i++ {
		kunci := DaftarFilm[i]
		j := i - 1
		for j >= 0 {
			perluPindah := false
			switch strings.ToLower(kriteria) {
			case "rating":
				if urutanNaik {
					if DaftarFilm[j].Rating > kunci.Rating {
						perluPindah = true
					}
				} else {
					if DaftarFilm[j].Rating < kunci.Rating {
						perluPindah = true
					}
				}
			case "genre":
				if urutanNaik {
					if strings.ToLower(DaftarFilm[j].Genre) > strings.ToLower(kunci.Genre) {
						perluPindah = true
					}
				} else {
					if strings.ToLower(DaftarFilm[j].Genre) < strings.ToLower(kunci.Genre) {
						perluPindah = true
					}
				}
			}
			if perluPindah {
				DaftarFilm[j+1] = DaftarFilm[j]
				j--
			} else {
				break
			}
		}
		DaftarFilm[j+1] = kunci
	}
}

func lihatDanUrutFilm() {
	spaceScreen()
	fmt.Println("--- DAFTAR FILM ---")
	garisPisah()
	if JumlahFilm == 0 {
		lihatDaftarFilm(DaftarFilm, JumlahFilm)
		enterLanjut()
		return
	}

	for {
		fmt.Println("\nOpsi Pengurutan:")
		fmt.Println("1. Urutkan berdasarkan Judul (Selection Sort)")
		fmt.Println("2. Urutkan berdasarkan Tahun Rilis (Selection Sort)")
		fmt.Println("3. Urutkan berdasarkan Rating (Insertion Sort)")
		fmt.Println("4. Urutkan berdasarkan Genre (Insertion Sort)")
		fmt.Println("0. Kembali ke Menu Utama")
		garisPisah()
		pilihanUrut := bacaInt("Pilih opsi pengurutan (0-4): ")

		if pilihanUrut == 0 {
			break
		}

		var urutanNaik bool
		fmt.Print("Urutkan secara Naik (a) atau Turun (d)? (a/d): ")
		urutanStr := strings.ToLower(bacaString(""))
		if urutanStr == "a" {
			urutanNaik = true
		} else if urutanStr == "d" {
			urutanNaik = false
		} else {
			fmt.Println("Pilihan urutan tidak valid. Default: Naik (Ascending).")
			urutanNaik = true
		}

		switch pilihanUrut {
		case 1:
			selectionSort("judul", urutanNaik)
			fmt.Println("\nDaftar Film setelah diurutkan berdasarkan Judul:")
		case 2:
			selectionSort("tahunrilis", urutanNaik)
			fmt.Println("\nDaftar Film setelah diurutkan berdasarkan Tahun Rilis:")
		case 3:
			insertionSort("rating", urutanNaik)
			fmt.Println("\nDaftar Film setelah diurutkan berdasarkan Rating:")
		case 4:
			insertionSort("genre", urutanNaik)
			fmt.Println("\nDaftar Film setelah diurutkan berdasarkan Genre:")
		default:
			fmt.Println("Pilihan tidak valid.")
			enterLanjut()
			continue
		}
		lihatDaftarFilm(DaftarFilm, JumlahFilm)
		enterLanjut()
		spaceScreen()
		fmt.Println("--- DAFTAR FILM ---")
		lihatDaftarFilm(DaftarFilm, JumlahFilm)
	}
}

func updateFilm() {
	spaceScreen()
	fmt.Println("--- UBAH DATA FILM ---")
	garisPisah()
	id := bacaInt("Masukkan ID Film yang ingin diubah: ")
	indeksFilm := cariBinerByID(id)
	if indeksFilm == -1 {
		fmt.Println("Film dengan ID tersebut tidak ditemukan.")
		enterLanjut()
		return
	}
	fmt.Printf("Film ditemukan: %s (ID: %d)\n", DaftarFilm[indeksFilm].Judul, DaftarFilm[indeksFilm].ID)
	fmt.Println("Silakan masukkan data baru (kosongkan jika tidak ingin mengubah):")

	newJudul := bacaString(fmt.Sprintf("Judul Film (sebelumnya: %s): ", DaftarFilm[indeksFilm].Judul))
	if newJudul != "" {
		DaftarFilm[indeksFilm].Judul = newJudul
	}

	newSutradara := bacaString(fmt.Sprintf("Sutradara (sebelumnya: %s): ", DaftarFilm[indeksFilm].Sutradara))
	if newSutradara != "" && newSutradara != DaftarFilm[indeksFilm].Sutradara {
		indeksSutradaraLama := cariSutradara(DaftarFilm[indeksFilm].Sutradara)
		if indeksSutradaraLama != -1 {
			DaftarSutradara[indeksSutradaraLama].JumlahFilm--
		}
		DaftarFilm[indeksFilm].Sutradara = newSutradara
		tambahSutradaraBaru(newSutradara)
	} else if newSutradara != "" {
		DaftarFilm[indeksFilm].Sutradara = newSutradara
	}

	newGenre := bacaString(fmt.Sprintf("Genre (sebelumnya: %s): ", DaftarFilm[indeksFilm].Genre))
	if newGenre != "" && newGenre != DaftarFilm[indeksFilm].Genre {
		indeksGenreLama := cariGenre(DaftarFilm[indeksFilm].Genre)
		if indeksGenreLama != -1 {
			DaftarKategoriGenre[indeksGenreLama].JumlahFilm--
		}
		DaftarFilm[indeksFilm].Genre = newGenre
		tambahGenreBaru(newGenre)
	} else if newGenre != "" {
		DaftarFilm[indeksFilm].Genre = newGenre
	}

	newTahunRilisStr := bacaString(fmt.Sprintf("Tahun Rilis (sebelumnya: %d): ", DaftarFilm[indeksFilm].TahunRilis))
	if newTahunRilisStr != "" {
		newTahunRilis, err := strconv.Atoi(strings.TrimSpace(newTahunRilisStr))
		if err == nil {
			DaftarFilm[indeksFilm].TahunRilis = newTahunRilis
		} else {
			fmt.Println("Input Tahun Rilis tidak valid. Menggunakan nilai sebelumnya.")
		}
	}

	newRatingStr := bacaString(fmt.Sprintf("Rating (sebelumnya: %.1f): ", DaftarFilm[indeksFilm].Rating))
	if newRatingStr != "" {
		newRating, err := strconv.ParseFloat(strings.TrimSpace(newRatingStr), 64)
		if err == nil && newRating >= 1.0 && newRating <= 5.0 {
			DaftarFilm[indeksFilm].Rating = newRating
		} else {
			fmt.Println("Input Rating tidak valid (harus 1.0-5.0). Menggunakan nilai sebelumnya.")
		}
	}

	newStatusTonton := strings.ToLower(bacaString(fmt.Sprintf("Status Tonton (sebelumnya: %s): ", DaftarFilm[indeksFilm].StatusTonton)))
	if newStatusTonton != "" {
		if newStatusTonton == "belum ditonton" || newStatusTonton == "sedang ditonton" || newStatusTonton == "sudah ditonton" {
			DaftarFilm[indeksFilm].StatusTonton = newStatusTonton
		} else {
			fmt.Println("Status Tonton tidak valid. Menggunakan nilai sebelumnya.")
		}
	}
	fmt.Println("\nData film berhasil diubah!")
	enterLanjut()
}

func deleteFilm() {
	spaceScreen()
	fmt.Println("--- HAPUS FILM ---")
	garisPisah()
	id := bacaInt("Masukkan ID Film yang ingin dihapus: ")
	indeksFilm := cariBinerByID(id)

	if indeksFilm == -1 {
		fmt.Println("Film dengan ID tersebut tidak ditemukan.")
	} else {
		fmt.Printf("Film '%s' (ID: %d) akan dihapus.\n", DaftarFilm[indeksFilm].Judul, DaftarFilm[indeksFilm].ID)
		konfirmasi := strings.ToLower(bacaString("Anda yakin ingin menghapus film ini? (y/n): "))
		if konfirmasi == "y" {
			indeksSutradara := cariSutradara(DaftarFilm[indeksFilm].Sutradara)
			if indeksSutradara != -1 {
				DaftarSutradara[indeksSutradara].JumlahFilm--
			}
			indeksGenre := cariGenre(DaftarFilm[indeksFilm].Genre)
			if indeksGenre != -1 {
				DaftarKategoriGenre[indeksGenre].JumlahFilm--
			}

			for i := indeksFilm; i < JumlahFilm-1; i++ {
				DaftarFilm[i] = DaftarFilm[i+1]
			}
			DaftarFilm[JumlahFilm-1] = TipeFilm{}
			JumlahFilm--
			fmt.Println("Film berhasil dihapus.")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	}
	enterLanjut()
}

func ubahStatusFilm() {
	spaceScreen()
	fmt.Println("--- UBAH STATUS TONTON FILM ---")
	garisPisah()
	id := bacaInt("Masukkan ID Film yang ingin diubah status tontonnya: ")
	indeksFilm := cariBinerByID(id)
	if indeksFilm == -1 {
		fmt.Println("Film dengan ID tersebut tidak ditemukan.")
		enterLanjut()
		return
	}

	fmt.Printf("Film '%s' (ID: %d) saat ini berstatus: %s\n", DaftarFilm[indeksFilm].Judul, DaftarFilm[indeksFilm].ID, DaftarFilm[indeksFilm].StatusTonton)
	for {
		newStatus := strings.ToLower(bacaString("Masukkan status baru (belum ditonton/sedang ditonton/sudah ditonton): "))
		if newStatus == "belum ditonton" || newStatus == "sedang ditonton" || newStatus == "sudah ditonton" {
			DaftarFilm[indeksFilm].StatusTonton = newStatus
			fmt.Println("Status tonton film berhasil diubah.")
			break
		}
		fmt.Println("Status tidak valid. Harap masukkan 'belum ditonton', 'sedang ditonton', atau 'sudah ditonton'.")
	}
	enterLanjut()
}

func lihatDaftarSutradara() {
	spaceScreen()
	fmt.Println("--- DAFTAR SUTRADARA ---")
	garisPisah()
	if JumlahSutradara == 0 {
		fmt.Println("Belum ada sutradara dalam daftar.")
		enterLanjut()
		return
	}
	fmt.Printf("%-5s | %-30s | %-15s\n", "ID", "Nama Sutradara", "Jumlah Film")
	garisPisah()
	for i := 0; i < JumlahSutradara; i++ {
		fmt.Printf("%-5d | %-30s | %-15d\n", DaftarSutradara[i].IDSutradara, DaftarSutradara[i].Nama, DaftarSutradara[i].JumlahFilm)
	}
	garisPisah()
	enterLanjut()
}

func lihatDaftarGenre() {
	spaceScreen()
	fmt.Println("--- DAFTAR KATEGORI GENRE ---")
	garisPisah()
	if JumlahKategoriGenre == 0 {
		fmt.Println("Belum ada kategori genre dalam daftar.")
		enterLanjut()
		return
	}
	fmt.Printf("%-5s | %-20s | %-15s\n", "ID", "Nama Genre", "Jumlah Film")
	garisPisah()
	for i := 0; i < JumlahKategoriGenre; i++ {
		fmt.Printf("%-5d | %-20s | %-15d\n", DaftarKategoriGenre[i].IDGenre, DaftarKategoriGenre[i].Nama, DaftarKategoriGenre[i].JumlahFilm)
	}
	garisPisah()
	enterLanjut()
}

func main() {
	fmt.Println("Memuat contoh data film...")
	tambahFilmDenganData("Inception", "Christopher Nolan", "Sci-Fi", 2010, 4.8, "sudah ditonton")
	tambahFilmDenganData("Interstellar", "Christopher Nolan", "Sci-Fi", 2014, 4.7, "sudah ditonton")
	tambahFilmDenganData("Dune: Part Two", "Denis Villeneuve", "Sci-Fi", 2024, 4.9, "belum ditonton")
	tambahFilmDenganData("Spirited Away", "Hayao Miyazaki", "Animasi", 2001, 4.9, "sudah ditonton")
	tambahFilmDenganData("Parasite", "Bong Joon-ho", "Thriller", 2019, 4.6, "sudah ditonton")
	tambahFilmDenganData("Forrest Gump", "Robert Zemeckis", "Drama", 1994, 4.7, "sudah ditonton")
	tambahFilmDenganData("Arrival", "Denis Villeneuve", "Sci-Fi", 2016, 4.5, "sedang ditonton")
	fmt.Println("Contoh data selesai dimuat. Tekan Enter untuk melanjutkan.")
	enterLanjut()

	for {
		spaceScreen()
		fmt.Println("=========================================")
		fmt.Println("  MY MOVIE TRACKER (MMT) - Tugas Besar Alpro")
		fmt.Println("=========================================")
		fmt.Println("1. Tambah Film Baru")
		fmt.Println("2. Lihat Semua Film (dan Urutkan)")
		fmt.Println("3. Cari Film")
		fmt.Println("4. Ubah Data Film")
		fmt.Println("5. Hapus Film")
		fmt.Println("6. Ubah Status Tonton Film")
		fmt.Println("7. Lihat Daftar Sutradara")
		fmt.Println("8. Lihat Daftar Kategori Genre")
		fmt.Println("9. Keluar")
		fmt.Println("=========================================")
		pilihanMenuUtama := bacaInt("Pilih menu (1-9): ")
		switch pilihanMenuUtama {
		case 1:
			tambahFilm()
		case 2:
			lihatDanUrutFilm()
		case 3:
			prosedurCariFilm()
		case 4:
			updateFilm()
		case 5:
			deleteFilm()
		case 6:
			ubahStatusFilm()
		case 7:
			lihatDaftarSutradara()
		case 8:
			lihatDaftarGenre()
		case 9:
			fmt.Println("Terima kasih telah menggunakan My Movie Tracker. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			enterLanjut()
		}
	}
}

func tambahFilmDenganData(judul, sutradara, genre string, tahun int, rating float64, status string) {
	if JumlahFilm >= MAKS_FILM {
		return
	}
	filmBaru := TipeFilm{
		ID:           generateNewID(),
		Judul:        judul,
		Sutradara:    sutradara,
		Genre:        genre,
		TahunRilis:   tahun,
		Rating:       rating,
		StatusTonton: status,
	}
	tambahSutradaraBaru(filmBaru.Sutradara)
	tambahGenreBaru(filmBaru.Genre)
	DaftarFilm[JumlahFilm] = filmBaru
	JumlahFilm++
}