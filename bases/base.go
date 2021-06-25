package bases

type Annotation struct {
	Import  string
	Content string
}

type Content struct {
	Imports []string
	Content []string
}

func GetEntityBody(domain string) string {
	classBody := "\t" + id().Content + "\t" + generatedValue().Content + "\tprivate Long id;\n\t//TODO implement parameters here"
	content := Content{
		[]string{
			"//TODO write your package here\n",
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
			getClassSkeleton(domain, classBody),
		},
	}
	return concatenateBody(content)
}

func GetDtoBody(name string) string {
	content := Content{
		[]string{
			"//TODO write your package here\n",
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

func getClassSkeleton(name string, body string) string {
	return "public class " + name + "{ \n\n" + body + "\n}"
}

func concatenateBody(content Content) string {
	var imports string
	for _, item := range content.Imports {
		imports = imports + item
	}

	var body string
	for _, item := range content.Content {
		body = body + item
	}

	return imports + "\n" + body
}
