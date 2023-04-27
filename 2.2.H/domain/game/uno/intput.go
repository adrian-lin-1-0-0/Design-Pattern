package uno

import (
	"bufio"
	"os"
	"strconv"
)

func inputString() string {
	reader := bufio.NewReader(os.Stdin)
	read_line, _ := reader.ReadString('\n')
	return read_line[:len(read_line)-1]
}

func inputInt() int {
	reader := bufio.NewReader(os.Stdin)
	read_line, _ := reader.ReadString('\n')
	read_line = read_line[:len(read_line)-1]
	num, _ := strconv.Atoi(read_line)
	return num
}
