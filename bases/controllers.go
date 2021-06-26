package bases

import "strings"

func GetControllerBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("controllers"),
			generatedService().Import,
			generatedDto().Import,
			generatedEntity().Import,
			staticToDto().Import,
			staticToModel().Import,

			newLine(),
			restController().Import,
			requestMapping().Import,
			allArgsConstructor().Import,
			postMapping("").Import,
			requestBody().Import,
			responseEntity().Import,
		},
		[]string{
			restController().Content,
			requestMapping().Content,
			newLine(),
			allArgsConstructor().Content,
			startClass(Domain + "Controller"),

			newLine(),
			generatedService().Content,

			newLine(),
			tab(),
			postMapping("").Content,
			newLine(),
			createMethod(),

			endClass(),
		},
	}
	return concatenateBody(content)
}

func createMethod() string {
	domain := strings.ToLower(Domain)
	signature := generatedEntity().Content + " create(" + requestBody().Content + " " + generatedDto().Content + " " + domain + "Dto)"
	toModel := generatedEntity().Content + " " + domain + " = " + staticToModel().Content + "(" + domain + "Dto)"
	savedIteration := generatedEntity().Content + " saved" + Domain + " = service.create(" + domain + ")"
	returnMethod := "return " + responseEntity().Content + ".ok(" + staticToDto().Content + "(saved" + Domain + "))"
	return method(signature, toModel, savedIteration, returnMethod)
}

func responseEntity() Class {
	return Class{
		"import org.springframework.http.ResponseEntity;\n",
		"ResponseEntity",
	}
}
