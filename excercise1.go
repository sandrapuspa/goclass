package main

import (
    "fmt"
    "os"
    "strconv"
    "flag"
)

var biodata = []struct {
        nama string
        alamat string
        pekerjaan string
        alasan string
    }{
        {"andi","cilegon","kuli","alasan1"},
        {"ojan","depok","pns","alasan2"},
        {"patrick","jaksel","mahasiswa","alasan3"},
        {"yoga","jaksel","programmer","alasan4"},
    }

func main() {

   var argsRaw = os.Args[1]
    argg,_ := strconv.Atoi(argsRaw)
    b := biodata[argg]

    panggilbiodata(b.nama,b.alamat,b.pekerjaan,b.alasan)
}

func panggilbiodata(nama string, alamat string, pekerjaan string, alasan string) {
    fmt.Println("nama  : ", nama)
    fmt.Println("alamat  : ", alamat)
    fmt.Println("pekerjaan  : ", pekerjaan)
    fmt.Println("alasan  : ", alasan)
    fmt.Println()
    fmt.Println()
}