package greet

import "fmt"

func English(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Japanese(name string) string {
	return fmt.Sprintf("こんにちは、%s", name)
}
