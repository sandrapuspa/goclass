package patrick

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func Exercise() {
	temanteman := []Teman{{"User1", "House1", "Dummy1", "Biar Pinter1"}, {"User2", "House2", "Dummy2", "Biar Pinter2"}, {"User3", "House3", "Dummy3", "Biar Pinter3"}}

	arg := os.Args[1:]

	var indexArg []int
	var insertArg []string
	var check int = -1

	for _, v := range arg {
		if check == -1 {
			var temp, err = strconv.Atoi(v)
			if err != nil {
				check = 0

			}
			if check == -1 {
				indexArg = append(indexArg, temp)
				fmt.Println(indexArg)
			}
		}

		if check != -1 {
			insertArg = append(insertArg, v)
			fmt.Println(insertArg)
			check = check + 1
			if check == 4 {
				check = -1
			}
		}
	}

	for i := 1; i < (len(insertArg)/4)+1; i++ {
		temp := Teman{insertArg[((4 * (i - 1)) + 0)], insertArg[(4*(i-1))+1], insertArg[(4*(i-1))+2], insertArg[(4*(i-1))+3]}

		temanteman = append(temanteman, temp)
	}

	for _, v := range indexArg {
		fmt.Println(temanteman[v-1])
	}

}
