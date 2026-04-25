package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"

	"tsBootstrup/internal/languages"
)

type Output struct {
	ShowInfo   bool
	ShowErrors bool
}

var ShowOnlyErrors = Output{
	ShowInfo:   false,
	ShowErrors: true,
}

var ShowAll = Output{
	ShowInfo:   true,
	ShowErrors: true,
}

var reader = bufio.NewReader(os.Stdin)

func GetArgs() []string {
	return os.Args[1:]
}

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
	prefix := color.GreenString(languages.Get("success_prefix"))

	if err != nil {
		prefix = color.RedString(languages.Get("fail_prefix"))
	}

	fmt.Println(prefix, message)
}

func Warn(message string) {
	prefix := color.YellowString(languages.Get("warning_prefix"))
	fmt.Println(prefix, message)
}

func Input(prompt string) string {
	fmt.Print(prompt)

	inp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(languages.Get("input_error"), err)
		return ""
	}

	return strings.TrimSpace(inp)
}

func Error(err error) {
	color.Red(err.Error())
}

func Ask(prompt string) bool {
	fullPrompt := prompt + " " + languages.Get("yes_no_suffix")

	inp := Input(fullPrompt)
	inp = strings.ToLower(strings.TrimSpace(inp))

	switch inp {
	case "y", "yes", "д", "да":
		return true
	default:
		return false
	}
}

func PressEnter() bool {
	fmt.Print(languages.Get("press_enter"))

	var b [1]byte
	os.Stdin.Read(b[:])

	fmt.Println()
	return true
}
