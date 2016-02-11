package hanlonlib

import (
	"github.com/codegangsta/cli"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
)

func ImageCommands() cli.Command {
	command := cli.Command{
		Name: "image",
		Usage: "Show available images",
		Action: printAvailableImages,
		Subcommands: []cli.Command {
			{
				Name: "remove",
				Usage: "Remove an image",
				Action: func(c *cli.Context) {
					println("removed image")
				},
			},
			{
				Name: "add",
				Usage: "Add an image",
				Action: func(c *cli.Context) {
					println("Add image")
				},
			},
		},
	}
	return command
}

type Image struct {
	Uuid string `json:"@uuid"`
	Classname string `json:"@classname"`
	Noun string `json:"@noun"`
	Uri string `json:"@uri"`
}

type ImageResponse struct {
	Resource string `json:"resource"`
	Command string `json:"command"`
	Result string `json:"result"`
	Http_err_code int `json:"http_err_code"`
	Errcode int `json:"errcode"`
	Response []Image `json:"response"`
}

func printAvailableImages(c *cli.Context)  {
	var image ImageResponse
	resp, err := http.Get("http://10.53.252.60:8026/hanlon/api/v1/image")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		err = json.Unmarshal(contents, &image)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, resp := range image.Response {
			fmt.Println(resp.Uri)
		}

	}
}