package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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

func Confirm(err error, message string) {
	prefix := color.GreenString("Success:")
	if err != nil {
		prefix = color.RedString("Fail:")
	}
	fmt.Println(prefix, message)
}

func Warn(message string) {
	prefix := color.YellowString("Warning:")
	fmt.Println(prefix, message)
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

func Error(err error) {
	color.Red(err.Error())
}

func Ask(prompt string) bool {
	new_prompt := prompt + " [y/n]: "
	inp := Input(new_prompt)
	inp = strings.ToLower(inp)
	switch inp {
	case "y", "yes":
		return true
	default:
		return false
	}
}

func PressEnter() bool {
	fmt.Print("Press enter to continue...")
	var b [1]byte
	os.Stdin.Read(b[:])
	fmt.Println()
	return true
}
