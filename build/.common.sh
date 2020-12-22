GOOS="linux"
VERSION=$(cat "$(dirname "${BASH_SOURCE[0]}")/../VERSION")


if [ "$(uname -m)" == "x86_64" ]; then
  GOARCH="amd64"
fi

NAME="dnsmonitor-${GOARCH}"
REGISTRIES=(
            "ghcr.io/felixfriedrich/"
            )

export GOOS
export GOARCH
export REGISTRIES
export VERSION
