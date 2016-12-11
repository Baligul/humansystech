package main

import (
    "os"
    "strconv"
    
	_ "github.com/Baligul/humansystech/routers"
	"github.com/astaxie/beego"
)

func main() {
    port := os.Getenv("PORT")
    
    if port == "" {
        port = "8088"
    }
        
    beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)
    
	beego.Run()
}