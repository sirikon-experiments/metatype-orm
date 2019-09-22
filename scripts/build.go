package scripts

func Build()  {
	Pipe(".", "mkdir", "-p", "./out")
	Pipe(".", "go", "build", "-o", "./out/metatype", "./main.go")
}
