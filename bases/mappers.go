package bases

import "strings"

func GetMapperBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("models.mapping"),
			generatedEntity().Import,
			generatedDto().Import,
			newLine(),
			component().Import,
		},
		[]string{
			component().Content,
			startClass(Domain + "Mapper"),
			newLine(),
			toModelMethod(),
			newLine(),
			newLine(),
			toDtoMethod(),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func importMapper() string {
	domain := strings.ToLower(Domain)
	return "import " + Package + "." + domain + ".models.mapping." + Domain + "Mapper;\n"
}

func toModelMethod() string {
	domain := strings.ToLower(Domain)
	signature := "static " + Domain + " toModel(" + Domain + "Dto " + domain + ")"
	return method(signature, "return "+Domain+".builder().build()")
}

func toDtoMethod() string {
	domain := strings.ToLower(Domain)
	signature := "static " + Domain + "Dto" + " toDto(" + Domain + " " + domain + ")"
	return method(signature, "return "+Domain+"Dto.builder().build()")
}

func staticToModel() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import static " + Package + "." + domain + ".models.mapping." + Domain + "Mapper.toModel;\n",
		"toModel",
	}
}

func staticToDto() Class {
	domain := strings.ToLower(Domain)
	return Class{
		"import static " + Package + "." + domain + ".models.mapping." + Domain + "Mapper.toDto;\n",
		"toDto",
	}
}
