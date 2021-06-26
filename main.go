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

	createEntity()
	createDto()
	createRepository()
	createMapper()
	createService()
	createController()
}

func readDomain() string {
	domain := os.Args[1]
	return strings.Title(domain)
}

func readPackage() string {
	return os.Args[2]
}

func createMapper() {
	fileName := bases.Domain + "Mapper"
	log.Println("Creating mapper " + fileName)
	path := basePath(bases.Domain) + "/models/mapping/"
	createJavaFile(path, fileName, bases.GetMapperBody())
}

func createController() {
	fileName := bases.Domain + "Controller"
	log.Println("Creating controller " + fileName)
	path := basePath(bases.Domain) + "/controllers/"
	createJavaFile(path, fileName, bases.GetControllerBody())
}

func createRepository() {
	fileName := bases.Domain + "Repository"
	log.Println("Creating repository " + fileName)
	path := basePath(bases.Domain) + "/repositories/"
	createJavaFile(path, fileName, bases.GetRepositoryBody())
}

func createService() {
	fileName := bases.Domain + "Service"
	log.Println("Creating service " + fileName)
	path := basePath(bases.Domain) + "/services/"
	createJavaFile(path, fileName, bases.GetServiceBody())
}

func createEntity() {
	log.Println("Creating entitiy " + bases.Domain)
	path := basePath(bases.Domain) + "/models/entities/"
	createJavaFile(path, bases.Domain, bases.GetEntityBody())
}

func createDto() {
	fileName := bases.Domain + "Dto"
	log.Println("Creating dto " + fileName)
	path := basePath(bases.Domain) + "/models/dtos/"
	createJavaFile(path, fileName, bases.GetDtoBody(fileName))
}

func createJavaFile(path string, fileName string, content string) {
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
	f, err := os.Create(path + fileName + fileExtention)
	check(err)
	defer f.Close()

	log.Println("Writing a content of a class " + fileName + fileExtention)
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
