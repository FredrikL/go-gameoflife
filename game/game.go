package game

type Pos struct {
	X, Y int
}

func unique[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// returns all 8 surrounding positions for given p
func getSurround(p Pos) []Pos {
	return []Pos{{p.X - 1, p.Y - 1}, {p.X, p.Y - 1}, {p.X + 1, p.Y - 1},
		{p.X - 1, p.Y}, {p.X + 1, p.Y},
		{p.X - 1, p.Y + 1}, {p.X, p.Y + 1}, {p.X + 1, p.Y + 1}}
}

// returns all alive cells and their surrounding positions
func toEvaluate(board map[Pos]bool) []Pos {
	r := make([]Pos, 0)
	for k := range board {
		r = append(r, k)
		r = append(r, getSurround(k)...)
	}
	return unique(r)
}

func shouldBeAlive(currentlyAlive bool, numberOfNeighbours int) bool {
	if currentlyAlive {
		return numberOfNeighbours == 2 || numberOfNeighbours == 3
	}
	return numberOfNeighbours == 3
}

// Advance the board, returns the next generation of `board`
func Advance(board map[Pos]bool) map[Pos]bool {
	newBoard := make(map[Pos]bool)
	for _, p := range toEvaluate(board) {
		n := countNeighbours(p, board)
		if shouldBeAlive(board[p], n) {
			newBoard[p] = true
		}
	}
	return newBoard
}

// returns the number of alive neighbours for a given position
func countNeighbours(p Pos, board map[Pos]bool) int {
	sum := 0
	for _, v := range getSurround(p) {
		if board[v] {
			sum += 1
		}
	}
	return sum
}
