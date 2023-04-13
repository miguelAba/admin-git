package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Folder struct {
	Name     string
	Type     string
	Path     string
	Children []Folder
	Content  string
}

func ApiGit(url string) *http.Response {
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

func GetFolder(patch string) []Folder {
	res := ApiGit("https://api.github.com/repos/miguelAba/my-protos/contents/" + patch + "?ref=main")
	repos := []Folder{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&repos)
	if err != nil {
		fmt.Println(err)
	}
	return repos
}

func GetFile(patch string) Folder {
	res := ApiGit("https://api.github.com/repos/miguelAba/my-protos/contents/" + patch + "?ref=main")
	file := Folder{}
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&file)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func GetFolderRepo(patch string) Folder {

	name := regexp.MustCompile(`\w+$`).FindString(patch)
	slave := Folder{Name: name, Path: patch, Type: "dir"}
	folders := GetFolder(patch)

	for _, folder := range folders {
		if folder.Type == "file" {
			file := GetFile(folder.Path)
			slave.Children = append(slave.Children, file)

		}
		if folder.Type == "dir" {
			slave.Children = append(slave.Children, GetFolderRepo(folder.Path))
		}
	}
	return slave
}
