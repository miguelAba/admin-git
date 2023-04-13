package main

import (
	ctr "admin-git/src/controller"
	"flag"
	"fmt"

	"github.com/manifoldco/promptui"
)

// struct data
// add cache

func getLang() string {
	items := []string{"Typescript", "Ruby"}

	prompt := promptui.Select{
		Label: "Selecciona el lenguaje de programación",
		Items: items,
		Size:  2,
		Templates: &promptui.SelectTemplates{
			Active:   "> {{ . | cyan }}",
			Inactive: "  {{ . | white }}",
			Selected: "{{ . | green }}",
		},
	}

	_, lang, err := prompt.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	if lang == "Typescript" {
		lang = ".ts"
	} else if lang == "Ruby" {
		lang = ".rb"
	}

	return lang
}

func getProjects(tree ctr.Folder) string {

	items := []string{}
	for _, sub := range tree.Children[0].Children {
		if sub.Type == "dir" {
			items = append(items, sub.Name)
		}
	}

	prompt := promptui.Select{
		Label: "Selecciona el lenguaje de programación",
		Items: items,
		Size:  len(items),
		Templates: &promptui.SelectTemplates{
			Active:   "> {{ . | cyan }}",
			Inactive: "  {{ . | white }}",
			Selected: "{{ . | green }}",
		},
	}

	_, lang, err := prompt.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}

	if lang == "Typescript" {
		lang = ".ts"
	} else if lang == "Ruby" {
		lang = ".rb"
	}

	return lang
}

func main() {
	route := flag.String("route", "", "route to save the files")
	flag.Parse()

	tree := ctr.GetFolderRepo("")
	ctr.CreateTree(tree, getLang(), getProjects(tree), *route)
}
