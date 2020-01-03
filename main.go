package main

import (
	"os"
	"net"
	"github.com/gin-gonic/gin"
)

func Ips() (map[string]string) {

    ips :=  make(map[string]string)

    interfaces, err := net.Interfaces()
    if err != nil {
        return nil
    }

    for _, i := range interfaces {
        byName, err := net.InterfaceByName(i.Name)
        if err != nil {
            return nil
        }
        addresses, err := byName.Addrs()
        for _, v := range addresses {
            ips[byName.Name] = v.String()
        }
    }
    return ips
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
                hostname, _ := os.Hostname()

		c.JSON(200, gin.H{
			"hostname": hostname,
			"ip": Ips(),	
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
