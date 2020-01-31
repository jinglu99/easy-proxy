package main

import (
	"github.com/jingleWang/easy-proxy/src/app"
	"os"
)

func main() {
	a := app.MainMenu()
	err := a.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}

