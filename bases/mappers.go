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
			toModelMethod(),
			newLine(),
			toDtoMethod(),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func toModelMethod() string {
	domain := strings.ToLower(Domain)
	signature := "static " + Domain + " toModel(" + Domain + "Dto " + domain + ")"
	return method(signature, "//TODO implement me")
}

func toDtoMethod() string {
	domain := strings.ToLower(Domain)
	signature := "static " + Domain + "Dto" + " toDto(" + Domain + " " + domain + ")"
	return method(signature, "//TODO implement me")
}
