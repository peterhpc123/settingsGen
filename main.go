package main

import (
	"awesomeProject/cmd"
	_ "awesomeProject/cmd"
	_ "github.com/beevik/etree"
)

func main() {
	//construct CLI
	cmd.Execute()

}
