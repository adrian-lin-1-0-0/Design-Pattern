package file

import (
	"os"
)

func Byte2File(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
