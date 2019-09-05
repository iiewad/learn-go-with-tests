package main

import "fmt"

const chinese = "Chinese"
const chineseHelloPrefix = "你好, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const englishPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", chinese))
}
