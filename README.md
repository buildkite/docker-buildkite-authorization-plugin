# Buildkite Job Authorization Docker Plugin

At present this is a placeholder to debug an issue getting authorization plugins working with Docker for Mac.

## Running

```bash
mkdir -p $GOPATH/github.com/buildkite
git clone https://github.com/buildkite/docker-buildkite-authorization-plugin.git docker-buildkite-authorization-plugin
cd $GOPATH/github.com/buildkite/docker-buildkite-authorization-plugin
make all enable
```

You can now verify it's installed and working with `docker plugin list`.

Running `docker run docker/whalesay cowsay boo` should fail.
