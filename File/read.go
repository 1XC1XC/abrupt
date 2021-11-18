package File 

import "os"
)

func Exists(arguments ...interface{}) interface{} {
	var Size int = len(arguments)
	Data, Error := os.Stat(arguments[0].(string))
	
	if Size == 1 && Error == nil {
		return true
	} else if Size == 2 {
		if arguments[1].(bool) != true {
			return false 
		}
		if Data.IsDir() {
			return "folder"
		}
	
		return "file" 
	}
	return false
}

func ReadDirectory(title string) interface{} {
	files, Error := os.ReadDir(title)
	if Error != nil {
		return false
	}
	var List []string 
	for _, file := range files {
		List = append(List, file.Name())
	}
	return List
}

func ReadFile(title string) interface{} {
	data, Error := os.ReadFile(title)
	if Error != nil {
		return false
	}
	return string(data) 
}

func Read(title string) interface{} {
	Data, Error := os.Stat(title)

	if Error != nil {
		return false
	}
	
	if Data.IsDir() {
		return ReadDirectory(title)
	}

	return ReadFile(title) 
}