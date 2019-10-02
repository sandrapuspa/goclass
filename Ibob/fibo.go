package main
 
import (
    "fmt"
)

func main() {
    fmt.Println("testo")
    n1 := 5
    n2 := 7
    n3 := 0

    fmt.Print(n1)
    fmt.Print(", ")
    fmt.Print(n2)

    for i:=1; i<=9; i++{
        n3 = n1+ n2
        fmt.Print(", ")
        fmt.Print(n3)
        n1=n2
        n2=n3
    }
}