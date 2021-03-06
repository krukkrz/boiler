package bases

import "strings"

func GetEntityBody() string {
	classBody := "\t" + id().Content + "\t" + generatedValue().Content + "\tprivate Long id;\n\t//TODO implement parameters here"
	content := Content{
		[]string{
			packageDeclarationLine("models.entities"),
			id().Import,
			generatedValue().Import,
			entity().Import,
			builder().Import,
			data().Import,
			equalsAndHashCode().Import,
			allArgsConstructor().Import,
			noArgsConstructor().Import,
		},
		[]string{
			entity().Content,
			builder().Content,
			data().Content,
			equalsAndHashCode().Content,
			allArgsConstructor().Content,
			noArgsConstructor().Content,
			getClassSkeleton(Domain, classBody),
		},
	}
	return concatenateBody(content)
}

func GetDtoBody(name string) string {
	content := Content{
		[]string{
			packageDeclarationLine("models.dtos"),
			builder().Import,
			data().Import,
			noArgsConstructor().Import,
			allArgsConstructor().Import,
		},
		[]string{
			builder().Content,
			data().Content,
			noArgsConstructor().Content,
			allArgsConstructor().Content,
			getClassSkeleton(name, "\t//TODO implement parameters here"),
		},
	}
	return concatenateBody(content)
}

func generatedEntity() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".models.entities." + Domain + ";\n",
		Domain,
	}
}

func generatedDto() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import " + Package + "." + domain + ".models.dtos." + Domain + "Dto;\n",
		Domain + "Dto",
	}
}
