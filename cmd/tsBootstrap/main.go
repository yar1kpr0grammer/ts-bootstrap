package main

import (
	"flag"
	"fmt"
	"log"

	"tsBootstrup/internal/cmd"
	"tsBootstrup/internal/config"
	"tsBootstrup/internal/languages"
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
	Lang     string
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

	flag.StringVar(&f.Lang, "lang", "", "language: ru / en")
	flag.StringVar(&f.Lang, "l", "", "language: ru / en")

	flag.Parse()

	return f
}

func mustStop(err error, msgKey string) bool {
	if err == nil {
		return false
	}

	cmd.Confirm(err, languages.Get(msgKey))
	cmd.PressEnter()
	return true
}

func measureProjectInit(settings project.Settings) {
	utils.MeasureTimeErr(
		languages.Get("init"),
		func() error {
			return project.Init(settings)
		},
	)
}

func initLanguage(lang string) {
	var err error

	if lang != "" {
		err = languages.Init(lang)
	} else {
		err = languages.AutoInit()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f := parseFlags()

	// сначала язык
	initLanguage(f.Lang)

	if utils.FileExists(config.BlockadeFileName) {
		cmd.Confirm(
			fmt.Errorf(languages.Get("error_blockade_found")),
			config.BlockadeMessage,
		)
		cmd.PressEnter()
		return
	}

	// Non-ASCII path can break npm init
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
		if cmd.Ask(languages.Get("ask_clear_current_dir")) {
			if err := project.Remove(); err != nil {
				mustStop(err, "error_remove_current_dir")
			}
		}
		return
	}

	// init mode
	settings := project.Settings{
		UseGit:       f.UseGit,
		CreateReadme: !f.NoReadme,
	}

	if f.Init {
		measureProjectInit(settings)
		return
	}

	if cmd.Ask(languages.Get("ask_create_project")) {
		measureProjectInit(settings)
	}
}
