package client

import (
	"github.com/kamontat/forgitgo/utils"
	"github.com/manifoldco/promptui"
)

func PromptKey(key *string, value []string, index int, commitDB []Commit) {
	if len(value) <= index {
		var s []string
		for _, v := range commitDB {
			s = append(s, v.String())
		}

		prompt := promptui.SelectWithAdd{
			Label:    "Select commit key",
			Items:    s,
			AddLabel: "custom",
		}

		_, result, err := prompt.Run()
		if err != nil {
			utils.Logger().WithError(err).PromptError("cause error while run prompt")
		}
		*key = result
	}
}

func PromptTitle(title *string, value []string, index int) {
	if len(value) <= index {
		prompt(title, "title")
	}
}

func PromptMessage(msg *string, value []string, index int) {
	if len(value) <= index {
		prompt(msg, "message")
	}
}

func prompt(result *string, message string) {
	prompt := promptui.Prompt{
		Label: "Enter " + message,
	}

	temp, err := prompt.Run()
	if err != nil {
		utils.Logger().WithError(err).PromptError("cause error while run prompt")
	}

	*result = temp
}
