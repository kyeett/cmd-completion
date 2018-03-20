package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

// https://godoc.org/github.com/mitchellh/cli
func main() {
	app := cli.NewCLI("hello", "0.0.0")
	app.Args = os.Args[1:]
	app.HiddenCommands = []string{}
	app.AutocompleteInstall = "magnus"
	app.Commands = map[string]cli.CommandFactory{
		"hello sub1": func() (cli.Command, error) {
			return &Hello{}, nil
		},

		"build slow subsub1": func() (cli.Command, error) {
			return &Hello{}, nil
		},

		"build fast": func() (cli.Command, error) {
			return &Hello{}, nil
		},
	}

	status, err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}

/*
https://github.com/mitchellh/cli/blob/518dc677a1e1222682f4e7db06721942cb8e9e4c/cli.go

   // Check if it implements ComandAutocomplete. If so, setup the autocomplete
   if c, ok := impl.(CommandAutocomplete); ok {
       subCmd.Args = c.AutocompleteArgs()
       subCmd.Flags = c.AutocompleteFlags()
   }

*/

type Hello struct {
}

func (*Hello) Help() string {
	return "hello"
}
func (*Hello) Run(args []string) int {
	fmt.Printf("hello, %v", args)
	return 0
}
func (h *Hello) Synopsis() string {
	return h.Help()
}

//https://blog.alexellis.io/golang-json-api-client/

type people struct {
	Number int `json:"number"`
}

func (h *Hello) AutocompleteArgs() complete.Predictor {
	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// var x map[string]interface{}
	// jsonErr := json.Unmarshal(body, &x)
	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(people1)

	return complete.PredictSet("context", "gotypes", "netipv6zone", "printerconfig")
}

// http://httpbin.org/headers
func (h *Hello) AutocompleteFlags() complete.Flags {
	return complete.Flags{}
}
