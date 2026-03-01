package main

import (
	"flag"
	"tsBootstrup/src/cmd"
	"tsBootstrup/src/git"
	"tsBootstrup/src/npm"
	"tsBootstrup/src/ts"
	"tsBootstrup/src/utils"
)

func initProject() {
	npm.Init(cmd.ShowOnlyErrors)
	npm.Install(cmd.ShowOnlyErrors, "typescript")
	ts.Init(cmd.ShowOnlyErrors)
	ts.UpdateConfig("tsconfig.json", tsconfigContent)
	ts.CreateIndexFile(indexFileContent)
	npm.UpdatePackageJSON()
	utils.ShowInfo()
}

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

func main() {
	args := cmd.GetArgs()

	if len(args) == 0 {
		createProject := cmd.Ask("Do you want to create a project?")
		if createProject {
			initProject()
		}
		return
	}

	if isInit {
		initProject()
	}

	if useGit {
		git.Init()
	}
	if !noReadme {
		utils.CreateReadMe(readMeContent)
	}
	if isRun {
		npm.RunProject(cmd.ShowAll)
	}
}
