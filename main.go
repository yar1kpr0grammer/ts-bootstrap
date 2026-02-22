package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func printConfirm(err error, message string) {
	if err != nil {
		fmt.Println("❌", message)
	} else {
		fmt.Println("✅", message)
	}
}

func nodeProjectInit() {
	err := runCommand("npm", "init", "-y")
	printConfirm(err, "npm init -y")
}

func npmInstall(pkg string) {
	err := runCommand("npm", "i", pkg)
	printConfirm(err, "npm i "+pkg)
}

func tsInit() {
	os.Remove("tsconfig.json")
	err := runCommand("npx", "tsc", "--init")
	printConfirm(err, "npx tsc --init")
}

func uncomment(line string) string {
	re := regexp.MustCompile(`//\s*`)
	return re.ReplaceAllString(line, "")
}

func updateTSConfig() {
	data, err := os.ReadFile("tsconfig.json")
	if err != nil {
		printConfirm(err, "read tsconfig.json")
		return
	}

	lines := regexp.MustCompile("\n").Split(string(data), -1)
	newLines := []string{}

	for _, line := range lines {
		if regexp.MustCompile(`rootDir|outDir`).MatchString(line) {
			line = uncomment(line)
		}
		newLines = append(newLines, line)
	}

	err = os.WriteFile("tsconfig.json", []byte(
		regexp.MustCompile("\n").ReplaceAllString(
			fmt.Sprint(joinLines(newLines)), "\n")), 0644)

	printConfirm(err, "update tsconfig.json")
}

func joinLines(lines []string) string {
	result := ""
	for i, l := range lines {
		if i != 0 {
			result += "\n"
		}
		result += l
	}
	return result
}

func updatePackageJSON() {
	data, err := os.ReadFile("package.json")
	if err != nil {
		printConfirm(err, "read package.json")
		return
	}

	var pkg map[string]interface{}
	json.Unmarshal(data, &pkg)

	scripts, ok := pkg["scripts"].(map[string]interface{})
	if !ok {
		scripts = make(map[string]interface{})
	}

	scripts["start"] = "npx tsc && node dist/index.js"
	pkg["scripts"] = scripts

	newData, _ := json.MarshalIndent(pkg, "", "  ")
	err = os.WriteFile("package.json", newData, 0644)

	printConfirm(err, "update package.json")
}

func createIndexFile() {
	path := filepath.Join("src", "index.ts")
	os.MkdirAll("src", os.ModePerm)

	content := []byte("console.log('Hello, world!')")
	err := os.WriteFile(path, content, 0644)

	printConfirm(err, "create src/index.ts")
}

func runProject() {
	cmd := exec.Command("npm", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	printConfirm(err, "npm start")
}

func showInfo(){
	fmt.Println("--------------------")
	fmt.Println("ℹ️\tЧтобы запустить проект:")
	fmt.Println("npm start ")
}

func main() {
	nodeProjectInit()
	npmInstall("typescript")
	tsInit()
	updateTSConfig()
	createIndexFile()
	updatePackageJSON()
	runProject()
	showInfo()
}