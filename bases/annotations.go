package bases

import "strings"

const webBindAnnotation = "import org.springframework.web.bind.annotation."

func requestBody() Annotation {
	return Annotation{
		springWebAnnotationImport("requestBody"),
		annotation("requestBody"),
	}
}

func postMapping(path string) Annotation {
	return springAnnotation("postMapping", path)
}

func putMapping(path string) Annotation {
	return springAnnotation("putMapping", path)
}

func getMapping(path string) Annotation {
	return springAnnotation("getMapping", path)
}

func deleteMapping(path string) Annotation {
	return springAnnotation("deleteMapping", path)
}

func springAnnotation(name string, path string) Annotation {
	return Annotation{
		springWebAnnotationImport(name),
		annotationWithParams(name, path),
	}
}

func component() Annotation {
	return Annotation{
		"import org.springframework.stereotype.Component;\n",
		annotation("component" + newLine()),
	}
}

func requestMapping() Annotation {
	return Annotation{
		"import org.springframework.web.bind.annotation.RequestMapping;\n",
		annotationWithParams("requestMapping", "\"/"+strings.ToLower(Domain)+"s\""),
	}
}

func restController() Annotation {
	return Annotation{
		"import org.springframework.web.bind.annotation.RestController;\n",
		annotation("restController" + newLine()),
	}
}

func service() Annotation {
	return Annotation{
		"import org.springframework.stereotype.Service;\n",
		annotation("service" + newLine()),
	}
}

func repository() Annotation {
	return Annotation{
		"import org.springframework.stereotype.Repository;\n",
		annotation("repository" + newLine()),
	}
}

func generatedValue() Annotation {
	return Annotation{
		"import javax.persistence.GeneratedValue;\n",
		annotation("generatedValue" + newLine()),
	}
}

func id() Annotation {
	return Annotation{
		"import javax.persistence.Id;\n",
		annotation("id" + newLine()),
	}
}

func noArgsConstructor() Annotation {
	return Annotation{
		"import lombok.NoArgsConstructor;\n",
		annotation("noArgsConstructor" + newLine()),
	}
}

func allArgsConstructor() Annotation {
	return Annotation{
		"import lombok.AllArgsConstructor;\n",
		annotation("allArgsConstructor" + newLine()),
	}
}

func entity() Annotation {
	return Annotation{
		"import javax.persistence.Entity;\n",
		annotation("entity" + newLine()),
	}
}

func builder() Annotation {
	return Annotation{
		"import lombok.Builder;\n",
		annotation("builder" + newLine()),
	}
}

func data() Annotation {
	return Annotation{
		"import lombok.Data;\n",
		annotation("data" + newLine()),
	}
}

func equalsAndHashCode() Annotation {
	return Annotation{
		"import lombok.EqualsAndHashCode;\n",
		annotation("EqualsAndHashCode" + newLine()),
	}
}

func annotation(annotation string) string {
	return "@" + strings.Title(annotation)
}

func annotationWithParams(annotation string, params string) string {
	return "@" + strings.Title(annotation) + "(" + params + ")"
}

func springWebAnnotationImport(annotation string) string {
	return webBindAnnotation + strings.Title(annotation) + ";" + newLine()
}
