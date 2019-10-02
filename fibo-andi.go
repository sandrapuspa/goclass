package main
 
import "fmt"
 
func main(){
    var n int
    t1:=5
    t2:=7
    nextnya:=0
    n = 10
     
    fmt.Print("TASK: Barisan Fibonacci : ")
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
