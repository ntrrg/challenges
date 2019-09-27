process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function main() {
  const Q = parseInt(readLine());

  for (let a0 = 0; a0 < Q; a0++) {
    let [L, R] = readLine().split(' ').map((n) => parseInt(n));
    let result = 0;

    while (L <= R) {
      if (L <= R && L % 4 === 0) { result ^= L; ++L; }
      if (L <= R && L % 4 === 1) { result ^= 1; ++L; }
      if (L <= R && L % 4 === 2) { result ^= L + 1; ++L; }
      if (L <= R && L % 4 === 3) { ++L; }
    }

    console.log(result);
  }
}
