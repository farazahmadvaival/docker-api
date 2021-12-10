package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	var listenAddress = "0.0.0.0:8088"
	gin.SetMode(gin.ReleaseMode)
	if os.Args != nil && len(os.Args) > 0 {
		listenAddress = os.Args[1]
	}
	log.Printf("server is listening at %s...", listenAddress)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	router.GET("/containers", func(c *gin.Context) {
		containers, err := getContainers()
		if err != nil {
			panic(err)
		}
		var list []types.Container
		for _, container := range containers {
			list = append(list, container)
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    list,
			"message": "OK",
		})
	})

	router.GET("/restart-container/:container_id", func(c *gin.Context) {
		containerId := c.Param("container_id")
		err := restartContainer(containerId)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": "OK",
		})
	})

	router.Run(listenAddress) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getContainers() ([]types.Container, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return containers, nil
}

func restartContainer(containerId string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerRestart(ctx, containerId, nil); err != nil {
		panic(err)
	}

	return nil
}
