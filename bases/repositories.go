package bases

import "strings"

func GetRepositoryBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("repositories"),
			generatedEntity().Import,
			newLine(),
			repository().Import,
			crudRepository().Import,
		},
		[]string{
			repository().Content,
			startInterfaceWithExtention(Domain+"Repository", crudRepository().Content+"<"+Domain+", Long"+">"),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func generatedRepository() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".repositories." + Domain + "Repository;\n",
		tab() + "private final " + Domain + "Repository repository;\n",
	}
}
