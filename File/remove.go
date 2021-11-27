package File 

import "os"

// Reasonably wouldn't require multiple functions, but for consistency, it will retain.

func RemoveFile(title string) bool {
	v := os.Remove(title)

	if v != nil {
		return false
	}

	return true
}

func RemoveFolder(title string) bool {
	v := os.RemoveAll(title)

	if v != nil {
		return false
	}

	return true
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
	
	return false
}