package File

import "os"

func CreateDirectory(title string) bool {
	err := os.Mkdir(title, 0755)
	if err != nil {
		return false
	}
	return true 
}

func CreateFile(title string, content string) bool {
	file, err := os.Create(title)
	if err != nil {
		return false
	}
	file.WriteString(content)
	file.Close()
	return true
} 

func Create(arguments ...string) bool {	
	var Return bool

	switch len(arguments) {
		case 1:
			Return = CreateDirectory(arguments[0])
		case 2:
			Return = CreateFile(arguments[0], arguments[1])
		default:
			Return = false 
	}

	return Return 
}