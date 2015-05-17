package main

import (
	"github.com/codegangsta/cli"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	app := cli.NewApp()
	app.Name = "Apigo CLI"
	app.Usage = "CLI for Apigo API gateway"
	app.Action = func(c *cli.Context) {
		println("Welcome to " + app.Name + ". Type 'apigo-cli help'")
	}

	app.Commands = []cli.Command{
		{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "Creates a JWT access token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Value: "",
					Usage: "private key file",
				},
			},
			Action: func(c *cli.Context) {
				keyPath, _ := filepath.Abs(c.String("key"))
				key, err := ioutil.ReadFile(keyPath)
				if err != nil {
					println("Failed loading private key file", keyPath, err)
				}
				token := jwt.New(jwt.SigningMethodHS256)
				token.Claims["foo"] = "bar"
				tokenString, err := token.SignedString(key)
				if err != nil {
					println("Failed creating token", err)
				} else {
					println(tokenString)
				}
			},
		},
	}

	app.Run(os.Args)
}
