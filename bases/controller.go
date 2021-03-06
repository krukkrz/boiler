package bases

import "strings"

func GetControllerBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("controllers"),
			generatedService().Import,
			generatedDto().Import,
			generatedEntity().Import,
			importMapper(),
			staticToDto().Import,
			staticToModel().Import,

			newLine(),
			restController().Import,
			requestMapping().Import,
			allArgsConstructor().Import,
			postMapping("").Import,
			getMapping("").Import,
			putMapping("").Import,
			deleteMapping("").Import,
			requestBody().Import,
			responseEntity().Import,
			responseStatus("").Import,
			pathVariable("").Import,
			"import static org.springframework.http.HttpStatus.OK;\n",

			newLine(),
			list().Import,
			collectors().Import,
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
			newLine(),

			newLine(),
			tab(),
			getMapping("").Content,
			newLine(),
			getAllMethod(),
			newLine(),

			newLine(),
			tab(),
			getMapping("\"/{id}\"").Content,
			newLine(),
			getByIdMethod(),
			newLine(),

			newLine(),
			tab(),
			putMapping("").Content,
			newLine(),
			editMethod(),
			newLine(),

			newLine(),
			tab(),
			deleteMapping("\"/{id}\"").Content,
			newLine(),
			tab(),
			responseStatus("OK").Content,
			newLine(),
			deleteMethod(),

			endClass(),
		},
	}
	return concatenateBody(content)
}

func deleteMethod() string {
	signature := "void delete(" + pathVariable("\"id\"").Content + "Long id)"
	serviceCall := "service.delete(id)"
	return method(signature, serviceCall)
}

func editMethod() string {
	domain := strings.ToLower(Domain)
	signature := responseEntity().Content + "<" + generatedDto().Content + "> edit(" + requestBody().Content + " " + generatedDto().Content + " " + domain + "Dto)"
	toModel := generatedEntity().Content + " " + domain + " = " + staticToModel().Content + "(" + domain + "Dto)"
	savedIteration := generatedEntity().Content + " saved" + Domain + " = service.edit(" + domain + ")"
	returnMethod := "return " + responseEntity().Content + ".ok(" + staticToDto().Content + "(saved" + Domain + "))"
	return method(signature, toModel, savedIteration, returnMethod)
}

func getByIdMethod() string {
	signature := responseEntity().Content + "<" + generatedDto().Content + "> getById(" + pathVariable("\"id\"").Content + " Long id)"
	serviceCall1 := "return " + responseEntity().Content + ".ok("
	serviceCall2 := newLine() + tab() + tab() + tab() + "toDto(service.getById(id))"
	serviceCall3 := newLine() + tab() + tab() + ")"
	serviceCall := serviceCall1 + serviceCall2 + serviceCall3
	return method(signature, serviceCall)
}

func getAllMethod() string {
	signature := responseEntity().Content + "<" + list().Content + "<" + generatedDto().Content + ">> getAll()"
	returnLine := "return " + responseEntity().Content + ".ok("
	serviceStream1 := newLine() + tab() + tab() + tab() + "service.getAll()"
	serviceStream2 := newLine() + tab() + tab() + tab() + tab() + ".stream()"
	serviceStream3 := newLine() + tab() + tab() + tab() + tab() + ".map(" + Domain + "Mapper::" + staticToDto().Content + ")"
	serviceStream4 := newLine() + tab() + tab() + tab() + tab() + ".collect(" + collectors().Content + ".toList()))"
	serviceStream := returnLine + serviceStream1 + serviceStream2 + serviceStream3 + serviceStream4
	return method(signature, serviceStream)
}

func createMethod() string {
	domain := strings.ToLower(Domain)
	signature := responseEntity().Content + "<" + generatedDto().Content + "> create(" + requestBody().Content + " " + generatedDto().Content + " " + domain + "Dto)"
	toModel := generatedEntity().Content + " " + domain + " = " + staticToModel().Content + "(" + domain + "Dto)"
	savedIteration := generatedEntity().Content + " saved" + Domain + " = service.create(" + domain + ")"
	returnMethod := "return " + responseEntity().Content + ".ok(" + staticToDto().Content + "(saved" + Domain + "))"
	return method(signature, toModel, savedIteration, returnMethod)
}
