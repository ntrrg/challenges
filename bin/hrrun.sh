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

  CHALLENGES=$(
    find "${CHALLENGES}" -type f -name "solution.*" |
    sed "s/ /_ESPACE_/g"
  )

  for CHALLENGE in ${CHALLENGES}; do
    CHALLENGE="$(echo ${CHALLENGE} | sed "s/_ESPACE_/ /g")"

    cd "$(dirname "${CHALLENGE}")" &&
    test_challenge.sh "${TEST_CASE}"

    ERROR_CODE=$?

    cd ${OLDPWD}

    if [ ! ${ERROR_CODE} -eq 0 ]; then
      return ${ERROR_CODE}
    fi
  done
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
