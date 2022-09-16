// $ docker run -d alpine /bin/sh -c "while :; do sleep 10; echo hello; done"

package main

import (
	"context"
	"os"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var ID string = "5268f4c0ae3f"

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	}
	out, err := cli.ContainerLogs(ctx, ID, options)
	if err != nil {
		panic(err)
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

// hello
// hello
// hello
// hello
// hello
// ...
