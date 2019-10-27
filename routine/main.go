package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

type Statuses struct {
	Statuses Status `json:"status"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gtpl"))
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func status(w http.ResponseWriter, r *http.Request) {

	// Open our jsonFile
	jsonFile, err := os.Open("data/status.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened status.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var statuses Statuses

	json.Unmarshal([]byte(byteValue), &statuses)

	errs := tpl.ExecuteTemplate(w, "base.gtpl", statuses.Statuses)
	if errs != nil {
		log.Fatalln(errs)
	}
}

func main() {

	go startwrite()
	http.HandleFunc("/", status)             // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}

func writejson() {
	dt := time.Now()
	data := Statuses{
		Status{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	err := ioutil.WriteFile("data/status.json", file, 0755)
	if err != nil {
		fmt.Println("Error writing Files: ", err)
	} else {
		fmt.Println("Success Write File: ", dt.Format("2006.01.02 15:04:05"))
	}
}
func startwrite() {
	for {
		<-time.After(15 * time.Second)
		go writejson()
	}
}
