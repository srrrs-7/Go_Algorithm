Solution6();

function Solution6() {
    const matrix = [
        [3, 4, 4],
        [4, 3, 4],
        [3, 2, 4],
        [2, 2, 2],
        [3, 3, 4],
        [1, 4, 4],
        [4, 1, 1],
    ];
    const cnt = solution(matrix);
    console.log(cnt);
}

function solution(A) {
    const N = A.length;
    const M = A[0].length;
    let sum = 0;

    if (N === 0 || M === 0) return 0;

    const B = A.map((inner) => inner.slice());
    for (let i = 0; i < N; i++) {
        for (let j = 0; j < M; j++) {
            if (B[i][j] >= 0) {
                checkNeighbor(A, B, i, j, N, M);
                sum += 1;
            }
        }
    }
    return sum;
}

const checkNeighbor = (A, B, i, j, N, M) => {
    if (B[i][j] === -1) return;
    B[i][j] = -1;
    if (i + 1 < N) {
        if (A[i + 1][j] === A[i][j]) {
            checkNeighbor(A, B, i + 1, j, N, M);
        }
    }
    if (i - 1 >= 0) {
        if (A[i - 1][j] === A[i][j]) {
            checkNeighbor(A, B, i - 1, j, N, M);
        }
    }
    if (j + 1 < M) {
        if (A[i][j + 1] === A[i][j]) {
            checkNeighbor(A, B, i, j + 1, N, M);
        }
    }
    if (j - 1 >= 0) {
        if (A[i][j - 1] === A[i][j]) {
            checkNeighbor(A, B, i, j - 1, N, M);
        }
    }
};