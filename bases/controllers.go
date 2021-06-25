package bases

func GetControllerBody() string {
	content := Content{
		[]string{
			packageDeclarationLine("controllers"),
			restController().Import,
			requestMapping().Import,
		},
		[]string{
			restController().Content,
			requestMapping().Content,
			startClass(Domain + "Controller"),
			endClass(),
		},
	}
	return concatenateBody(content)
}
