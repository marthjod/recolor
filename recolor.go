package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var (
		replacements = []struct {
			regex *regexp.Regexp
			repl  string
		}{
			{
				regex: regexp.MustCompile(`\[(3[1-8])m`),
				repl:  "\033[${1}m",
			},
			{
				regex: regexp.MustCompile(`\[0m`),
				repl:  "\033[0m",
			},
			{
				regex: regexp.MustCompile(`\[39m`),
				repl:  "",
			},
		}
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		for _, replacement := range replacements {
			if replacement.regex.MatchString(line) {
				line = replacement.regex.ReplaceAllString(line, replacement.repl)
			}
		}
		fmt.Println(line)
	}
}
