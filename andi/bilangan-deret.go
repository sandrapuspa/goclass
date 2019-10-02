package main
 
import "fmt"
 
func main(){
    fmt.Print("TASK 3: Hitung deret : 3 (10 angka) = ")
    hitungderet(1,10,3)
       
}

func hitungderet(start int, n int, mp int) {
    var i int 
    count := 0
    next := start
    for i = start; i <= 190; i++ {

        if count == n {
            break
        }

        fmt.Print(next, " ")
        next = next + mp
        mp = mp + 2
        count = count + 1
    }

}