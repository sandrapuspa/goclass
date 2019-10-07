package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type DataTeman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

var Data = []DataTeman{
	{"Yoga", "Jakarta", "Software Engineer"},
	{"Jodi", "Jakarta", "Web Developer"},
	{"Ebob", "Jakarta", "Database Admin"},
	{"Patrick", "Jakarta", "Mahasiswa"},
}

func main() {
	logrus.Info("Start Program")
	defer logrus.Info("Quit Program")
	l := len(os.Args)
	if l < 1 {
		logrus.Error("Argument Invalid!")
		return
	}
	if os.Args[1] == "add" {
		if len(os.Args) < 5 {
			logrus.Error("Argument Invalid!")
			return
		}
		dt := DataTeman{
			Nama:      os.Args[2],
			Alamat:    os.Args[3],
			Pekerjaan: os.Args[4],
		}
		Data = append(Data, dt)
		for _, val := range Data {
			logrus.Info("Nama: ", val.Nama)
			logrus.Info("Alamat: " + val.Alamat)
			logrus.Info("Pekerjaan: " + val.Pekerjaan)
		}
	} else {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil || i > len(Data) || i < 0 {
			logrus.Error("Argument Invalid!")
			return
		}
		logrus.Info("Nama: " + Data[i].Nama)
		logrus.Info("Alamat: " + Data[i].Alamat)
		logrus.Info("Pekerjaan: " + Data[i].Pekerjaan)
	}
}
