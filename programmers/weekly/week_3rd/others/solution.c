#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

#define _USE_MATH_DEFINES
#include <math.h>

int dfs_board[50][50] = { 0 }, dfs_table[50][50] = { 0 }, dot_game_board[101][101] = { 0 }, dot_table[101][101] = { 0 };
int board_order = 0, table_order = 0, indices = 0, count = 0;
char game_board_polyomino[6][625][50] = { 0 }, table_polyomino[6][625][50] = { 0 };
int polyomino[50] = { 0 };

int dot_dfs(int board[][101], int x, int y, int previous, int rows, int cols) {
    board[x][y] = 0;
    if (board[x][y + 1] && y + 1 < rows) {
        if (x - 1 >= 0) {
            if (board[x - 1][y]) {
                polyomino[indices++] = (int)sin((previous - 270.0) * M_PI / 180.0) + 2; // 1 = clockwise, 2 = straight, 3 = counterclockwise
                dot_dfs(board, x - 1, y, 270, rows, cols);
            }
            else {
                polyomino[indices++] = (int)sin((previous - 0.0) * M_PI / 180.0) + 2;
                dot_dfs(board, x, y + 1, 0, rows, cols);
            }
        }
        else {
            polyomino[indices++] = (int)sin((previous - 0.0) * M_PI / 180.0) + 2;
            dot_dfs(board, x, y + 1, 0, rows, cols);
        }
    }
    else if (board[x + 1][y] && x + 1 < cols) {
        polyomino[indices++] = (int)sin((previous - 90.0) * M_PI / 180.0) + 2;
        dot_dfs(board, x + 1, y, 90, rows, cols);
    }
    else if (board[x][y - 1] && y - 1 >= 0) {
        polyomino[indices++] = (int)sin((previous - 180.0) * M_PI / 180.0) + 2;
        dot_dfs(board, x, y - 1, 180, rows, cols);
    }
    else if (board[x - 1][y] && x - 1 >= 0) {
        polyomino[indices++] = (int)sin((previous - 270.0) * M_PI / 180.0) + 2;
        dot_dfs(board, x - 1, y, 270, rows, cols);
    }
    else {
        polyomino[indices++] = (int)sin((previous - 270.0) * M_PI / 180.0) + 2;
    }
    return count;
}

int dfs(int board[][50], int x, int y, int rows, int cols) {
    board[x][y] = 0;
    if (board[x][y + 1] && y + 1 < rows) {
        count++;
        dfs(board, x, y + 1, rows, cols);
    }
    if (board[x + 1][y] && x + 1 < cols) {
        count++;
        dfs(board, x + 1, y, rows, cols);
    }
    if (board[x][y - 1] && y - 1 >= 0) {
        count++;
        dfs(board, x, y - 1, rows, cols);
    }
    if (board[x - 1][y] && x - 1 >= 0) {
        count++;
        dfs(board, x - 1, y, rows, cols);
    }
    return count;
}

void LMSR(int S[]) {
    int temp[100] = { 0 };
    int len_S = indices;
    int max = len_S * 2;
    for (int y = 0; y < len_S; y++) {
        temp[y] = S[y];
    }
    for (int y = len_S; y < max; y++) { // Concatenate string to it self to avoid modular arithmetic
        temp[y] = S[y - len_S];
    }
    int f[50] = { 0 };
    for (int y = 0; y < max; y++) {  // Failure function
        f[y] = -1;
    }
    int i = 0;
    int sj = 0, k = 0; // Least rotation of string found so far
    for (int j = 1; j < max; j++) {
        sj = temp[j];
        i = f[j - k - 1];
        while (i != -1 && sj != S[k + i + 1]) {
            if (sj < temp[k + i + 1]) {
                k = j - i - 1;
            }
            i = f[i];
        }
        if (sj != temp[k + i + 1]) {  // if sj != S[k+i+1], then i == -1
            if (sj < temp[k]) {
                k = j;
            }
            f[j - k] = -1;
        }
        else {
            f[j - k] = i + 1;
        }
    }

    // k represents rotation

    int S_left[50] = { 0 };
    int S_right[50] = { 0 };
    for (int y = 0; y < k; y++) {
        S_left[y] = temp[y];
    }
    for (int y = k; y < len_S; y++) {
        S_right[y - k] = temp[y];
    }
    for (int y = 0; y < len_S - k; y++) {
        S[y] = S_right[y];
    }
    for (int y = len_S - k; y < len_S; y++) {
        S[y] = S_left[y + k - len_S];
    }
}

int solution(int** game_board, size_t game_board_rows, size_t game_board_cols, int** table, size_t table_rows, size_t table_cols) {
    for (int i = 0; i < game_board_rows; i++) {
        for (int j = 0; j < game_board_cols; j++) {
            if (!game_board[i][j]) {
                dot_game_board[2 * i][2 * j] += 1;
                dot_game_board[2 * i + 1][2 * j] += 1;
                dot_game_board[2 * i][2 * j + 1] += 1;
                dot_game_board[2 * i + 1][2 * j + 1] += 1;
                dfs_board[i][j] = game_board[i][j] = 1;
            }
            else if(game_board[i][j]){
                dfs_board[i][j] = game_board[i][j] = 0;
            }

            if (table[i][j]) {
                dot_table[2 * i][2 * j] += 1;
                dot_table[2 * i + 1][2 * j] += 1;
                dot_table[2 * i][2 * j + 1] += 1;
                dot_table[2 * i + 1][2 * j + 1] += 1;
                dfs_table[i][j] = 1;
            }
            else {
                dfs_table[i][j] = 0;
            }
        }
    }

    // 10 -> 1100
    // 01    1100
    //       0011
    //       0011

    int dot_game_board_rows = 2 * game_board_rows + 1;
    int dot_game_board_cols = 2 * game_board_cols + 1;

    int delete_board[300][2] = { 0 }, delete_board_index = 0;
    int delete_table[300][2] = { 0 }, delete_table_index = 0;
    for (int i = 0; i < dot_game_board_rows - 4; i++) {
        for (int j = 0; j < dot_game_board_cols - 4; j++) {
            int sum_board = 0;
            int sum_table = 0;
            int x = i;
            int y = j;
            for (x; x < i + 4; x++) {
                y = j;
                for (y; y < j + 4; y++) {
                    sum_board += dot_game_board[x][y];
                    sum_table += dot_table[x][y];
                }
            }
            if (sum_table == 16) {
                delete_table[delete_table_index][0] = i + 1;
                delete_table[delete_table_index++][1] = j + 1;
                delete_table[delete_table_index][0] = i + 1;
                delete_table[delete_table_index++][1] = j + 2;
                delete_table[delete_table_index][0] = i + 2;
                delete_table[delete_table_index++][1] = j + 1;
                delete_table[delete_table_index][0] = i + 2;
                delete_table[delete_table_index++][1] = j + 2;
            }
            if (sum_board == 16) {
                delete_board[delete_board_index][0] = i + 1;
                delete_board[delete_board_index++][1] = j + 1;
                delete_board[delete_board_index][0] = i + 1;
                delete_board[delete_board_index++][1] = j + 2;
                delete_board[delete_board_index][0] = i + 2;
                delete_board[delete_board_index++][1] = j + 1;
                delete_board[delete_board_index][0] = i + 2;
                delete_board[delete_board_index++][1] = j + 2;
            }
        }
    }
    for (int i = 0; i < delete_board_index; i++) {
        dot_game_board[delete_board[i][0]][delete_board[i][1]] = 0;
    }
    for (int i = 0; i < delete_table_index; i++) {
        dot_table[delete_table[i][0]][delete_table[i][1]] = 0;
    }

     for (int i = 0; i < dot_game_board_rows; i++) {
        for (int j = 0; j < dot_game_board_cols; j++) {
            if (dot_game_board[i][j]) {
                count = 1;
                indices = 0;
                dot_dfs(dot_game_board, i, j, 270, dot_game_board_rows, dot_game_board_cols);
                dfs(dfs_board, i / 2, j / 2, game_board_rows, game_board_cols);
                LMSR(polyomino);
                for (int z = 0; z < indices; z++) {
                    game_board_polyomino[count - 1][board_order][z] = polyomino[z];
                }
                board_order++;
            }
            if (dot_table[i][j]) {
                count = 1;
                indices = 0;
                dot_dfs(dot_table, i, j, 270, dot_game_board_rows, dot_game_board_cols);
                dfs(dfs_table, i / 2, j / 2, table_rows, table_cols);
                LMSR(polyomino);
                for (int z = 0; z < indices; z++) {
                    table_polyomino[count - 1][table_order][z] = polyomino[z];
                }
                table_order++;
            }
        }
    }

    int answer = 0;
    for (int i = 0; i < 6; i++) {
        for (int j = 0; j < table_order; j++) {
            for (int k = 0; k < board_order; k++) {
                if (game_board_polyomino[i][k][0] > 0 && table_polyomino[i][j][0] > 0) {
                    if (!strcmp(game_board_polyomino[i][k], table_polyomino[i][j])) {
                        game_board_polyomino[i][k][0] = 0;
                        table_polyomino[i][j][0] = 0;
                        answer += (i + 1);
                    }
                }
            }
        }
    }
    return answer;
}
