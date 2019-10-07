package library

import "fmt"

func HitungFibonacci(n int, t1 int, t2 int, nextnya int) {
    for i:=1;i<=n;i++ {
        if(i==1){
            fmt.Print(" ",t1)
            continue
        }
        if(i==2){
            fmt.Print(" ",t2)
            continue
        }
        nextnya = t1 + t2
        t1 = t2
        t2 = nextnya
        fmt.Print(" ",nextnya)
    }
}

func Hitungprima(awal int, akhir int) {
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

func Hitungderet(start int, n int, mp int) {
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