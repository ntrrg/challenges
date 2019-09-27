#!/bin/sh
# Copyright (c) 2019 Miguel Angel Rivera Notararigo
# Released under the MIT License

set -e

main() {
  case $1 in
    -h | --help )
      show_help
      return 0
      ;;

    -t | --test )
      TEST_CASES="input/input$2.txt"
      shift; shift
      ;;
  esac

  CHALLENGES="$(find "${1:-$CHALLENGES_DIR}" -type f -name ".env")"

  for CHALLENGE in $CHALLENGES; do
    DIR="$(dirname "$CHALLENGE")"
    PREFIX="$(get_prefix "$DIR")"
    NAME=""
    LANGUAGE=""

    . "$CHALLENGE"

    printf "%s%s\n" "$PREFIX" "$NAME"

    if [ -z "$LANGUAGE" ]; then
      continue
    fi

    cd "$DIR"

    TEST_CASES="${TEST_CASES:-$(find "input" -name "input??.txt")}"

    for TEST_CASE in $TEST_CASES; do
      TEST_CASE="$(echo "$TEST_CASE" | grep -o "[0-9][0-9]")"
      printf "%s  * Test case %s: " "$PREFIX" "$TEST_CASE"
      (run "$TEST_CASE" && echo "[PASS]") || echo "[FAIL]"
    done

    cd "$OLDPWD"
  done
}

get_prefix() {
  COUNT=$(( ($(echo "$1" | tr "/" "\n" | wc -l) - 2) * 2 ))

  while [ $COUNT -gt 0 ]; do
    printf " "
    COUNT=$(( $COUNT - 1 ))
  done

  return 0
}

run() {
  TEST_CASE="$1"
  GOT="$(cat "input/input$TEST_CASE.txt" | go run main.go)"

  printf "%s" "$GOT" | diff "output/output$TEST_CASE.txt" -
}

show_help() {
  cat <<EOF
$0 - Challenges runner.

Usage: $0 [OPTIONS] [PATH]

Arguments:
  PATH
    Challenges folder, looks for solutions recursively. ($CHALLENGES_DIR)

Options:
  -h, --help
    Show this help text and exit.
  -t, --test=TEST_CASE
    Run 'solution.*' against TEST_CASE test case. e.g. 00, 01.
EOF

  return 0
}

CHALLENGES_DIR="challenges"
TEST_CASES=""

main "$@"

