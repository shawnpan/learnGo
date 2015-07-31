package testbase

import (
	"bufio"
	"fmt"
	"os"
)

func TestStdin() {
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(line, hasMoreInLine, err)
}
