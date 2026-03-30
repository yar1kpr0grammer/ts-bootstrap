package main

import (
	"tsBootstrup/internal/cmd"
	"tsBootstrup/internal/npm"
	"tsBootstrup/internal/project"
	"tsBootstrup/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("tsBootstrap")

	header := widget.NewLabel("tsBootstrap")

	content := container.NewVBox(header)

	loader := widget.NewProgressBarInfinite()
	loader.Hide()
	content.Add(loader)

	ui := func(f func()) {
		fyne.Do(f)
	}

	setLoading := func(v bool) {
		ui(func() {
			if v {
				loader.Show()
			} else {
				loader.Hide()
			}
			content.Refresh()
		})
	}

	if err := utils.CheckPathForNPM(); err != nil {
		content.Add(widget.NewLabel("Path error"))
	}

	if v, _ := project.IsEmpty(); !v {
		btn := widget.NewButton("Очистить папку", func() {
			setLoading(true)

			go func() {
				project.Remove()
				setLoading(false)
			}()
		})

		content.Add(btn)
	}

	if !project.Exists() {
		createReadMe := true
		useGit := false

		createRMBox := widget.NewCheck("README", func(v bool) {
			createReadMe = v
		})

		useGitBox := widget.NewCheck("Git", func(v bool) {
			useGit = v
		})

		btn := widget.NewButton("Создать проект", func() {
			setLoading(true)

			go func() {
				settings := project.Settings{
					UseGit:       useGit,
					CreateReadme: createReadMe,
				}

				project.Init(settings)

				setLoading(false)
			}()
		})

		content.Add(container.NewVBox(createRMBox, useGitBox, btn))

	} else {
		runBtn := widget.NewButton("Запустить", func() {
			setLoading(true)

			go func() {
				npm.RunProject(cmd.ShowAll)
				setLoading(false)
			}()
		})

		content.Add(runBtn)
	}

	w.SetContent(content)
	w.ShowAndRun()
}
