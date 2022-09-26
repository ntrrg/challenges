#!/bin/sh
# Copyright (c) 2019 Miguel Angel Rivera Notararigo
# Released under the MIT License

set -e

CHALLENGES_DIR="challenges"
TEST_CASES_FLAG=""

main() {
  case $1 in
    -h | --help )
      show_help
      return 0
      ;;

    -t | --test )
      TEST_CASES_FLAG="input/input$2.txt"
      shift; shift
      ;;
  esac

  local ERRORS=0

  local CHALLENGES="$(
    find "${1:-$CHALLENGES_DIR}" -type f -name ".env" |
    LC_ALL="C" sort
  )"

  local CHALLENGE=""

  for CHALLENGE in $CHALLENGES; do
    local DIR="$(dirname "$CHALLENGE")"
    local PREFIX="$(get_prefix "$DIR")"
    local NAME=""
    local LANGUAGE=""

    . "$CHALLENGE"

    printf "%s%s\n" "$PREFIX" "$NAME"

    if [ -z "$LANGUAGE" ]; then
      continue
    fi

    cd "$DIR"

    case $LANGUAGE in
      go )
        go build -o solution main.go
        ;;

      * )
        echo "Unsupported language '$LANGUAGE'"
        return 1
        ;;
    esac

    local TEST_CASES="${TEST_CASES_FLAG:-$(
      find "input" -name "input??.txt" |
      LC_ALL="C" sort
    )}"

    local TEST_CASE=""

    for TEST_CASE in $TEST_CASES; do
      TEST_CASE="$(echo "$TEST_CASE" | grep -o "[0-9][0-9]")"
      local INPUT="input/input$TEST_CASE.txt"
      local OUTPUT="output/output$TEST_CASE.txt"

      if [ ! -f "$INPUT" ] || [ ! -f "$OUTPUT" ]; then
        continue
      fi

      printf "%s  * Test case %s: " "$PREFIX" "$TEST_CASE"
      run "$INPUT" "$OUTPUT" || ERRORS=$?
    done

    cd "$OLDPWD"
  done

  return $ERRORS
}

get_prefix() {
  local COUNT=$(( ($(echo "$1" | tr "/" "\n" | wc -l) - 2) * 2 ))

  while [ $COUNT -gt 0 ]; do
    printf " "
    COUNT=$(( $COUNT - 1 ))
  done
}

run() {
  local INPUT="$1"
  local OUTPUT="$2"
  local GOT="$(cat "$INPUT" | ./solution)"

  local WANT="$(cat "$OUTPUT")"

  if [ "$GOT" != "$WANT" ]; then
    echo "[FAIL]\nGot:\n$GOT\nWant:\n$WANT"
    return 1
  fi

  echo "[PASS]"
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
}

main "$@"

