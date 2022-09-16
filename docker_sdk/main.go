package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	_"github.com/docker/docker/api/types/filters"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID[:12], container.Image)
	}
}

// 9e616abed730 ubuntu
// 3f72d5bbdba4 golang
// 15c41393877f golang
// 6c6076228da6 ubuntu
// 660f64d02c15 ubuntu
