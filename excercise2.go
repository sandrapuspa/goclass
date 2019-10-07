package main

import (
    "fmt"
    "flag"
)

func main() {
    var namanya = flag.String("nama", "anonymous", "type your name")
    var alamatnya = flag.String("alamat", "alamatnya", "type your location")
    var pekerjaannya = flag.String("pekerjaan", "pekerjaannya", "type your profession")
    var alasannya = flag.String("alasan", "alasannya", "type your reason")
    
    flag.Parse()
    fmt.Printf("nama\t: %s\n", *namanya)
    fmt.Printf("alamat\t: %s\n", *alamatnya)
    fmt.Printf("pekerjaan\t: %s\n", *pekerjaannya)
    fmt.Printf("alasan\t: %s\n", *alasannya)
}

func panggilbiodata(nama string, alamat string, pekerjaan string, alasan string) {
    fmt.Println("nama  : ", nama)
    fmt.Println("alamat  : ", alamat)
    fmt.Println("pekerjaan  : ", pekerjaan)
    fmt.Println("alasan  : ", alasan)
    fmt.Println()
    fmt.Println()
}