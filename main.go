// DOCKER_API_VERSION=1.42 go run main.go

package main

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	argsWithProg := os.Args

	if len(argsWithProg) < 2 {
		fmt.Println("Need container ID")
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if argsWithProg[1] == container.ID[:12] {
			syscall.Exec("/usr/local/bin/docker", []string{"docker", "exec", "-it", container.ID[:10], "bash"}, os.Environ())
		} else {
			fmt.Println("Container does't exists.")
		}
	}
}