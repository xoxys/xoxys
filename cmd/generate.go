package main

import (
	"log"
	"os"
	"text/template"

	"github.com/mmcdole/gofeed"
)

// Declare type pointer to a template
var tmpl *template.Template

type TmplVars struct {
	FeedItems []*gofeed.Item
}

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	tmpl = template.Must(template.ParseFiles("src/readme.tmpl"))
}

func readFeed() []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://thegeeklab.de/atom.xml")

	return feed.Items[0:5]
}

func main() {
	output, err := os.Create("README.md")
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	vars := TmplVars{
		FeedItems: readFeed(),
	}

	err = tmpl.Execute(output, vars)
	if err != nil {
		log.Fatalln(err)
	}
}
