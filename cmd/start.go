/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"os/signal"

	"github.com/bketelsen/gcbot/bot"
	"github.com/spf13/cobra"
)

// GopherGuild is the constant id of the guild
const GopherGuild = "755435423177638059"

// LogChannel is the constant id of the logging channel for bots
const LogChannel = "765966830380515389"

// HelpDesk is the constant id of the help desk channel
const HelpDesk = "763762406417498142"

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start your bot",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		// get configs from env and constants
		config := bot.NewConfig("$DISCORD_TOKEN", LogChannel, GopherGuild)
		//create a new discord connection
		discord := bot.NewDiscord(config)

		// make a channel for signals
		sig := make(chan os.Signal, 1)

		// channel for telling the bot to hang up
		botQuit := make(chan bool, 1)
		signal.Notify(sig, os.Interrupt)
		go func() {
			<-sig
			// when the os signals are caught, send a message to the bot channel
			// so it will quit
			botQuit <- true
			os.Exit(0)
		}()

		fmt.Println("kicking off the bot")
		if err := discord.Run(botQuit); err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
