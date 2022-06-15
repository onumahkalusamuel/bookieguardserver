package helpers

import "io/ioutil"

func GetFileContent(filepath string) string {

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return ""
	}

	return string(file)
}
