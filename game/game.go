package game

type Pos struct {
	X, Y int
}

func getSurround(p Pos) []Pos {
	return []Pos{{p.X - 1, p.Y - 1}, {p.X, p.Y - 1}, {p.X + 1, p.Y - 1},
		{p.X - 1, p.Y}, {p.X + 1, p.Y},
		{p.X - 1, p.Y + 1}, {p.X, p.Y + 1}, {p.X + 1, p.Y + 1}}
}

func toEvaluate(board map[Pos]bool) []Pos {
	r := make([]Pos, 0)
	for k := range board {
		r = append(r, k)
		r = append(r, getSurround(k)...)
	}
	return r
}

func Advance(board map[Pos]bool) map[Pos]bool {
	newBoard := make(map[Pos]bool)
	positions := toEvaluate(board)
	for _, p := range positions {
		n := countNeighbours(p, board)
		alive := board[p]
		if alive && (n == 2 || n == 3) {
			newBoard[p] = true
		}
		if !alive && n == 3 {
			newBoard[p] = true
		}
	}
	return newBoard
}

func countNeighbours(p Pos, board map[Pos]bool) int {
	surround := getSurround(p)
	sum := 0
	for _, v := range surround {
		if board[v] {
			sum += 1
		}
	}
	return sum
}
