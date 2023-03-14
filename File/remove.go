package File 

import "os"

func RemoveFile(title string) bool {
	v := os.Remove(title)
	return v != nil
}

func RemoveFolder(title string) bool {
	v := os.RemoveAll(title)
	return v != nil
}

func Remove(title string) bool {
	Data, Error := os.Stat(title)

	if Error != nil { 
		return false
	}

	if Data.IsDir() {
		return RemoveFile(title)
	} else {
		return RemoveFolder(title)
	}
}