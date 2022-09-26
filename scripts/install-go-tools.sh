#!/bin/bash

GOPATH=$(go env GOPATH)
GOBINDIR="${GOPATH}/bin"

install_gometalinter() {
  GMLVER="${GO_METALINTER_VERSION}"
  if type gometalinter >/dev/null 2>&1; then
    echo "gometalinter already installed"
    return
  fi

  echo "Installing gometalinter. Version: ${GMLVER}"
  curl -L https://raw.githubusercontent.com/alecthomas/gometalinter/master/scripts/install.sh | bash -s -- -b "${GOBINDIR}" "${GMLVER}"
}

install_gometalinter
