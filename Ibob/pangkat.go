package main
import "fmt"

// return list of primes less than N

func main() {
    for start := 1; start <=10; start++{
        fmt.Print(start*start)
        fmt.Print(", ")
    }
}