package hangmanweb

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetRandomWord(filename string) string {
	lines := GetFile(filename, "\n")
	cc := 0
	test := false
	for {
		pick := lines[RandInt(len(lines))]
		if pick == "" {
			cc++
			continue
		}
		for _, char := range pick {
			if char < 65 || char > 90 && char < 97 || char > 122 && char < 224 {
				test = true
				cc++
				if cc > 4 {
					fmt.Println("La liste contient trop d'invalide selection du mot impossible ^^'")
					os.Exit(0)
				}
				fmt.Print("Le mot choisit est invalide selection d'un nouveau mot - Tentative numéro " + string(rune(cc)+48))
			}
		}
		if !test {
			return pick
		}
	}
}

func GetFile(filename, sep string) []string {
	var testfile []string
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if string(content) == "" {
		fmt.Printf("Le fichier fournit %v est vide, action impossible\n", filename)
		os.Exit(0)
	} else if sep == "\n" {
		testfile = strings.Split(string(content), "\r\n")
		if len(testfile) <= 1 {
			testfile = strings.Split(string(content), "\n")
		}
	} else {
		testfile = strings.Split(string(content), "\r\n\r\n")
		if len(testfile) <= 1 {
			testfile = strings.Split(string(content), "\n\n")
		}
	}
	return testfile
}
