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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shoris",
	Short: "Be able to shorten the long URL",
	Long:  `Be able to shorten the long URL`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		alias := ""
		if len(args) != 0 {
			longUrl := args[0]
			if len(args) == 2 {
				alias = args[1]
			}
			getShortenURL(longUrl, alias)
		} else {
			fmt.Println("command error")
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

//getShortenURL function mainly generate a short url
func getShortenURL(URL string, alias string) {
	token := "uYcUCgvOMB3XYGAXAUmO3uojC9TjqZHlWG8sQJ4IeoKZxlOvSou706TgTpg7"
	requestAPI := "https://api.tinyurl.com/create?api_token=" + token
	requestBody, _ := json.Marshal(map[string]string{
		"url":    URL,
		"domain": "tinyurl.com",
		"alias":  alias,
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
		fmt.Print("( *・ω・)✄╰ひ╯ ")
		fmt.Printf("\x1b[34m%s\x1b[0m", urlResponse.Data.TinyURL+"\n")
		writeClip(urlResponse.Data.TinyURL)
		fmt.Println("Short URL has been saved to clipboard!")
	} else {
		fmt.Println(urlResponse.Errors[0])
	}
}

//Write data to clipboard
func writeClip(URL string) {
	cmd := exec.Command("pbcopy")
	str, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer str.Close()
		io.WriteString(str, URL)
	}()

	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shoris.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".shoris" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".shoris")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
