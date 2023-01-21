package main

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	dockerfileName string = "Dockerfile-test"
	imageNameAndTags []string = []string{"name:tag"}
)

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	buildOptions := types.ImageBuildOptions{
		Dockerfile: dockerfileName,
		Remove: true,
		Tags: imageNameAndTags,
	}
	res, err := cli.ImageBuild(
		ctx,
		getArchivedDockerfile(dockerfileName),
		buildOptions,
	)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func getArchivedDockerfile(dockerfile string) *bytes.Reader {
	// read the Dockerfile
	f, err := os.Open(dockerfile)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panic(err)
	}

	// archive the Dockerfile
	// SDKにDockerfileの内容を渡すとき、tarで圧縮する必要がある
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(b)),
	}
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Panic(err)
	}
	_, err = tw.Write(b)
	if err != nil {
		log.Panic(err)
	}

	return bytes.NewReader(buf.Bytes())
}

// go run dockerfile.go
// {"stream":"Step 1/2 : FROM alpine:3.8"}
// {"stream":"\n"}
// {"status":"Pulling from library/alpine","id":"3.8"}
// {"status":"Pulling fs layer","progressDetail":{},"id":"486039affc0a"}
// {"status":"Downloading","progressDetail":{"current":23273,"total":2207101},"progress":"[\u003e                                                  ]  23.27kB/2.207MB","id":"486039affc0a"}
// {"status":"Verifying Checksum","progressDetail":{},"id":"486039affc0a"}
// {"status":"Download complete","progressDetail":{},"id":"486039affc0a"}
// {"status":"Extracting","progressDetail":{"current":32768,"total":2207101},"progress":"[\u003e                                                  ]  32.77kB/2.207MB","id":"486039affc0a"}
// {"status":"Extracting","progressDetail":{"current":2207101,"total":2207101},"progress":"[==================================================\u003e]  2.207MB/2.207MB","id":"486039affc0a"}
// {"status":"Pull complete","progressDetail":{},"id":"486039affc0a"}
// {"status":"Digest: sha256:2bb501e6173d9d006e56de5bce2720eb06396803300fe1687b58a7ff32bf4c14"}
// {"status":"Status: Downloaded newer image for alpine:3.8"}
// {"stream":" ---\u003e c8bccc0af957\n"}
// {"stream":"Step 2/2 : RUN /bin/sh -c 'echo \"hello world!\"'"}
// {"stream":"\n"}
// {"stream":" ---\u003e Running in 264c3a75f2b6\n"}
// {"stream":"hello world!\n"}
// {"stream":"Removing intermediate container 264c3a75f2b6\n"}
// {"stream":" ---\u003e 9b2532bb3bc6\n"}
// {"aux":{"ID":"sha256:9b2532bb3bc652ee85a8feb396b1690217f905af1f1157db7c48138604359dcc"}}
// {"stream":"Successfully built 9b2532bb3bc6\n"}
// {"stream":"Successfully tagged name:tag\n"}
