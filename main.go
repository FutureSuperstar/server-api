package main

import (
	"encoding/base64"
	"fmt"
	"server-api/controller"
	"server-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	useMiddlewares(r)
	r.Any("/", controller.Decompress)
	r.Run(":6606")
}

func main1() {
	str := "KLUv/WAVAgUQANZcZDxAa9tiAwhGURTBYMyE1McgC5sMhoHe8KoUveqyHkGQ/9huKnmOKsID1jZvUyEgtNRCSyvxL2iRwTAM9wZPAFQAVwAYvM9DJXx9hrvKnyZvWOe7H1h/EwvqNbdxnPV3rS+GdZTXW2CSv0MEgJBCATSazxQlKukjRQB32wri9RfNbn01fD1Tr7ttN+XqvF30lXgmGxmcChw4Ghid11/NTqgYVQQy9RvD3Na4yljL7N8UoXzinFL0wGmKwCfWTyLihCAooWiiEjRNkSjrL4mBqvL3CWVQZjCCEpQgGMVEIXBSQJCZZREkHihwMHTgwHA7d9s7XckWF7A+IcIHCxk6ZrxYxhcbd/mb0MR6B1f/vZK3ne3kGhidjQROBrMx9qeZ9RbZSmT5atqGXe7voZiNqZ0vZy9/j26zQKjk7bIeGs2vhq9Hxtks2015aNh0Nmw6LmxwbDI4urJsjABCSp/6fSkVFRHY7v6klHlTcDW8/z6AuiQHGFyUJIH13aDT0CNjpHNKKMoI3BZf3GprnNmMLZT1LSBAQgRVcgPEDlBpsncCPwRwuywbO8CZgG7fpiaJeFAMHnm+nkbMl+rYoveJBFcQ+QWXQLoVj1kmM3YvNgnvIgUbZwKACz6TAPewH4J8ADaC8t0xthFqaalFIhOIurDULWs1kFInrYQowTsD"
	testBase64Decode(str)
	testZstd(str)
}

func testZstd(str string) {
	data, err := controller.DeCompressWithZstd(testBase64Decode(str))
	if err != nil {
		fmt.Println("解压失败", err.Error())
		return
	}
	fmt.Println("解压成功：", string(data))
}

func testBase64Decode(encoded string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("base解压失败", err.Error())
		return nil
	}
	return decoded
}

func useMiddlewares(r *gin.Engine) {
	r.Use(middlewares.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
