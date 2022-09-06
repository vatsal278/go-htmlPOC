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

// main function
func main() {
	var class []Student
	// defining struct instance
	std1 := Student{"A", 90, "1"}
	std2 := Student{"B", 100, "2"}
	std3 := Student{"C", 88, "3"}
	std4 := Student{"D", 65, "4"}
	std5 := Student{"E", 35, "5"}
	class = append(class, std4, std2, std3, std1, std5)
	os.Setenv("list", "merit")
	list := os.Getenv("list")
	if list == "fail" {
		for _, student := range class {
			if student.Marks < 40 {
				t, err := template.ParseFiles("failure.html")
				if err != nil {
					log.Print(err)
				}
				// standard output to print merged data
				err = t.Execute(os.Stdout, student)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
	if list == "merit" {
		for _, student := range class {
			if student.Marks > 90 {
				log.Print(student.Name)
				t, err := template.ParseFiles("search.html")
				if err != nil {
					log.Print(err)
				}
				// standard output to print merged data
				err = t.Execute(os.Stdout, student)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}
