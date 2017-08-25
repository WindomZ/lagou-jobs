package output

import "github.com/WindomZ/go-develop-kit/path"

// Access returns a error if the file cannot be accessed
func Access(filePath string) error {
	return path.Ensure(filePath, false)
}

// Write returns a error if the file cannot be written
func Write(filePath string, s ...string) error {
	return path.OverwriteFile(filePath, s...)
}
