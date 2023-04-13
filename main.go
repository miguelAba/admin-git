package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
)

// struct data
// add cache
// select console

type Folder struct {
	Name     string
	Type     string
	Path     string
	Children []Folder
	Content  string
}

func apiGit(url string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header = http.Header{
		"Accept":               {"application/vnd.github+json"},
		"X-GitHub-Api-Version": {"2022-11-28"},
		"Authorization":        {"Bearer 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"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func getFolder(patch string) []Folder {
	res := apiGit("https://api.github.com/repos/miguelAba/my-protos/contents/" + patch + "?ref=main")
	repos := []Folder{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&repos)
	if err != nil {
		fmt.Println(err)
	}
	return repos
}

func getFile(patch string) Folder {
	res := apiGit("https://api.github.com/repos/miguelAba/my-protos/contents/" + patch + "?ref=main")
	file := Folder{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&file)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func getFolderRepo(patch string) Folder {

	name := regexp.MustCompile(`\w+$`).FindString(patch)
	slave := Folder{Name: name, Path: patch, Type: "dir"}
	folders := getFolder(patch)

	for _, folder := range folders {
		if folder.Type == "file" {
			file := getFile(folder.Path)
			slave.Children = append(slave.Children, file)

		}
		if folder.Type == "dir" {
			slave.Children = append(slave.Children, getFolderRepo(folder.Path))
		}
	}
	return slave
}

func createTree(folder Folder, language string, project string) {

	for _, sub := range folder.Children {
		if sub.Type == "dir" && (sub.Name == project || sub.Name == "protos") {
			os.MkdirAll(sub.Path, os.ModePerm)
			createTree(sub, language, project)
		}

		if sub.Type == "file" {

			matchLang, _ := regexp.MatchString(language, sub.Name)
			matchProto, _ := regexp.MatchString(".proto", sub.Name)

			if matchLang || matchProto {

				dec, err := base64.StdEncoding.DecodeString(sub.Content)
				if err != nil {
					panic(err)
				}

				f, err := os.Create(sub.Path)
				if err != nil {
					panic(err)
				}
				defer f.Close()

				if _, err := f.Write(dec); err != nil {
					panic(err)
				}
				if err := f.Sync(); err != nil {
					panic(err)
				}
			}
		}

	}

}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

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

func getProjects(tree Folder) string {

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
	tree := getFolderRepo("")
	createTree(tree, getLang(), getProjects(tree))
}
