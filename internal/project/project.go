package project

import (
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
	isempty, err := isEmpty()
	if err != nil {
		return err
	}
	if !isempty {
		if cmd.Ask("Directory is not empty. Continue?") {
			if err := Remove(); err != nil {
				return err
			}
			cmd.Confirm(nil, "clear directory")
		}
	}

	var wg sync.WaitGroup

	// 1. npm init (синхронно, обязательно первым)
	npm.Init(cmd.ShowOnlyErrors)

	// 2. параллельные задачи
	wg.Add(2)

	go func() {
		defer wg.Done()
		npm.Install(cmd.ShowOnlyErrors, "typescript")
		npm.DevInstall(cmd.ShowOnlyErrors, "@types/node")
	}()

	go func() {
		defer wg.Done()
		ts.CreateIndexFile(config.IndexFileContent)
	}()

	wg.Wait()

	// 3. зависимые шаги
	ts.Init(cmd.ShowOnlyErrors)
	ts.SetConfig("tsconfig.json", config.TsconfigContent)
	npm.UpdatePackageJSON(config.NodeType)

	if s.UseGit {
		git.Init()
	}

	if s.CreateReadme {
		utils.CreateReadMe(config.ReadMeContent)
	}

	return nil
}

func isEmpty() (bool, error) {
	exeName, err := utils.ThisFile()
	if err != nil {
		return false, err
	}

	context, err := os.ReadDir(".")
	if err != nil {
		cmd.Confirm(err, "read current directory")
		return false, err
	}

	for _, file := range context {
		name := file.Name()

		// если нашли НЕ служебный файл -> папка не пустая
		if name != exeName && name != config.BlockadeFileName {
			return false, nil
		}
	}

	return true, nil
}

func Remove() error {
	content, err := os.ReadDir(".")
	if err != nil {
		cmd.Confirm(err, "read current directory")
		return err
	}

	exeName, err := utils.ThisFile()
	if err != nil {
		return err
	}

	for _, file := range content {
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
