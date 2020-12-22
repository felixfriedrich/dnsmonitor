GOOS=$(uname | tr '[:upper:]' '[:lower:]')

if [ "$(uname -p)" == "x86_64" ]; then
  GOARCH="amd64"
fi

export $GOOS
export $GOARCH