package src

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ReadDialogue() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dialogue := viper.New()
	dialogue.AddConfigPath(path + "/Resources/Dialogue")
	dialogue.SetConfigName("CTM-1")
	dialogue.SetConfigType("yaml")
	if err := dialogue.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return dialogue.GetString("Contents")
}
func ReadMission() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	mission := viper.New()
	mission.AddConfigPath(path + "/Resources/Missions")
	mission.SetConfigName("CTM-Mission")
	mission.SetConfigType("yaml")
	if err := mission.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	return mission.GetString("Mission.Contents")
}
