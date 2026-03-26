package project

import (
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
	var wg sync.WaitGroup

	// 1. npm init (синхронно, обязательно первым)
	npm.Init(cmd.ShowOnlyErrors)

	// 2. параллельные задачи
	wg.Add(2)

	go func() {
		defer wg.Done()
		npm.Install(cmd.ShowOnlyErrors, "typescript")
	}()

	go func() {
		defer wg.Done()
		ts.CreateIndexFile(config.IndexFileContent)
	}()

	wg.Wait()

	// 3. зависимые шаги
	ts.Init(cmd.ShowOnlyErrors)
	ts.SetConfig("tsconfig.json", config.TsconfigContent)
	npm.UpdatePackageJSON()

	if s.UseGit {
		git.Init()
	}

	if s.CreateReadme {
		utils.CreateReadMe(config.ReadMeContent)
	}

	return nil
}
