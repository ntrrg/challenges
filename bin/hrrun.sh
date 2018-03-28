#!/bin/sh

main() {
  case $1 in
    -h | --help )
      print_help
      return
      ;;

    -t | --test )
      TEST_CASE=$2
      shift; shift
      ;;
  esac

  CHALLENGES=${@:-challenges}

  find "${CHALLENGES}" \
    -type f \
    -name "solution.*" \
    -execdir test_challenge.sh "${TEST_CASE}" \;
}

print_help() {
  cat <<EOF
Usage: hrrun [-t TEST_CASE] [PATH]

Arguments:

  PATH
    Challenges folder, searches solutions recursively. (default: challenges)

Options:

  -h, --help
    Shows this help text and exit.

  -t, --test TEST_CASE
    Runs 'solution.*' against TEST_CASE test case. e.g. 00, 01.
EOF
}

main $@
