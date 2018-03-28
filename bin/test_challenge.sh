#!/bin/sh

main() {
  echo "$(basename "${PWD}")"

  TEST_CASE=$1

  if [ -n "${TEST_CASE}" ]; then
    run
  else
    TEST_CASES="$(find "input" -name "input??.txt")"

    for TEST_CASE in ${TEST_CASES}; do
      TEST_CASE="$(echo "${TEST_CASE}" | grep -oe "[0-9][0-9]")"

      echo "  Test case ${TEST_CASE}:"
      run
    done
  fi
}

run() {
  INPUT="input/input${TEST_CASE}.txt"
  OUTPUT="output/output${TEST_CASE}.txt"

  if [ ! -f "${INPUT}" ]; then
    echo "  The test case ${TEST_CASE} doesn't exists"
    echo
    exit 1
  fi

  cat "${INPUT}" |
    go run solution.go |
    diff -yw --suppress-common-lines --color=always - "${OUTPUT}" &&
      echo "    Pass" ||
      echo "    Fail"

  echo
}

main $@
