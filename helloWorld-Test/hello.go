package main

import "fmt"

const (
 spanish = "Spanish"
 french = "French"
 urdu = "Urdu"

 urduPrefixHello = "Salam, "
 englishPrefixHello = "Hello, "
 spanishPrefixHola = "Hola, "
 frenchPrefixHello = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return gettingPrefix(language) + name
}
func gettingPrefix(language string) (prefix string)  {

	switch language {
		case spanish:
			prefix = spanishPrefixHola
		case french:
			prefix = frenchPrefixHello
		case urdu:
			prefix = urduPrefixHello 
		default:
			prefix = englishPrefixHello
		}
		return 
	}


func main() {
	fmt.Println(Hello("", "Spanish"))
}
