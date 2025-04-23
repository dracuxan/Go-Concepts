package main

const (
	helloString = "Hello, "
)

func helloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return helloString + name
}
