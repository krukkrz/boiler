package bases

import "strings"

func GetRepositoryBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("repositories"),
			repository().Import,
			crudRepository().Import,
			generatedEntity().Import,
		},
		[]string{
			repository().Content,
			startClassWithExtention(Domain+"Repository", crudRepository().Content+"<"+Domain+", Long"+">"),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func crudRepository() Class {
	return Class{
		"import org.springframework.data.repository.CrudRepository;\n",
		"CrudRepository",
	}
}

func generatedRepository() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".repository." + Domain + "Repository;\n",
		tab() + "private final " + Domain + "Repository repository;\n",
	}
}
