package bases

var httpStatusNotFound = "import static org.springframework.http.HttpStatus.NOT_FOUND;\n"

func GetExceptionControllerBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("controllers"),
			controllerAdvice().Import,
			responseStatus("").Import,
			exceptionHandler("").Import,
			newLine(),
			entityNotFoundException().Import,
			newLine(),
			httpStatusNotFound,
		},
		[]string{
			controllerAdvice().Content,
			newLine(),
			startClass("ExceptionController"),

			newLine(),
			tab(),
			responseStatus("value = NOT_FOUND, reason = \"No such entity\"").Content,
			newLine(),
			tab(),
			exceptionHandler(entityNotFoundException().Content + ".class").Content,
			newLine(),
			notFoundHandlerMethod(),
			endClass(),
		},
	}
	return concatenateBody(content)
}

func notFoundHandlerMethod() string {
	signature := "void notFound()"
	return method(signature)
}
