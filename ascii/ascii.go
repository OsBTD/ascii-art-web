package ascii

import (
	"log"
	"os"
	"strings"
)

// return the map fully populated with all printable characters as keys and the corresponding ascii art characters as values
func Checknewline(inputsplit []string) bool {
	for _, line := range inputsplit {
		if len(line) != 0 {
			return false
		}
	}
	return true
}

func PrintArt(input, banner string) string {
	// here we read the text file containing the ascii art characters
	content, err := os.ReadFile("Banner/" + banner + ".txt")
	if err != nil {
		log.Fatal("Error : couldn't read file ", err)
	}
	noreturn := strings.ReplaceAll(string(content), "\r", "")
	Lines := strings.Split(noreturn, "\n")

	Replace := make(map[rune]([]string))

	Char := 32
	for i := 0; i < len(Lines); i += 9 {
		if i+9 <= len(Lines)-1 {
			Replace[rune(Char)] = Lines[i+1 : i+9]
		}
		if Char <= 126 {
			Char++
		}

	}

	var result string
	inputsplit := strings.Split(input, "\r\n")
	// for i := 0; i < len(inputsplit)-1; i++ {
	// 	if inputsplit[i] < " " || inputsplit[i] > "~" {
	// 		inputsplit = append(inputsplit[:i], inputsplit[:i+1]...)
	// 	}
	// }
	for idx, line := range inputsplit {
		if Checknewline(inputsplit) && idx != len(inputsplit)-1 {
			result += "\n"
			continue
		} else if len(line) == 0 && !Checknewline(inputsplit) {
			result += "\n"
		} else if len(line) != 0 && !Checknewline(inputsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(line); j++ {
					inputrune := rune(line[j])
					result += Replace[inputrune][i]

				}
				result += "\n"
			}
		}
	}
	return result
}
