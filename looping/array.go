package main

import (
	"fmt"
	"sort"
	"time"
)

var data = [...]string{"0001", "Roman Alamsyah", "Bandar Lampung", "21/05/1998", "Membaca"}
var dateValue = "21 May 1998"
var date time.Time

var NewDate = [...]string{"21", "05", "1998"}
var value string = NewDate[1]

func dataHandling2(add []string) []string {
	var input = data
	fmt.Println(input)

	add = append(data[0:1], add...)
	return add
}

func sortingDate() string {
	switch value {
	case NewDate[0]:
		fmt.Println(21)
	case NewDate[1]:
		fmt.Println("Mei")
	case NewDate[2]:
		fmt.Println(1998)
	}

	fmt.Println("===============================")

	return value
}
func sortDate() {
	strs := NewDate
	sort.Strings(strs[:])
	fmt.Println(strs)

}
func main() {
	fmt.Println(dataHandling2([]string{"Roman Alamsyah Elsharawy", "Provinsi Bandar Lampung", "21/05/1998", "Pria", "SMA Internasional Metro"}))
	fmt.Println(sortingDate())
	fmt.Println("=======================================")
	sortDate()
}
