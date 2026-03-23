package main

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"

	"tsBootstrup/src/cmd"
	"tsBootstrup/src/git"
	"tsBootstrup/src/npm"
	"tsBootstrup/src/ts"
	"tsBootstrup/src/utils"
)

var isInit bool
var isRun bool
var useGit bool
var noReadme bool

func init() {
	flag.BoolVar(&isInit, "init", false, "Don't ask confirm to create a project")
	flag.BoolVar(&isInit, "i", false, "Don't ask confirm to create a project")

	flag.BoolVar(&isRun, "run", false, "To run the project")
	flag.BoolVar(&isRun, "r", false, "To run the project")

	flag.BoolVar(&useGit, "git", false, "To init git repo and create init commit")
	flag.BoolVar(&useGit, "g", false, "To init git repo and create init commit")

	flag.BoolVar(&noReadme, "noReadme", false, "To decline creating Readme")

	flag.Parse()
}

func initProject() {
	start := time.Now()
	var wg sync.WaitGroup

	// 1. npm init (синхронно, обязательно первым)
	npm.Init(cmd.ShowOnlyErrors)

	// 2. параллельные задачи
	wg.Go(func() {
		npm.Install(cmd.ShowOnlyErrors, "typescript")
	})

	wg.Go(func() {
		ts.CreateIndexFile(indexFileContent)
	})

	wg.Wait()

	// 3. зависимые шаги
	ts.Init(cmd.ShowOnlyErrors)
	ts.UpdateConfig("tsconfig.json", tsconfigContent)
	npm.UpdatePackageJSON()

	// 4. git (если нужно)
	if useGit {
		git.Init()
	}

	// 5. README (ВСЕГДА после initProject)
	if !noReadme {
		utils.CreateReadMe(readMeContent)
	}

	// 6. финальный лог
	prefix := color.CyanString("Ended for:")
	elapsed := time.Since(start).Round(time.Millisecond)
	fmt.Println(prefix, elapsed)

	utils.ShowInfo()
}

func main() {
	if utils.FileExists(blockadeFileName) {
		err := errors.New("BLOCKADE FILE FOUND")
		cmd.Confirm(err, blockadeMessage)
		cmd.PressEnter()
		return
	}
	if err := utils.CheckASCIIPath(); err != nil {
		cmd.Confirm(err, "Invalid project path for npm")
		cmd.PressEnter()
		return
	}
	args := cmd.GetArgs()

	// Eсли пользователь просто запустил бинарь
	if len(args) == 0 {
		if cmd.Ask("Do you want to create a project?") {
			initProject()
		}
		return
	}

	// Явный init через флаг
	if isInit {
		initProject()
	}

	// run можно отдельно
	if isRun {
		npm.RunProject(cmd.ShowAll)
	}
}
