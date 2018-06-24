package client

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/kamontat/forgitgo/client/validation"
	"github.com/kamontat/forgitgo/utils"
	git "gopkg.in/src-d/go-git.v4"
)

type Runner struct {
	Dir  string
	Repo *git.Repository
	Tree *git.Worktree
}

func BeforeRun() Runner {
	dir, err := os.Getwd()
	if err != nil {
		utils.Logger().WithError(err).PathError("Cannot get current path")
	}

	return Runner{
		Dir: dir,
	}
}

func (runner Runner) SetRepository() Runner {
	err := validation.HasGit(runner.Dir)
	if err == nil {
		return runner.OpenRepository()
	}
	return runner.InitRepository()
}

func (runner Runner) InitRepository() Runner {
	repo, err := git.PlainInit(runner.Dir, false)
	if err != nil {
		utils.Logger().WithError(err).GitError("init", "Cannot initial repository")
	}

	runner.Repo = repo
	return runner
}

func (runner Runner) OpenRepository() Runner {
	repo, err := git.PlainOpen(runner.Dir)
	if err != nil {
		utils.Logger().WithError(err).GitError("repo", "Cannot open current git repository")
	}

	runner.Repo = repo
	return runner
}

func (runner Runner) SetWorktree() Runner {
	if runner.Repo == nil {
		runner = runner.SetRepository()
	}

	tree, err := runner.Repo.Worktree()
	if err != nil {
		utils.Logger().WithError(err).GitError("worktree", "Cannot get worktree")
	}

	runner.Tree = tree
	return runner
}

func (runner Runner) Add(pattern string) {
	if runner.Tree == nil {
		runner = runner.SetWorktree()
	}

	err := runner.Tree.AddGlob(pattern)
	if err != nil {
		utils.Logger().WithError(err).GitError("add", "Cannot add "+pattern)
	}
}

func (runner Runner) Commit(name string, email string, key string, title string, message string) {
	if runner.Tree == nil {
		runner = runner.SetWorktree()
	}

	h, err := runner.Tree.Commit(fmt.Sprintf("[%s] %s\n%s", key, title, message), &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		utils.Logger().WithError(err).GitError("worktree", "Cannot commit")
	}
	utils.Logger().Info("commit", h.String())
}

func (runner Runner) Status() git.Status {
	if runner.Tree == nil {
		runner = runner.SetWorktree()
	}

	s, err := runner.Tree.Status()
	if err != nil {
		utils.Logger().WithError(err).GitError("worktree", "Cannot show status")
	}

	return s
}
