package library

import fmt

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