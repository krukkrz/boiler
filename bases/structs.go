package bases

type Annotation struct {
	Import  string
	Content string
}

type Content struct {
	Imports []string
	Content []string
}

type Class struct {
	Import  string
	Content string
}
