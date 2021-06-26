package bases

import "strings"

const dependency = "repository"

func GetServiceBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("services"),
			generatedRepository().Import,
			generatedEntity().Import,
			newLine(),
			service().Import,
			allArgsConstructor().Import,
			entityNotFoundException().Import,
			newLine(),
			list().Import,
		},
		[]string{
			service().Content,
			allArgsConstructor().Content,
			startClass(Domain + "Service"),
			newLine(),
			generatedRepository().Content,
			newLine(),
			methodCreate(),
			newLine(),
			newLine(),
			methodGetAll(list().Content),
			newLine(),
			newLine(),
			methodGetById(entityNotFoundException().Content),
			newLine(),
			newLine(),
			methodEdit(),
			newLine(),
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

func methodDelete() string {
	serviceSignature := "void delete(Long id)"
	repositoryMethod := dependency + ".deleteById(id)"
	return method(serviceSignature, repositoryMethod)
}

func methodEdit() string {
	domain := strings.ToLower(Domain)

	serviceSignature := Domain + " edit(" + Domain + " " + domain + ")"
	repositoryMethod := "return " + dependency + ".save(" + domain + ")"
	return method(serviceSignature, repositoryMethod)
}

func methodGetById(exception string) string {
	serviceSignature := Domain + " getById(Long id)"
	repositoryMethod := "return " + dependency + ".findById(id).orElseThrow(" + exception + "::new)"
	return method(serviceSignature, repositoryMethod)
}

func methodCreate() string {
	domain := strings.ToLower(Domain)

	serviceSignature := Domain + " create(" + Domain + " " + domain + ")"
	repositoryMethod := "return " + dependency + ".save(" + domain + ")"
	return method(serviceSignature, repositoryMethod)
}

func methodGetAll(collection string) string {
	returnType := collection + "<" + Domain + ">"
	serviceSignature := returnType + " getAll()"
	repositoryMethod := "return (" + list().Content + "<" + Domain + ">" + ") " + dependency + ".findAll()"
	return method(serviceSignature, repositoryMethod)
}

func generatedService() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".services." + Domain + "Service;\n",
		tab() + "private final " + Domain + "Service service;\n",
	}
}
