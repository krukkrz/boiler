package main

import (
	"log"
	"os"
	"strings"

	"pl.krukkrz/boiler/bases"
)

const fileExtention = ".java"

func main() {
	log.Println("Staring the app")

	log.Println("Clean /target directory..")
	os.RemoveAll("./target")

	domain := readDomain()
	bases.Domain = domain
	bases.Package = readPackage()

	createEntity(domain)
	createDto(domain)
	//createMapper(domain)
	createService(domain)
	//createController(domain)
}

func readDomain() string {
	domain := os.Args[1]
	return strings.Title(domain)
}

func readPackage() string {
	return os.Args[2]
}

func createService(domain string) {
	fileName := domain + "Service"
	log.Println("Creating service " + fileName)
	path := basePath(domain) + "/services/"
	createJavaFile(path, fileName, bases.GetServiceBody(domain))
}

func createEntity(domain string) {
	log.Println("Creating entitiy " + domain)
	path := basePath(domain) + "/models/entities/"
	createJavaFile(path, domain, bases.GetEntityBody(domain))
}

func createDto(domain string) {
	fileName := domain + "Dto"
	log.Println("Creating dto " + fileName)
	path := basePath(domain) + "/models/dtos/"
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
