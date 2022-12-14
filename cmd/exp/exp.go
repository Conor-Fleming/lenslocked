package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Con Sledge",
		Bio:  "<script>alert(\"Haha, you have been h4x0r3d!\");</script>",
		Age:  27,
	}

	err = t.Execute(os.Stdout, user)
}
