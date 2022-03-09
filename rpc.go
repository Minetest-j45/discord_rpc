package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type conf struct {
	ClientID     string `json:"client_id"`
	State        string `json:"state"`
	Details      string `json:"details"`
	LargeID      string `json:"large_id"`
	LargeImgText string `json:"large_img_text"`
	SmallID      string `json:"small_id"`
	SmallImgText string `json:"small_img_text"`
	Buttons      int    `json:"buttons"`
	Button1      string `json:"button_1"`
	Button1Url   string `json:"button_1_url"`
	Button2      string `json:"button_2"`
	Button2Url   string `json:"button_2_url"`
	PartyPlayers int    `json:"party_players"`
	MaxPlayers   int    `json:"max_players"`
}

func openConf() conf {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		os.Create("config.json")
		openConf()
	}
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var config conf

	err = json.Unmarshal(byteVal, &config)
	if err != nil {
		fmt.Println(err)
	}

	return config
}

func main() {
	config := openConf()

	if config.ClientID == "" {
		fmt.Println("`client_id` is required in the config.json file")
		os.Exit(1)
	}

	var buttonsArr []*client.Button
	if config.Buttons > 2 {
		fmt.Println("Too many buttons, there is a maximum of 2 buttons")
		os.Exit(1)
	} else if config.Buttons == 1 {
		buttonsArr = append(buttonsArr, &client.Button{
			Label: config.Button1,
			Url:   config.Button1Url,
		})
	} else if config.Buttons == 2 {
		buttonsArr = append(buttonsArr, &client.Button{
			Label: config.Button1,
			Url:   config.Button1Url,
		})
		buttonsArr = append(buttonsArr, &client.Button{
			Label: config.Button2,
			Url:   config.Button2Url,
		})
	}

	err := client.Login(config.ClientID)
	if err != nil {
		panic(err)
	}

	if config.PartyPlayers > 0 && config.MaxPlayers > 0 {
		now := time.Now()
		err = client.SetActivity(client.Activity{
			State:      config.State,
			Details:    config.Details,
			LargeImage: config.LargeID,
			LargeText:  config.LargeImgText,
			SmallImage: config.SmallID,
			SmallText:  config.SmallImgText,
			Party: &client.Party{
				ID:         "-1",
				Players:    config.PartyPlayers,
				MaxPlayers: config.MaxPlayers,
			},
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: buttonsArr,
		})
		if err != nil {
			panic(err)
		}
	} else {
		now := time.Now()
		err = client.SetActivity(client.Activity{
			State:      config.State,
			Details:    config.Details,
			LargeImage: config.LargeID,
			LargeText:  config.LargeImgText,
			SmallImage: config.SmallID,
			SmallText:  config.SmallImgText,
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: buttonsArr,
		})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Running...")
	fmt.Println("Press Ctrl+C to exit, whenever you want")
	for {
		time.Sleep(time.Hour * 12)
	}
}
