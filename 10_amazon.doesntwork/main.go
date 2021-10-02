package main

import (
	"fmt"
	"github.com/DDRBoxman/go-amazon-product-api"
)

// 2021-08 Doesn't work. I haven't found a good Golang solution for AWS
// Furthermore, it looks like Amazon doesn't allow low-traffic clients, it's
// meant for active affiliates only.
func main() {
	var api amazonproduct.AmazonProductAPI

	api.AccessKey = ""
	api.SecretKey = ""
	api.Host = "webservices.amazon.com"
	api.AssociateTag = ""

	result,err := api.ItemSearchByKeyword("sgt+frog")
	if (err != nil) {
        fmt.Println(err)
	}

	fmt.Println(result)
}
