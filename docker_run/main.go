// https://haibara-works.hatenablog.com/entry/2020/12/05/235227

package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	// イメージプル
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader) // ログを出力

	// コンテナを作成
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd: []string{"echo", "hello world"},
		Tty: false,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	// コンテナを起動
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// 処理を待つ
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// 実行ログ出力
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out) // ログを出力
}
