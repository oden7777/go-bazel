package main

import "fmt"

func main() {
	fmt.Println(Hello("Service"))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
