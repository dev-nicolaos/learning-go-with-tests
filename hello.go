package main

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"

	hello = "Hello"
)

var languages = map[string]map[string]string{
	english: map[string]string{
		hello: hello,
	},
	spanish: map[string]string{
		hello: "Hola",
	},
	french: map[string]string{
		hello: "Bonjour",
	},
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = english
	}

	return languages[language][hello] + " " + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
