package bases

import "strings"

var Domain string

var Package string

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
	return startClass(name) + body + endClass()
}

func startClass(name string) string {
	return "public class " + name + "{ \n\n"
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
