package bases

import "strings"

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
	return getClassSkeleton(name, "//TODO implement parameters here")
}

func generatedValue() Annotation {
	return Annotation{
		"import javax.persistence.GeneratedValue;\n",
		annotation("generatedValue"),
	}
}

func id() Annotation {
	return Annotation{
		"import javax.persistence.Id;\n",
		annotation("id"),
	}
}

func noArgsConstructor() Annotation {
	return Annotation{
		"import lombok.NoArgsConstructor;\n",
		annotation("noArgsConstructor"),
	}
}

func allArgsConstructor() Annotation {
	return Annotation{
		"import lombok.AllArgsConstructor;\n",
		annotation("allArgsConstructor"),
	}
}

func entity() Annotation {
	return Annotation{
		"import javax.persistence.Entity;\n",
		annotation("entity"),
	}
}

func builder() Annotation {
	return Annotation{
		"import lombok.Builder;\n",
		annotation("builder"),
	}
}

func data() Annotation {
	return Annotation{
		"import lombok.Data;\n",
		annotation("data"),
	}
}

func equalsAndHashCode() Annotation {
	return Annotation{
		"import lombok.EqualsAndHashCode;\n",
		annotation("EqualsAndHashCode"),
	}
}

func annotation(annotation string) string {
	return "@" + strings.Title(annotation) + "\n"
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
