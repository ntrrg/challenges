#!/bin/sh
# HackerRank challenges tester.

set -e

on_error() {
  return 1
}

trap on_error INT TERM EXIT

main() {
  case $1 in
    -h | --help )
      printHelp
      return
      ;;

    -t | --test )
      testCases=$2
      shift; shift
      ;;
  esac

  challenges="$@"

  challenges=$(
    find "${challenges:-HackerRank}" -type f -name "solution.*" |
    sort |
    sed "s/ /_ESPACE_/g"
  )

  for challenge in $challenges; do
    challenge="$(dirname "$(echo "$challenge" | sed "s/_ESPACE_/ /g")")"
    echo
    echo "$challenge"
    cd "$challenge"

    if [ -z "$testCases" ]; then
      testCases="$(find "input" -name "input??.txt")"
    fi

    for testCase in $testCases; do
      testCase="$(echo "$testCase" | grep -oe "[0-9][0-9]")"
      echo -n "  Test case $testCase: "
      run $testCase
    done

    cd "$OLDPWD"
  done
}

run() {
  input="input/input$1.txt"
  output="output/output$1.txt"

  if [ ! -f "$input" ]; then
    echo "  The test case $testCase doesn't exists"
    return 1
  fi

  echo "$(cat "${output}")" > /tmp/EXPECT
  echo "$(cat "${input}" | go run solution.go)" > /tmp/RESULT
  diff /tmp/EXPECT /tmp/RESULT && echo "PASS" || echo "FAIL"
}

printHelp() {
  cat <<EOF
Usage: $0 [-t testCase] [PATH]

Arguments:

  PATH
    Challenges folder, searches solutions recursively. (default: HackerRank)

Options:

  -h, --help
    Shows this help text and exit.

  -t, --test TEST_CASE
    Runs 'solution.*' against TEST_CASE test case. e.g. 00, 01.
EOF
}

main $@