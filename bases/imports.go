package bases

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

func crudRepository() Class {
	return Class{
		"import org.springframework.data.repository.CrudRepository;\n",
		"CrudRepository",
	}
}

func list() Class {
	return Class{
		"import java.util.List;\n",
		"List",
	}
}

func entityNotFoundException() Class {
	return Class{
		"import javax.persistence.EntityNotFoundException;\n",
		"EntityNotFoundException",
	}
}
