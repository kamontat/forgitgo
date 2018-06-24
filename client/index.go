package client

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/kamontat/forgitgo/client/validation"
	"github.com/kamontat/forgitgo/utils"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Runner struct {
	Dir  string
	Repo *git.Repository
	Tree *git.Worktree
	CMDB []Commit
}

type Commit struct {
	Name string
	Key  string
}

func (c Commit) String() string {
	return c.Name
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

func (runner Runner) SetCommitDatabase(path string) Runner {
	var commitList []Commit

	file, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger().WithError(err).GitError("commit", "Cannot get commit database")
	}

	yaml.Unmarshal(file, &commitList)

	runner.CMDB = commitList
	return runner
}

func (runner Runner) Commit(name string, email string, all bool, format string, args ...string) {
	if runner.Tree == nil {
		runner = runner.SetWorktree()
	}

	var key string
	var title string
	var message string

	PromptKey(&key, args, 0, runner.CMDB)
	PromptTitle(&title, args, 1)
	PromptMessage(&message, args, 2)

	utils.Logger().Debug("Commit option", fmt.Sprintf("user=%s <%s>, all=%t", name, email, all))
	h, err := runner.Tree.Commit(fmt.Sprintf(format, key, title, message), &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		},
		All: all,
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
