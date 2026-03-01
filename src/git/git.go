package git

import "tsBootstrup/src/cmd"

func Init() error {
	err := initRepo()
	if err != nil {
		return err
	}
	cmd.Confirm(err, "init git repo")
	err = addAll()
	if err != nil {
		return err
	}
	err = commit("Init")
	if err != nil {
		return err
	}
	cmd.Confirm(err, "commit into git repo")
	return nil
}

func initRepo() error {
	return cmd.Run(cmd.ShowOnlyErrors, "git", "init")
}
func addAll() error {
	return cmd.Run(cmd.ShowOnlyErrors, "git", "add", ".")
}
func commit(message string) error {
	return cmd.Run(cmd.ShowOnlyErrors, "git", "commit", "-m", message)
}
