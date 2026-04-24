package project

import (
	"fmt"
	"os"
	"sync"

	"tsBootstrup/internal/cmd"
	"tsBootstrup/internal/config"
	"tsBootstrup/internal/git"
	"tsBootstrup/internal/npm"
	"tsBootstrup/internal/ts"
	"tsBootstrup/internal/utils"
)

type Settings struct {
	UseGit       bool
	CreateReadme bool
}

func Init(s Settings) error {
	name, err := prepareProjectDir()
	if err != nil {
		return err
	}

	if err := ensureDirectoryReady(); err != nil {
		return err
	}

	if err := initBaseProject(); err != nil {
		return err
	}

	if err := createFiles(); err != nil {
		return err
	}

	if err := applyOptionalFeatures(s); err != nil {
		return err
	}

	printPostInfo(name)
	return nil
}

func prepareProjectDir() (string, error) {
	name := cmd.Input(`How to name a project? (Type "." for current dir): `)

	if name == "" {
		name = "."
	}

	// current directory
	if name == "." {
		return ".", os.Chdir(".")
	}

	info, err := os.Stat(name)

	// папка не существует -> создаём
	if os.IsNotExist(err) {
		if err := os.Mkdir(name, 0755); err != nil {
			return "", err
		}

		if err := os.Chdir(name); err != nil {
			return "", err
		}

		cmd.Confirm(nil, "created project directory")
		return name, nil
	}

	// другая ошибка доступа и т.д.
	if err != nil {
		return "", err
	}

	// существует, но это файл
	if !info.IsDir() {
		return "", fmt.Errorf("%s exists and is not a directory", name)
	}

	// существует папка -> заходим
	if err := os.Chdir(name); err != nil {
		return "", err
	}

	cmd.Warn("directory already exists")

	empty, err := isEmpty()
	if err != nil {
		return "", err
	}

	if !empty {
		if cmd.Ask("Directory contains files. Clear it?") {
			if err := Remove(); err != nil {
				return "", err
			}
			cmd.Confirm(nil, "directory cleaned")
		} else {
			return "", fmt.Errorf("directory is not empty")
		}
	}

	return name, nil
}

func ensureDirectoryReady() error {
	empty, err := isEmpty()
	if err != nil {
		return err
	}

	if empty {
		return nil
	}

	if !cmd.Ask("Directory is not empty. Continue?") {
		return fmt.Errorf("operation cancelled")
	}

	return Remove()
}

func initBaseProject() error {
	npm.Init(cmd.ShowOnlyErrors)
	return nil
}

func createFiles() error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		installDependencies()
	}()

	go func() {
		defer wg.Done()
		createSourceFiles()
	}()

	wg.Wait()

	return finalizeTypescript()
}

func installDependencies() {
	npm.Install(cmd.ShowOnlyErrors, "typescript")
	npm.DevInstall(cmd.ShowOnlyErrors, "@types/node")
}

func createSourceFiles() {
	ts.CreateIndexFile(config.IndexFileContent)
}

func finalizeTypescript() error {
	ts.Init(cmd.ShowOnlyErrors)
	ts.SetConfig("tsconfig.json", config.TsconfigContent)
	npm.UpdatePackageJSON(config.NodeType)
	return nil
}

func applyOptionalFeatures(s Settings) error {
	if s.UseGit {
		git.Init()
	}

	if s.CreateReadme {
		utils.CreateReadMe(config.ReadMeContent)
	}

	return nil
}

func printPostInfo(name string) {
	fmt.Println("----------------")
	fmt.Println("Now:")

	if name != "." && name != "" {
		fmt.Printf("cd %s\n", name)
	}

	fmt.Println("npm start")
	fmt.Println("----------------")
}

func isEmpty() (bool, error) {
	exeName, err := utils.ThisFile()
	if err != nil {
		return false, err
	}

	files, err := os.ReadDir(".")
	if err != nil {
		return false, err
	}

	for _, file := range files {
		name := file.Name()

		if name != exeName && name != config.BlockadeFileName {
			return false, nil
		}
	}

	return true, nil
}

func Remove() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	exeName, err := utils.ThisFile()
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == exeName {
			continue
		}

		if err := os.RemoveAll(file.Name()); err != nil {
			return err
		}
	}

	cmd.Confirm(nil, "removed all files (except self)")
	return nil
}
