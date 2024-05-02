package main

import (
	"DischordGptBot/bot"

	"github.com/labstack/echo"
)


func main(){
	

	e:=echo.New()
	


   bot.Start()

	e.Start(":8080")

}