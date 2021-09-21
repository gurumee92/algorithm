from collections import Counter
from dataclasses import dataclass
from itertools import product


@dataclass(frozen=True)
class Pos:
    x: int
    y: int

    def neighbors(self):
        return [
            Pos(self.x, self.y - 1),
            Pos(self.x + 1, self.y),
            Pos(self.x, self.y + 1),
            Pos(self.x - 1, self.y),
        ]


def make_tile_from_positions(positions):
    """Smallest possible representation with rotation"""

    def rotate90(tile):
        return tuple(
            tuple(tile[i][j] for i in range(len(tile)))
            for j in reversed(range(len(tile[0])))
        )

    positions = set(positions)

    xs = [pos.x for pos in positions]
    min_x = min(xs)
    max_x = max(xs)

    ys = [pos.y for pos in positions]
    min_y = min(ys)
    max_y = max(ys)

    tile_representations = [
        tuple(
            tuple(Pos(i, j) in positions for j in range(min_y, max_y + 1))
            for i in range(min_x, max_x + 1)
        )
    ]

    for __ in range(3):
        tile_representations.append(rotate90(tile_representations[-1]))

    return min(tile_representations)


def get_tile_size(tile):
    return sum(sum(row) for row in tile)


def parse_tiles(board, tile_value=1):
    n = len(board)

    # Add sentinel boundaries
    sentinel = 1 - tile_value

    board = [
        [sentinel] * (n + 2),
        *([sentinel] + row + [sentinel] for row in board),
        [sentinel] * (n + 2),
    ]

    # Detect tiles
    tile_positions = []
    for i, j in product(range(1, n + 1), range(1, n + 1)):
        if board[i][j] == tile_value:
            stack = [Pos(i, j)]
            squares = []
            while stack:
                curr = stack.pop()
                board[curr.x][curr.y] = sentinel
                squares.append(curr)
                for neighbor in curr.neighbors():
                    if board[neighbor.x][neighbor.y] == tile_value:
                        stack.append(neighbor)
            tile_positions.append(squares)

    # Make tiles
    tiles = [make_tile_from_positions(p) for p in tile_positions]

    return tiles


def solution(game_board, table):
    tiles = parse_tiles(table, 1)
    empty_spaces = parse_tiles(game_board, 0)

    tile_counter = Counter(tiles)
    empty_space_counter = Counter(empty_spaces)

    used_tiles = tile_counter & empty_space_counter

    return sum(get_tile_size(tile) * occ for tile, occ in used_tiles.items())