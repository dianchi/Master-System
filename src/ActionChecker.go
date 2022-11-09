package src

import (
	"github.com/tidwall/gjson"
)

func DialogueChecker(option string) string {
	var GoDialogue string
	if gjson.Get(option, "Action.Go").String() == "Dialogue" {
		GoDialogue = gjson.Get(option, "Action.Contents").String()
	}
	return GoDialogue
}
