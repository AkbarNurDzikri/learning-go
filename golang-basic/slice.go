package main

import "fmt"

func main() {
	names := [...]string{
		"Dzikri",
		"Fauzan",
		"Rizki",
		"Alif",
		"Fajar",
		"Rizal",
		"Abdul",
		"Udin",
		"Abud",
		"Fadli",
	}

	// slice sama dengan array, hanya tidak ditentukan ukurannya
	var slice1 []string = names[4:7]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[5:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// memeriksa kapasitas slice
	fmt.Println(cap(slice4))

	// memeriksa panjang slice
	fmt.Println(len(slice4))

	// membuat slice baru dengan menambahkan data baru ke posisi terakhir slice
	slice5 := append(slice4, "Eko")
	fmt.Println(slice5)

	// membuat slice baru dari array baru

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)

	// membuat slice langsung
	var slice []string = make([]string, 2, 5)
	slice[0] = "Dzikri"
	slice[1] = "Akbar"
	// slice[2] = "Akbar" // akan error karena kapasitas slice hanya 2

	fmt.Println(slice)

	// menambahkan data ke slice dengan append
	var newSlice []string = append(slice, "Khalid")
	fmt.Println(newSlice)

	// membuktikan bahwa newSlice menggunakan array milik slice
	newSlice[0] = "Dzikri Baru"
	fmt.Println(newSlice)
	fmt.Println(slice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan membuat array dan slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
}