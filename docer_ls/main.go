// https://zenn.dev/okaki_se/articles/c81c4dbf672a34

package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	all()
	project_filter()
	nest()
}

func all() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	// コンテナ一覧
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	fmt.Println("all: ", containers)
}

func project_filter() {
	// フィルタ条件を追加する
	// プロジェクト名testを指定してコンテナを起動しておく
	// docker-compose -p test up -d

	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	labelFilters := filters.NewArgs()
	labelFilters.Add("label", "com.docker.compose.project=test")
	containers_filtered, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
		Filters: labelFilters,
	})
	fmt.Println("project_filter: ", containers_filtered)
}

type ServiceContainerMap map[string][]types.Container

func nest() (ServiceContainerMap, error){
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		return nil, err
	}

	// ラベルありのコンテナを取得
	labelFilters := filters.NewArgs()
	labelFilters.Add("label", "com.docker.compose.project") // ラベルがついているものすべて
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
		Filters: labelFilters,
	})
	if err != nil {
		return nil, err
	}
	// fmt.Println(containers)
	// return nil, err

	// プロジェクトでグループ化
	containersByProject := ServiceContainerMap{}
	for _, container := range containers {
		label := container.Labels["com.docker.compose.project"]
		cons := containersByProject[label]
		if cons == nil {
			cons = []types.Container{}
		}
		cons = append(cons, container)
		containersByProject[label] = cons
	}
	fmt.Println("grouping: ", containersByProject)
	return containersByProject, nil
}
