package main

import (
	"hr_manajemen/routes"


)

var err error

func main() {
	r := routes.SetupRouter()
	//running
	// add comment from eko
	r.Run(":7200")
}
