package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s ", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author  /*author struct*/
}

/*Create slice of struct(blogPost)*/
type website struct {
	blogPosts []blogPost
}

func (b blogPost) details() {
	fmt.Println("Title ", b.title)
	fmt.Println("Content ", b.content)
	fmt.Println("Author ", b.author.fullName()) /*fullName from author*/
	fmt.Println("Bio ", b.author.bio)
}

func (w website) contents() {
	fmt.Println("Content of website")
	for _, value := range w.blogPosts {
		fmt.Println()
		value.details()

	}
}

/* Main function*/
func main() {
	author1 := author{
		"Chayan",
		"Roy",
		"CSE student",
	}

	b := blogPost{
		"Golang Learning ",
		"Learning composition instead of inheritance",
		author1,
	}

	b.details()

	/*Create slice of struct(blogPost)*/
	w := website{
		blogPosts: []blogPost{
			{"Concurrency",
				"Golang is a concurrent language",
				author1},
			{"Faster",
				"Golang is a faster language",
				author1,
			},
		},
	}
	w.contents()

}
