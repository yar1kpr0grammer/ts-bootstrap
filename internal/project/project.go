package project

import (
	"fmt"
	"os"
	"sync"

	"tsBootstrup/internal/cmd"
	"tsBootstrup/internal/config"
	"tsBootstrup/internal/git"
	"tsBootstrup/internal/languages"
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
	name := cmd.Input(languages.Get("ask_project_name"))

	if name == "" {
		name = "."
	}

	if name == "." {
		return ".", os.Chdir(".")
	}

	info, err := os.Stat(name)

	if os.IsNotExist(err) {
		if err := os.Mkdir(name, 0755); err != nil {
			return "", fmt.Errorf("%s: %w", languages.Get("error_create_dir"), err)
		}

		if err := os.Chdir(name); err != nil {
			return "", fmt.Errorf("%s: %w", languages.Get("error_open_dir"), err)
		}

		cmd.Confirm(nil, languages.Get("created_project_dir"))
		return name, nil
	}

	if err != nil {
		return "", fmt.Errorf("%s: %w", languages.Get("error_check_dir"), err)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("%s", languages.Get("error_not_directory"))
	}

	if err := os.Chdir(name); err != nil {
		return "", fmt.Errorf("%s: %w", languages.Get("error_open_dir"), err)
	}

	cmd.Warn(languages.Get("directory_exists"))

	empty, err := isEmpty()
	if err != nil {
		return "", err
	}

	if !empty {
		if cmd.Ask(languages.Get("ask_clear_dir")) {
			if err := Remove(); err != nil {
				return "", err
			}

			cmd.Confirm(nil, languages.Get("directory_cleaned"))
		} else {
			return "", fmt.Errorf("%s", languages.Get("error_dir_not_empty"))
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

	if !cmd.Ask(languages.Get("ask_continue_not_empty")) {
		return fmt.Errorf("%s", languages.Get("error_cancelled"))
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
	fmt.Println(languages.Get("now"))

	if name != "." && name != "" {
		fmt.Printf("cd %s\n", name)
	}

	fmt.Println("npm start")
	fmt.Println("----------------")
}

func isEmpty() (bool, error) {
	exeName, err := utils.ThisFile()
	if err != nil {
		return false, fmt.Errorf("%s: %w", languages.Get("error_get_exe"), err)
	}

	files, err := os.ReadDir(".")
	if err != nil {
		return false, fmt.Errorf("%s: %w", languages.Get("error_read_dir"), err)
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
		return fmt.Errorf("%s: %w", languages.Get("error_read_dir"), err)
	}

	exeName, err := utils.ThisFile()
	if err != nil {
		return fmt.Errorf("%s: %w", languages.Get("error_get_exe"), err)
	}

	for _, file := range files {
		if file.Name() == exeName {
			continue
		}

		if err := os.RemoveAll(file.Name()); err != nil {
			return fmt.Errorf("%s %s: %w",
				languages.Get("error_remove_file"),
				file.Name(),
				err,
			)
		}
	}

	cmd.Confirm(nil, languages.Get("removed_all_files"))
	return nil
}
