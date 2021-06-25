package bases

import "strings"

func service() Annotation {
	return Annotation{
		"import org.springframework.stereotype.Service;\n",
		annotation("service"),
	}
}

func repository() Annotation {
	return Annotation{
		"import org.springframework.stereotype.Repository;\n",
		annotation("repository"),
	}
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
