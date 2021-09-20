/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/borischen0203/shoris/config"
	"github.com/spf13/cobra"
)

type UrlResponse struct {
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
	Data   struct {
		URL     string   `json:"url"`
		Domain  string   `json:"domain"`
		Alias   string   `json:"alias"`
		Tags    []string `json:"tags"`
		TinyURL string   `json:"tiny_url"`
	} `json:"data"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Shorten the URL",
	Long:  `Shorten the URL`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 && len(args) == 1 {
			longUrl := args[0]
			getShortenURL(longUrl)
		} else {
			fmt.Println("command not exist")
		}
	},
}

func getShortenURL(URL string) {
	token := config.Env.Api_token
	// token := "uYcUCgvOMB3XYGAXAUmO3uojC9TjqZHlWG8sQJ4IeoKZxlOvSou706TgTpg7"
	requestAPI := "https://api.tinyurl.com/create?api_token=" + token
	requestBody, _ := json.Marshal(map[string]string{
		"url":    URL,
		"domain": "tiny.one",
		"alias":  "",
	})

	res, err := http.Post(requestAPI,
		"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		print(err)
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		print(err)
	}
	var urlResponse UrlResponse
	json.Unmarshal([]byte(response), &urlResponse)

	if res.StatusCode == 200 {
		fmt.Println(urlResponse.Data.TinyURL)
	} else {
		fmt.Println(urlResponse.Errors[0])
	}
}

func init() {
	// config.Setup()
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
