package lib

import "fmt"

func Greet(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s!", name)
	}
	return "Hello, anonymous!"
}
