package main

import (
	"flag"
	"fmt"

	"tsBootstrup/internal/cmd"
	"tsBootstrup/internal/config"
	"tsBootstrup/internal/npm"
	"tsBootstrup/internal/project"
	"tsBootstrup/internal/utils"
)

type Flags struct {
	Init     bool
	Run      bool
	UseGit   bool
	NoReadme bool
	Clear    bool
}

func parseFlags() Flags {
	var f Flags

	flag.BoolVar(&f.Init, "init", false, "create project without prompt")
	flag.BoolVar(&f.Init, "i", false, "create project without prompt")

	flag.BoolVar(&f.Clear, "clear", false, "clear current directory")
	flag.BoolVar(&f.Clear, "c", false, "clear current directory")

	flag.BoolVar(&f.Run, "run", false, "run project")
	flag.BoolVar(&f.Run, "r", false, "run project")

	flag.BoolVar(&f.UseGit, "git", false, "init git repo")
	flag.BoolVar(&f.UseGit, "g", false, "init git repo")

	flag.BoolVar(&f.NoReadme, "noReadme", false, "skip README creation")

	flag.Parse()

	return f
}

func mustStop(err error, msg string) bool {
	if err == nil {
		return false
	}

	cmd.Confirm(err, msg)
	cmd.PressEnter()
	return true
}
func measuteProjectInit(settings project.Settings) {
	utils.MeasureTimeErr("Init", func() error {
		return project.Init(settings)
	})

}

func main() {
	f := parseFlags()

	if utils.FileExists(config.BlockadeFileName) {
		cmd.Confirm(
			fmt.Errorf("blockade file found"),
			config.BlockadeMessage,
		)
		cmd.PressEnter()
		return
	}

	// Not ASCII path can break npm init
	if err := utils.CheckPathForNPM(); err != nil {
		cmd.Warn(err.Error())
	}

	// run mode
	if f.Run {
		npm.RunProject(cmd.ShowAll)
		return
	}

	// clear mode
	if f.Clear {
		if cmd.Ask("Do you want to clear current diectory?") {
			if err := project.Remove(); err != nil {
				mustStop(err, "remove current directory")
			}
		}
		return
	}

	// Init mode
	settings := project.Settings{
		UseGit:       f.UseGit,
		CreateReadme: !f.NoReadme,
	}

	if f.Init {
		measuteProjectInit(settings)
		return
	}

	if cmd.Ask("Do you want to create a project?") {
		measuteProjectInit(settings)
	}
}
