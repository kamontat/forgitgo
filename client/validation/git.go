package validation

import (
	"errors"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

func HasGit(dir string) (err error) {
	_, err = os.Stat(dir + "/.git")
	return
}

func HasRemote(repo *git.Repository) (err error) {
	remotes, err := repo.Remotes()
	if err != nil {
		return
	}

	if len(remotes) < 1 {
		return errors.New("remote not exist")
	}

	return
}
