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
			getMapping("").Import,
			requestBody().Import,
			responseEntity().Import,
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

			endClass(),
		},
	}
	return concatenateBody(content)
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
	signature := responseEntity().Content + "<" + generatedEntity().Content + "> create(" + requestBody().Content + " " + generatedDto().Content + " " + domain + "Dto)"
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

func collectors() Class {
	return Class{
		"import java.util.stream.Collectors;\n",
		"Collectors",
	}
}
