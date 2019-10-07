package patrick

import "fmt"

func Prima() {
	//2, 3, 5, 7, 11, 13, 17, 19, 23, 29

	prima := []int{2, 3, 5}
	number := 1
	for len(prima) <= 10 {
		for _, element := range prima {
			temp := number % element
			for i := 0; i < len(prima); i++ {
				if temp != 0 && temp != 1 && temp != prima[i] {
					prima = append(prima, temp)
				}
			}

		}
		number++
	}
	fmt.Println(prima)
	// for _, element := range prima {
	// 	fmt.Println(element)
	// }

}
