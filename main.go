package main

import (
	"mcd-clone/cli"
	_ "mcd-clone/storage"
	"mcd-clone/utils"
)

func main() {
	utils.ClearTerm(0, "")
	cli.HomeMenu(0)
}
