package service
 import (
	"github.com/go-martini/martini"
)

 func NewServer(port string) {
 	ns := martini.Classic()
 	//get the input name by using marini.Params and return
 	ns.Get("/hello/:name", func(params martini.Params) string {
 		return params["name"] +" go to sleep!!!!"
 		})
 	//run
 	ns.RunOnAddr(":" + port)
 }
