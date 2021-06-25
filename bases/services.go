package bases

import "strings"

func GetServiceBody(domain string) string {
	content := Content{
		[]string{
			packageLine("services"),
			service().Import,
			allArgsConstructor().Import,
		},
		[]string{
			service().Content,
			allArgsConstructor().Content,
			startClass(domain + "Service"),
			importRepository().Content,
			methodCreate(),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func importRepository() Class {
	return Class{
		"//TODO update it respectively to you directory import " + Domain + ".repository.IterationRepository;\n",
		tab() + "private final " + Domain + "Repository repository;\n",
	}
}

func methodCreate() string {
	domain := strings.ToLower(Domain)
	method := []string{
		newLine(),
		tab(),
		"public " + Domain + " create(" + Domain + " " + domain + ") {",
		newLine(),
		tab(),
		tab(),
		"repository.save(" + domain + ");",

		newLine(),
		tab(),
		"}",
	}
	return concatenateStrings(method)
}
