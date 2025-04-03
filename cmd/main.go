package main

import (
	"fmt"

	"tutorial.ent.go/app/internal/db/ent"
)

func main() {
	fmt.Printf("Hello, World! %+v", ent.User{})
}
