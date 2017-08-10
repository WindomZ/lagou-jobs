package output

import "github.com/WindomZ/go-develop-kit/path"

func Access(filePath string) error {
	return path.Ensure(filePath, false)
}

func Write(filePath string, s ...string) error {
	return path.OverwriteFile(filePath, s...)
}
