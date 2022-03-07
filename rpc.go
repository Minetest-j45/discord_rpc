package main

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func main() {
	err := client.Login("950420436020240415")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Powered by rich-go",
		Details:    "A very difficult game",
		LargeImage: "pfpimg",
		LargeText:  "pfp",
		SmallImage: "verifiedimg",
		SmallText:  "verified",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "My GitHub",
				Url:   "https://github.com/Minetest-j45",
			},
			&client.Button{
				Label: "My Website",
				Url:   "https://j1233.minetest.land/",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Sleeping...")
	time.Sleep(time.Hour * 1)
}
