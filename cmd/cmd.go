package main

import (
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func defaultURL() string {
	if len(os.Args) > 1 {
		if a := os.Args[1]; a == "--h" {
			return "www.example.com"
		}
		return os.Args[1]
	}
	return "www.example.com"
}

func sprintIfExists(n string, v string) {
	if v != "" {
		fmt.Printf("----> %s: %s\n", n, v)
	}
}

func aprintIfExists(n string, v []string, a bool) {
	if len(v) != 0 {
		fmt.Printf("----> %s: ", n)
		for i := range v {
			if !a && i == 10 {
				fmt.Printf("...")
				break
			}
			fmt.Printf("%s ", v[i])
		}
		fmt.Printf("\n")
	}
}

func prettyspector(gs *gscrape.gscrape, a bool) {
	sprintIfExists("Title", gs.Title())
	sprintIfExists("Author", gs.Author())
	sprintIfExists("Description", gs.Description())
	sprintIfExists("Generator", gs.Generator())
	sprintIfExists("Charset", gs.Charset())
	sprintIfExists("Language", gs.Language())
	sprintIfExists("Feed URL", gs.Feed())
	aprintIfExists("Keywords", gs.Keywords(), a)
	aprintIfExists("Links", gs.Links(), a)
	aprintIfExists("Images", gs.Images(), a)
}

func main() {
	var url = flag.String("u", defaultURL(), "URL to gscrape.")
	var all = flag.Bool("all", false, "Show full results.")
	flag.Parse()
	mi, err := gscrape.New(*url)
	if err != nil {
		exit("Something went wrong. Please, try again.")
	}
	prettyspector(mi, *all)
}
