package main
 
import "fmt"
 
func main(){
    fmt.Print("TASK 2: Menampilkan angka prima start : 3 (10 angka) = ")
    hitungprima(3,35)
       
}

func hitungprima(awal int, akhir int) {
    for i:=awal;i<=akhir;i++ {
        a := 0
        for j:=1;j<=i;j++ {
            if(i % j == 0){
                a++
            }
        }
        if(a==2){
            fmt.Print(" ",i)
        }
    }
}
