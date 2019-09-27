'use strict';

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

function bigSorting(arr) {
  arr.sort((a, b) => {
    if (a.length === b.length) {
      if (a === b) {
        return 0;
      } else if (a < b) {
        return -1;
      } else {
        return 1;
      }
    } else {
      return a.length - b.length;
    }
  });

  return arr;
}

function main() {
  var n = parseInt(readLine());
  var arr = [];
  for(var arr_i = 0; arr_i < n; arr_i++){
     arr[arr_i] = readLine();
  }
  var result = bigSorting(arr);
  console.log(result.join("\n"));
}
