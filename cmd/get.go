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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
type UrlResponse struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
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
	requestAPI := "https://short-url-sample.herokuapp.com/api/url-shortener/v1/url"
	requestBody, _ := json.Marshal(map[string]string{
		"longUrl": URL,
	})

	res, err := http.Post(requestAPI,
		"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		print(err)
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if res.StatusCode == 200 {
		if err != nil {
			print(err)
		}
		var urlResponse UrlResponse
		json.Unmarshal([]byte(response), &urlResponse)
		fmt.Println(urlResponse.ShortUrl)
	} else {
		fmt.Println(string(response))
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
