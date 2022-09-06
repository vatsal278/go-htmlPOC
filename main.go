package main

import (
	"html/template"
	"log"
	"os"
)

// declaring struct
type Student struct {
	Name  string
	Marks int
	Id    string
}

type Class []Student

// main function
func main() {
	var class Class
	// defining struct instance
	std1 := Student{"A", 90, "1"}
	std2 := Student{"B", 100, "2"}
	std3 := Student{"C", 88, "3"}
	std4 := Student{"D", 25, "4"}
	std5 := Student{"E", 35, "5"}
	class = append(class, std4, std2, std3, std1, std5)
	//os.Setenv("list", "")
	//list := os.Getenv("list")
	t, err := template.ParseFiles("failure.html")
	if err != nil {
		log.Print(err)
	}
	f, err := os.Create("newfile")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = t.Execute(f, class)
	if err != nil {
		log.Print(err)
	}
	// standard output to print merged data
	err = t.Execute(os.Stdout, class)
	if err != nil {
		log.Print(err)
	}

}
