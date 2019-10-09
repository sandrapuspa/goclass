package main
import "fmt"

// return list of primes less than N
func prima(N int) (primes []int) {
    b := make([]bool, N)
    for i := 2; i < N; i++ {
        if b[i] == true { continue }
        primes = append(primes, i)
        for x2 := i * i; x2 < N; x2 += i {
            b[x2] = true
        }
    }
    return
}

func main() {
    primes := prima(30)
    for _, p := range primes {
        fmt.Print(p,", ")
    }
}