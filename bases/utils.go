package bases

import "strings"

var Domain string

var Package string

func method(serviceMethod string, dependencyMethod ...string) string {
	method := []string{
		tab(),
		"public " + serviceMethod + " {",
		newLine(),
		dependencyMethods(dependencyMethod...),
		tab(),
		"}",
	}
	return concatenateStrings(method)
}

func dependencyMethods(dependencyMethod ...string) string {
	var methods string
	for _, method := range dependencyMethod {
		methods = methods + tab() + tab() + method + ";" + newLine()
	}
	return methods
}

func concatenateBody(content Content) string {
	var imports string
	for _, item := range content.Imports {
		imports = imports + item
	}

	var body string
	for _, item := range content.Content {
		body = body + item
	}

	return imports + newLine() + body
}

func concatenateStrings(strings []string) string {
	var concatenated string
	for _, item := range strings {
		concatenated = concatenated + item
	}
	return concatenated
}

func packageDeclarationLine(location string) string {
	return "package " + packageLocation(location) + "; \n\n"
}

func packageLocation(location string) string {
	domain := strings.ToLower(Domain)
	return Package + "." + domain + "." + location
}

func getClassSkeleton(name string, body string) string {
	return startClass(name) + newLine() + body + endClass()
}

func startClass(name string) string {
	return "public class " + name + " { \n"
}

func startInterfaceWithExtention(name string, extends string) string {
	return "public interface " + name + " extends " + extends + " { \n"
}

func endClass() string {
	return closeBlock()
}

func closeBlock() string {
	return "\n}"
}

func tab() string {
	return "\t"
}

func newLine() string {
	return "\n"
}
