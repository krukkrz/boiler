package main

import (
	"log"
	"os"
	"strings"

	"pl.krukkrz/spring-mvc-generator/bases"
)

func main() {

	//jako argumenty:
	//- nazwa encji

	//tworzę folder z nazwą encji
	//w nim:
	// repositories
	// services
	// controllers
	// models
	//	entities
	// 	dto
	// mappers

	log.Println("Staring the app")

	domain := readDomain()

	createEntity(domain)

	// -------------------------------
	// os.Mkdir("./tmp", 0755)

	// f, err := os.Create("./tmp/dat2")
	// check(err)
	// defer f.Close()

	// n3, err := f.WriteString("writes\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n3)

	// f.Sync()
}

func readDomain() string {
	domain := os.Args[1]
	return strings.Title(domain)
}

func createEntity(domain string) {
	log.Println("Creating entitiy " + domain)

	path := "./target/" + strings.ToLower(domain) + "/models/entities/"

	err := os.MkdirAll(path, os.ModePerm)
	check(err)
	f, err := os.Create(path + domain + ".class")
	check(err)
	defer f.Close()

	log.Println("writing a content of s class " + domain + ".java")
	_, err = f.WriteString(bases.GetEntityBody(domain))
	check(err)
	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
