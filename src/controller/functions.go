package controller

import (
	"encoding/base64"
	"os"
	"regexp"
)

func CreateTree(folder Folder, language string, project string, path string) {

	if path != "" {
		path = path + "/"
	}

	for _, sub := range folder.Children {
		if sub.Type == "dir" && (sub.Name == project || sub.Name == "protos") {
			os.MkdirAll(path+sub.Path, os.ModePerm)
			CreateTree(sub, language, project, path)
		}

		if sub.Type == "file" {

			matchLang, _ := regexp.MatchString(language, sub.Name)
			matchProto, _ := regexp.MatchString(".proto", sub.Name)

			if matchLang || matchProto {

				dec, err := base64.StdEncoding.DecodeString(sub.Content)
				if err != nil {
					panic(err)
				}

				f, err := os.Create(path + sub.Path)
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
