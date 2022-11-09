package src

import (
	"fmt"

	"strconv"

	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

const path = "."

func ReadDialogue() string {
	//path, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	dialogue := viper.New()
	dialogue.AddConfigPath(path + "/Resources/Dialogue")
	dialogue.SetConfigName("Hello")
	dialogue.SetConfigType("yaml")
	if err := dialogue.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return dialogue.GetString("Contents")
}

func ReadMission() string {
	//path, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//}
	mission := viper.New()
	mission.AddConfigPath(path + "/Resources/Missions")
	mission.SetConfigName("CTM-Mission")
	mission.SetConfigType("yaml")
	if err := mission.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return mission.GetString("Mission.Contents")
}

func ReadOptions(content string) []string {
	var options []string
	var checker []string
	//content, err := os.ReadFile(path + "/Resources/Dialogue/Hello.json")
	OptionsNum, _ := strconv.Atoi(gjson.Get(content, "Options.#").String())
	for i := 0; i < OptionsNum; i++ {
		options = append(options, gjson.Get(content, "Options."+strconv.Itoa(i)+".Option-"+strconv.Itoa(i+1)+".Contents").String())
		checker = append(checker, DialogueChecker(gjson.Get(content, "Options."+strconv.Itoa(i)+".Option-"+strconv.Itoa(i+1)).String()))
	}
	return options
}
