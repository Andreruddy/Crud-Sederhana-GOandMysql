package main

import (
	"hr_manajemen/routes"


)

var err error

func main() {
	r := routes.SetupRouter()
	//running
	r.Run(":7200")
}
