package npm

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"tsBootstrup/internal/cmd"
)

func Init(outputMode cmd.Output) error {
	err := cmd.Run(outputMode, "npm", "init", "-y")
	cmd.Confirm(err, "npm init -y")
	return err
}

func Install(outputMode cmd.Output, pkg string) error {
	err := cmd.Run(outputMode, "npm", "i", pkg)
	cmd.Confirm(err, "npm i "+pkg)
	return err
}

func DevInstall(outputMode cmd.Output, pkg string) error {
	err := cmd.Run(outputMode, "npm", "i", "-D", pkg)
	cmd.Confirm(err, "npm i "+pkg)
	return err
}

func UpdatePackageJSON(nodeType string) error {
	data, err := os.ReadFile("package.json")
	if err != nil {
		cmd.Confirm(err, "read package.json")
		return err
	}

	var pkg map[string]any
	json.Unmarshal(data, &pkg)
	pkg["type"] = nodeType
	scripts, ok := pkg["scripts"].(map[string]any)
	if !ok {
		scripts = make(map[string]any)
	}

	scripts["start"] = "npx tsc && node dist/index.js"
	pkg["scripts"] = scripts

	newData, _ := json.MarshalIndent(pkg, "", "  ")
	err = os.WriteFile("package.json", newData, 0644)

	cmd.Confirm(err, "update package.json")
	return err
}

func RunProject(outputMode cmd.Output) {
	err := cmd.Run(outputMode, "npm", "start")
	cmd.Confirm(err, "npm start")
}

var npmNameRegexp = regexp.MustCompile(`^[a-z0-9]+([\-_.][a-z0-9]+)*$`)

func IsValidName(name string) bool {
	if len(name) == 0 || len(name) > 214 {
		return false
	}

	if strings.HasPrefix(name, ".") {
		return false
	}

	if strings.HasPrefix(name, "-") || strings.HasPrefix(name, "_") {
		return false
	}

	return npmNameRegexp.MatchString(name)
}
