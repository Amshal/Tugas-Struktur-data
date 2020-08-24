
package main

import (
	"fmt"
)

type Stack struct {
    slice []int
}


func (s *Stack) Push(i int) {
   s.slice = append(s.slice, i) 
} 


func (s *Stack) Peek() int {
 return s.slice[len(s.slice)-1]
}


func (s *Stack) Pop() int {
 var ret int = s.Peek()
 s.slice = s.slice[0:len(s.slice)-1]	
 return ret
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func main () {
 var s *Stack = new(Stack)
 var jumlah,index,rak,total int
 var pilih string

 fmt.Println()
 fmt.Println()
 fmt.Println("---------------------------------------------------------------------------")
 fmt.Println("SELAMAT MENGGUNAKAN APLIKASI MONITORING JUMLAH BUKU DI PERPUSTAKAAN SEKOLAH")
 fmt.Println("===========================================================================")
 fmt.Println("                              Silahkan Masukkan Data                      ")
 fmt.Println("---------------------------------------------------------------------------")
 fmt.Println()
 fmt.Print("Berapa jumlah rak buku yang ada di Perpustakaan Sekolah ?")
 fmt.Scanln(&rak)
index = 1
 for index != rak + 1 {
	fmt.Print("Masukkan Jumlah Buku di Rak ke ")
	fmt.Print(index)
	fmt.Print(" : ")
	fmt.Scanln(&jumlah)
	s.Push(jumlah)
	


	fmt.Print("Apakah data ingin dihapus Y/N :")
	fmt.Scanln(&pilih)

	fmt.Println()
	fmt.Println()
	if pilih == "Y" {
		s.Pop()
		fmt.Println("Data sudah terhapus")

		fmt.Print("Masukkan Ulang Jumlah Buku di Rak ke ")
		fmt.Print(index)
		fmt.Print(" : ")
		fmt.Scanln(&jumlah)
		s.Push(jumlah)
	} 
	index = index + 1
    total = total + jumlah	
}


 fmt.Println("Terimakasih telah menginput data")
 fmt.Println("Jumlah Buku di rak :")
 index = index - 1
 fmt.Println()
 fmt.Println("------------------------------")
 for index >= 1 {
	fmt.Print("| Rak ke : ")
	fmt.Print(index)
	fmt.Print("       | ") 
	fmt.Print(s.Pop())
	fmt.Println(" Buku  |")
	fmt.Println("------------------------------")
 index = index - 1
 }
fmt.Print("| Total Jumlah Buku yang Ada di Rak | ")
fmt.Print(total)
fmt.Println(" Buku |")
fmt.Println("------------------------------")
fmt.Println()


}