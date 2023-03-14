package cmd

import (
	"bufio"
	"fmt"
	"github.com/beevik/etree"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"os"
	"strings"
)

var userName string

var apiKey string

var url string

var github string
var cli = &cobra.Command{
	Use:     "cli",
	Example: "example",
	Short:   "short",
	Long:    "long",
	Run: func(cmd *cobra.Command, args []string) {
		userHome, _ := os.UserHomeDir()
		fmt.Printf(userHome)
		fmt.Printf("Please enter file in which to save the key (%s): ", userHome+"\\.m2\\settings.xml")

		locationReader := bufio.NewReader(os.Stdin)
		location, _ := locationReader.ReadString('\n')
		//TODO: download a template file if location hasn't settings.xml
		if _, err := os.Stat(location); err != nil {
			downloadFromGitHub(github, location)
		}

		//fmt.Printf("%s\n", location)

		for count := 1; ; count++ {
			username, password := input(userHome)
			// access maven repo with auth
			if apiKey = access(username, password, url); apiKey != "" { //登录成功
				break
			} else {
				fmt.Printf("please input with no more 3 times")
				if count == 4 {
					os.Exit(1)
				}
			}
		}

		//assume this is a true template file
		parse(location)

	},
}

func downloadFromGitHub(g string, location string) {

}

func parse(location string) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(location); err != nil {
		return err.Error()
	}
	settings := doc.SelectElement("settings")
	//TODO: get each server and put it in to map[key:serverid, vlaue:serverNode]
	if settings != nil {
		repos := settings.FindElements("./profiles/profile")
		for range repos {
			//TODO: get each repo and judge if it's repo id exists in server map, if not then create a server node, or replace it with api key
		}
	}
	return nil
}

func input(userHome string) (string, string) {
	paths := strings.Split(userHome, "\\")
	fmt.Printf("Enter userName (%s):", paths[len(paths)-1])
	usernameReader := bufio.NewReader(os.Stdin)
	username, _ := usernameReader.ReadString('\n')
	if username == "" {
		username = paths[len(paths)-1]
	}
	fmt.Printf("Enter passphrase:")
	passwordReader := bufio.NewReader(os.Stdin)
	password, _ := passwordReader.ReadString('\n')
	return username, password
}

func access(username string, password string, url string) string {
	//TODO: http get api key
	return "apikey"
}

func Execute() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
