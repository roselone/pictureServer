package controllers

import "github.com/revel/revel"
import "strconv"
import "io/ioutil"

type App struct {
	*revel.Controller
}

func (c App) Index(i string) revel.Result {
	var index int
	var pictures string
	buf,err := ioutil.ReadFile("conf")
	if err != nil {
		pictures = err.Error()	
	}else{
		pictures = string(buf)
	}
	if i == "" {
		index = 0
	}else {
		index,_ = strconv.Atoi(i)
	}
	return c.Render(index,pictures)
}

