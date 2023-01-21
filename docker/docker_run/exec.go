package main

import (
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"golang.org/x/net/context"
)

var (
	ID string = "6c6076228da6"
	cmd []string = []string{"ls", "/"}
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Tty: true,
		Cmd: cmd,
	}

	IDResp, err := cli.ContainerExecCreate(ctx, ID, config)
	if err != nil {
		log.Panic(err)
	}

	resp, err := cli.ContainerExecAttach(ctx, IDResp.ID, types.ExecStartCheck{})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := resp.Conn.Close(); err != nil {
			log.Panic(err)
		}

		log.Println("connection closed")
	}()

	stdcopy.StdCopy(os.Stdout, os.Stderr, resp.Reader)
}
