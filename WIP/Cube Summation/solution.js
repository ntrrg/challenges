'use strict';

let input = "";

process.stdin.resume();
process.stdin.setEncoding("ascii");

process.stdin.on("data", (_input) => {
  input += _input;
});

process.stdin.on("end", () => {
  processData(input);
});

function processData(input) {
  input = input.split('\n');

  let T = input.shift();

  while (T--) {
    let [, M] = input.shift().split(' ');
    const matrix = new Map();

    // Actions

    while (M--) {
      const [action, ...args] = input.shift().split(' ');

      switch(action) {
        case 'UPDATE':
          const coords = args.slice(0, -1).map((coord) => coord * 1);
          const value = args.slice(-1)[0] * 1;

          matrix.set(coords, value);
          break;

        case 'QUERY':
          const p1 = args.slice(0, 3).map((coord) => coord * 1);
          const p2 = args.slice(3).map((coord) => coord * 1);
          let result = 0;

          for (const [point, value] of matrix.entries()) {
            if (point[0] >= p1[0] && point[1] >= p1[1] && point[2] >= p1[2] &&
                point[0] <= p2[0] && point[1] <= p2[1] && point[2] <= p2[2]) {
              console.log(p1, p2, point);
              result += value;
            }
          }

          console.log(result);
          break;
      }
    }
  }
}
