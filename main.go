package main

import (
	"log"
	"os"
	"strings"

	"pl.krukkrz/spring-mvc-generator/bases"
)

const fileExtention = ".java"

func main() {
	log.Println("Staring the app")

	log.Println("Clean /target directory..")
	os.RemoveAll("./target")

	domain := readDomain()
	createEntity(domain)
	createDto(domain)
	//createMapper(domain)
	//createService(domain)
	//createController(domain)
}

func readDomain() string {
	domain := os.Args[1]
	return strings.Title(domain)
}

func createEntity(domain string) {
	log.Println("Creating entitiy " + domain)
	path := basePath(domain) + "/models/entities/"
	createJavaFile(path, domain, bases.GetEntityBody(domain))
}

func createDto(domain string) {
	log.Println("Creating dto " + domain)
	path := basePath(domain) + "/models/dtos/"
	fileName := domain + "Dto"
	createJavaFile(path, fileName, bases.GetDtoBody(fileName))
}

func createJavaFile(path string, fileName string, content string) {
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
	f, err := os.Create(path + fileName + fileExtention)
	check(err)
	defer f.Close()

	log.Println("writing a content of class " + fileName + fileExtention)
	_, err = f.WriteString(content)
	check(err)
	f.Sync()
}

func basePath(domain string) string {
	return "./target/" + strings.ToLower(domain)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
