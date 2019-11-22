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
      TEST_CASES_FLAG="input/input$2.txt"
      shift; shift
      ;;
  esac

  ERRORS=0
  CHALLENGES="$(find "${1:-$CHALLENGES_DIR}" -type f -name ".env" | sort)"

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

    TEST_CASES="${TEST_CASES_FLAG:-$(find "input" -name "input??.txt" | sort)}"

    for TEST_CASE in $TEST_CASES; do
      TEST_CASE="$(echo "$TEST_CASE" | grep -o "[0-9][0-9]")"
      INPUT="input/input$TEST_CASE.txt"
      OUTPUT="output/output$TEST_CASE.txt"

      if [ ! -f "$INPUT" ] || [ ! -f "$OUTPUT" ]; then
        continue
      fi

      printf "%s  * Test case %s: " "$PREFIX" "$TEST_CASE"
      run "$LANGUAGE" "$INPUT" "$OUTPUT" || ERRORS=$?
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

  return 0
}

run() {
  local GOT=""
  local WANT=""

  case $1 in
    go )
      go build -o solution main.go
      GOT="$(cat "$2" | ./solution)"
      ;;

    * )
      echo "Unsupported language '$1'"
      return 1
      ;;
  esac

  WANT="$(cat "$3")"

  if [ "$GOT" != "$WANT" ]; then
    echo "[FAIL]\nGot:\n$GOT\nWant:\n$WANT"
    return 1
  fi

  echo "[PASS]"
  return 0
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
TEST_CASES_FLAG=""

main "$@"

