#!/bin/sh

main() {
  echo "${PWD}"

  TEST_CASE=$1

  if [ -n "${TEST_CASE}" ]; then
    run

    ERROR_CODE=$?

    if [ ! ${ERROR_CODE} -eq 0 ]; then
      return ${ERROR_CODE}
    fi
  else
    TEST_CASES="$(find "input" -name "input??.txt")"

    for TEST_CASE in ${TEST_CASES}; do
      TEST_CASE="$(echo "${TEST_CASE}" | grep -oe "[0-9][0-9]")"

      echo "  Test case ${TEST_CASE}:"
      run

      ERROR_CODE=$?

      if [ ! ${ERROR_CODE} -eq 0 ]; then
        return ${ERROR_CODE}
      fi
    done
  fi
}

run() {
  INPUT="input/input${TEST_CASE}.txt"
  OUTPUT="output/output${TEST_CASE}.txt"

  if [ ! -f "${INPUT}" ]; then
    echo "  The test case ${TEST_CASE} doesn't exists"
    echo
    return 1
  fi

  go fmt solution.go
  go vet solution.go || return $?

  echo "$(cat "${OUTPUT}")" > /tmp/EXPECT
  echo "$(cat "${INPUT}" | go run solution.go)" > /tmp/RESULT
  diff /tmp/EXPECT /tmp/RESULT

  ERROR_CODE=$?

  if [ ${ERROR_CODE} -eq 0 ]; then
    echo "    Pass"
    echo
  else
    echo "    Fail"
    echo
    return ${ERROR_CODE}
  fi
}

main $@
