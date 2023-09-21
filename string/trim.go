package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CleanupMessage("******* HI Hemant! ******"))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.TrimLeft(oldMsg, "*")
	oldMsg = strings.TrimRight(oldMsg, "*")
	return strings.TrimSpace(oldMsg)

}
