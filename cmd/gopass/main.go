package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	genpass "github.com/ductnn/gopass/internal/generator"
	"github.com/fatih/color"
)

const banner = `

██████╗  ██████╗     ██████╗  █████╗ ███████╗███████╗
██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██╔════╝██╔════╝
██║  ███╗██║   ██║    ██████╔╝███████║███████╗███████╗
██║   ██║██║   ██║    ██╔═══╝ ██╔══██║╚════██║╚════██║
╚██████╔╝╚██████╔╝    ██║     ██║  ██║███████║███████║
 ╚═════╝  ╚═════╝     ╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝
`

var (
	length      = flag.Int("length", 16, "Specify the password length")
	numDigits   = flag.Int("digits", 4, "Number of digits to include in the password")
	numSymbols  = flag.Int("symbols", 4, "Number of digits to include in the password")
	noUpper     = flag.Bool("no-upper", true, "Excludes uppercase letters from the results")
	allowRepeat = flag.Bool("allow-repeat", true, "Allows characters to repeat")
	loops       = flag.Int("copies", 1, "Generate copies of password")
)

// End
func end_program() {
	var goodbye, githubURL, dockerURL string

	goodbye = "Star the project on GitHub if you liked this tool"
	githubURL = "https://github.com/ductnn/gopass"
	dockerURL = "https://hub.docker.com/r/ductn4/gopass"

	color.HiWhite("\n🐳 You can pull docker image in: " + color.HiCyanString(dockerURL))
	color.HiGreen("\n⭐️ " + goodbye)
	color.HiYellow("\n👉 " + githubURL + " 👈")
	color.HiWhite("\n🎉 Thank you so much 🎉")

	fmt.Printf("\n")
}

func main() {
	var passwords []string
	fmt.Println(banner)

	flag.Parse()

	// generate passwords by loop times
	for i := 0; i < *loops; i++ {
		res, err := genpass.Generate(*length, *numDigits, *numSymbols, *noUpper, *allowRepeat)
		if err != nil {
			log.Fatal(err)
		}

		passwords = append(passwords, res)
	}

	result := strings.Join(passwords, "\n")

	// print generated passwords at last
	color.Green("====================== Starting ======================")
	fmt.Println()
	color.Green("        Your Password: " + color.HiYellowString(result))
	fmt.Println()
	color.Green("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	end_program()
}
