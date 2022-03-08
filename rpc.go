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
	ClientID     string `json:"clientid"`
	State        string `json:"state"`
	Details      string `json:"details"`
	LargeID      string `json:"largeid"`
	LargeImgText string `json:"largeimgtext"`
	SmallID      string `json:"smallid"`
	SmallImgText string `json:"smallimgtext"`
	Buttons      int    `json:"buttons"`
	Button1      string `json:"button1"`
	Button1Url   string `json:"button1url"`
	Button2      string `json:"button2"`
	Button2Url   string `json:"button2url"`
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
		fmt.Println("`clientid` is required in the config.json file")
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

	fmt.Println("Running...")
	fmt.Println("Press Ctrl+C to exit, whenever you want")
	for {
		time.Sleep(time.Hour * 12)
	}
}
