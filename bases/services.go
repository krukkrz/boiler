package bases

import "strings"

func GetServiceBody(domain string) string {
	content := Content{
		[]string{
			packageDeclarationLine("services"),
			importRepository().Import,
			service().Import,
			allArgsConstructor().Import,
			entityNotFoundException().Import,
			list().Import,
		},
		[]string{
			service().Content,
			allArgsConstructor().Content,
			startClass(domain + "Service"),
			importRepository().Content,
			methodCreate(),
			newLine(),
			methodGetAll(list().Content),
			newLine(),
			methodGetById(entityNotFoundException().Content),
			newLine(),
			methodEdit(),
			newLine(),
			methodDelete(),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func list() Class {
	return Class{
		"import java.util.List;\n",
		"List",
	}
}

func entityNotFoundException() Class {
	return Class{
		"import javax.persistence.EntityNotFoundException;\n",
		"EntityNotFoundException",
	}
}

func importRepository() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".repository." + Domain + "Repository;\n",
		tab() + "private final " + Domain + "Repository repository;\n",
	}
}

func methodDelete() string {
	serviceSignature := "void delete(Long id)"
	repositoryMethod := "deleteById(id)"
	return method(serviceSignature, repositoryMethod)
}

func methodEdit() string {
	domain := strings.ToLower(Domain)

	serviceSignature := Domain + " edit(" + Domain + " " + domain + ")"
	repositoryMethod := "save(" + domain + ")"
	return method(serviceSignature, repositoryMethod)
}

func methodGetById(exception string) string {
	serviceSignature := Domain + " getById(Long id)"
	repositoryMethod := "findById(id).orElseThrow(" + exception + "::new)"
	return method(serviceSignature, repositoryMethod)
}

func methodCreate() string {
	domain := strings.ToLower(Domain)

	serviceSignature := Domain + " create(" + Domain + " " + domain + ")"
	repositoryMethod := "save(" + domain + ")"
	return method(serviceSignature, repositoryMethod)
}

func methodGetAll(collection string) string {
	returnType := collection + "<" + Domain + ">"
	serviceSignature := returnType + " getAll()"
	repositoryMethod := "findAll()"
	return method(serviceSignature, repositoryMethod)
}

func method(serviceMethod string, repositoryMethod string) string {
	method := []string{
		newLine(),
		tab(),
		"public " + serviceMethod + " {",
		newLine(),
		tab(),
		tab(),
		"repository." + repositoryMethod + ";",
		newLine(),
		tab(),
		"}",
	}
	return concatenateStrings(method)
}
