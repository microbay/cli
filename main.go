package main

import (
	"fmt"
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
			Usage:   "JWT access token base command",
			Subcommands: []cli.Command{
				{
					Name:  "sign",
					Usage: "sign a token",
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
							return
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
				{
					Name:  "verify",
					Usage: "verfiy a token",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "token, t",
							Value: "",
							Usage: "JWT token",
						},
						cli.StringFlag{
							Name:  "key, k",
							Value: "",
							Usage: "public key",
						},
					},
					Action: func(c *cli.Context) {
						input := c.String("token")
						keyPath, _ := filepath.Abs(c.String("key"))
						key, err := ioutil.ReadFile(keyPath)
						if err != nil {
							println("Failed loading public key file", keyPath, err)
							return
						}
						token, err := jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
							// Don't forget to validate the alg is what you expect:
							if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
								return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
							}
							return key, nil
						})

						if err == nil && token.Valid {
							println("Successfuly verfied token", input)
						} else {
							println("Failed")
						}
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
