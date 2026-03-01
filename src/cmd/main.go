package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Output struct {
	ShowInfo   bool
	ShowErrors bool
}

var ShowOnlyErrors = Output{ShowInfo: false, ShowErrors: true}
var ShowAll = Output{ShowInfo: true, ShowErrors: true}

func GetArgs() []string {
	return os.Args[1:]
}

var reader = bufio.NewReader(os.Stdin)

func Run(outputSettings Output, name string, args ...string) error {
	cmd := exec.Command(name, args...)

	if outputSettings.ShowInfo {
		cmd.Stdout = os.Stdout
	}

	if outputSettings.ShowErrors {
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func redText(text string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", text)
}

func greenText(text string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", text)
}

func Confirm(err error, message string) {
	if err != nil {
		fmt.Println(redText("Fail:"), message, err)
	} else {
		fmt.Println(greenText("Success:"), message)
	}
}

func Input(prompt string) string {
	fmt.Print(prompt)
	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return ""
	}
	inp = strings.TrimSpace(inp)
	return inp

}

func Ask(prompt string) bool {
	new_prompt := prompt + " [y/n]: "
	inp := Input(new_prompt)
	inp = strings.ToLower(inp)
	switch inp {
	case "y":
		return true
	default:
		return false
	}
}
