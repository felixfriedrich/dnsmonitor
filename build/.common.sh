GOOS=$(uname | tr '[:upper:]' '[:lower:]')

if [ "$(uname -p)" == "x86_64" ]; then
  GOARCH="amd64"
fi

NAME="dnsmonitor-${GOOS}-${GOARCH}"
REGISTRIES=(
            "docker.pkg.github.com/felixfriedrich/dnsmonitor/"
            "ghcr.io/felixfriedrich/"
            )

export GOOS
export GOARCH
export REGISTRIES
