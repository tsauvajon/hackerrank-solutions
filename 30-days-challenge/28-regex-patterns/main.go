package main

// TODO : use By sorting https://golang.org/pkg/sort/#example__sortKeys

import (
	"fmt"
	"regexp"
	"sort"
)

type person struct {
	name  string
	email string
}

type byFirstName []person

func main() {
	var (
		n           int
		persons     byFirstName
		name, email string
	)

	fmt.Scan(&n)

	persons = make(byFirstName, 0)

	r, _ := regexp.Compile("([a-z.]+)@gmail.com")

	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&email)

		p := person{
			name:  name,
			email: email,
		}

		if !r.MatchString(email) {
			continue
		}

		persons = append(persons, p)
	}

	sort.Sort(persons)

	for _, p := range persons {
		fmt.Println(p.name)
	}
}

func (p byFirstName) Len() int {
	return len(p)
}

func (p byFirstName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byFirstName) Less(i, j int) bool {
	return p[i].name < p[j].name
}
