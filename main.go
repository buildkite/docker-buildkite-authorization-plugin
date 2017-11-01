package main

import (
	"flag"
	"log"
	"os/user"
	"strconv"

	"github.com/docker/go-plugins-helpers/authorization"
)

const (
	pluginSocket = "/run/docker/plugins/buildkite-docker.sock"
)

func main() {
	flag.Parse()

	log.Printf("Listening for authz...")
	pl, err := NewTestPlugin()
	if err != nil {
		log.Fatal(err)
	}

	h := authorization.NewHandler(pl)

	u, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(u.Gid)

	if err := h.ServeUnix(pluginSocket, gid); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening...")
}
