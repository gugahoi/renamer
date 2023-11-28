package main

import (
	"fmt"
	"os"
	"strings"
)

// Renamer renames environment variables from OLD_PREFIX to NEW_PREFIX.
// Usage: renamer OLD_PREFIX NEW_PREFIX
func main() {
	// read the 1st and 2nd command line argumens into variables
	oldPrefix := os.Args[1]
	newPrefix := os.Args[2]

	envVars := os.Environ()

	for _, v := range envVars {
		if strings.HasPrefix(v, oldPrefix) {
			parts := strings.SplitN(v, "=", 2)

			oldName := parts[0]
			value := parts[1]

			newName := strings.Replace(oldName, oldPrefix, newPrefix, 1)

			os.Setenv(newName, value)
			fmt.Printf("%s becomes %s=%s\n", oldName, newName, value)
		}
	}
}
