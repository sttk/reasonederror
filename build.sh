#!/usr/bin/env bash

compile() {
  go fmt ./...
  go vet ./...
  go build ./...
}

test() {
  go test -v ./...
}

cover() {
  #go test -cover ./...

  mkdir -p coverage
  go test -coverprofile=coverage/cover.out ./...
  go tool cover -html=coverage/cover.out -o coverage/cover.html
}

bench() {
  local dir=$2
  if [[ "$dir" == "" ]]; then
    dir="."
  fi
  pushd $dir
  go test -bench . -benchmem -run=^$
  popd
}

case $1 in
test)
  test
  ;;
cover)
  cover
  ;;
bench)
  bench
  ;;
'')
  compile
  ;;
*)
  echo "Bad task: $1"
  exit 1
  ;;
esac
