package main

import (
    "fmt"
	"os"
)

type Student struct {
	id string
    name string
    address string
	occupation string
	reason string
}


func main() {
	/* Showing Selected Data */
	input := os.Args[1]
	//input := 2

	var students []Student
	students = append(students, Student{"1", "Djodi Ramadhan", "Jatiwaringin", "SE", "Training"})
	students = append(students, Student{"2", "Yoga Budhi Prawira", "Jatiwaringin", "SE", "Training"})
	
	PrintSelectedStudent(students, input)
	
	/* Appending Slice of Struct */
	student_data := os.Args[2:]
	students = append(students, Student{student_data[0], student_data[1], student_data[2], student_data[3], student_data[4]})
	fmt.Println(students);
}

func PrintSelectedStudent(students []Student, input string){
	for _, v := range students {
		if v.id == input{
			fmt.Println("ID : ", v.id)
			fmt.Println("Name : ", v.name)
			fmt.Println("Address : ", v.address)
			fmt.Println("Occupation : ", v.occupation)
			fmt.Println("Reason : ", v.reason)
		}
	}
}