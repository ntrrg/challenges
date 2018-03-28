#!/bin/sh

OLD_PATH="${PATH}"
RELEASE=${1:-1.10}
GOENVS=${GOENVS:-${HOME}/.local/share/go}
TARGET="${GOENVS}/go${RELEASE}.linux-amd64"

clean_env() {
  unset GOENVS PACKAGE RELEASE TARGET
  trap - INT TERM EXIT
}

on_error() {
  clean_env
  export PATH="${OLD_PATH}"
  unset OLD_PATH TARGET

  return 1
}

trap on_error INT TERM EXIT

if [ "${GOROOT}" = "${TARGET}" ]; then
  echo "Go v${RELEASE} is active"
else
  export GOROOT="${TARGET}"

  if [ ! -d "${GOROOT}" ]; then
    echo "Downloading Go v${RELEASE}.."
    echo

    PACKAGE="go${RELEASE}.linux-amd64.tar.gz"

    (cd "/tmp" && wget -c "https://dl.google.com/go/${PACKAGE}") &&
    mkdir -p "${GOENVS}" &&
    (cd "${GOENVS}" && tar -xf "/tmp/${PACKAGE}" && mv go "${TARGET}")

    echo
    echo "Done"
  fi

  export PATH="${GOROOT}/bin:${PATH}"

  if [ -z "$GOPATH" ]; then
    echo "Where is your workspace? (~/go)"
    read GOPATH
    export GOPATH=${GOPATH:-${HOME}/go}
  fi

  echo "Go v${RELEASE} activated"
  clean_env
fi
