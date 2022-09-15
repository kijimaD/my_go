// https://zenn.dev/okaki_se/articles/c81c4dbf672a34

package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	_ "github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	// コンテナ一覧
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	fmt.Println(containers)
}
