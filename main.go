package main

import (
	"fmt"

	"github.com/sumit-behera-in/ginBaseServer/db"
)

func main() {
	postgres := "postgresql://postgres:admin@localhost:5432/gin_http_register?sslmode=disable"

	c, err := db.New(postgres)
	if err != nil {
		panic(err)
	}

	defer c.Close()

	// err = c.AddCluster(db.ClusterInfo{})
	// if err != nil {
	// 	panic(err)
	// }

	var n int
	n, _ = c.CountTotalClusters()
	println("total", n)

	k, _ := c.GetAllClusters()
	fmt.Println(k)

	x, _ := c.GetClusterInfoById(10)
	fmt.Println(x)

}
