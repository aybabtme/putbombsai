package putbombsai

import (
	"github.com/aybabtme/bomberman/objects"
	"github.com/aybabtme/bomberman/player"
)

func planToPutBomb(plan *MoveQueue, state *player.State) []Point {
	bombs := make([]Point, 0)

	destination := findWhereToPutBomb(state)
	bombs = append(bombs, destination)

	for _, move := range findPathTo(destination, state) {
		plan.Enqueue(move)
	}

	plan.Enqueue(player.PutBomb)

	return bombs
}

// find destination to put bomb
func findWhereToPutBomb(state *player.State) Point {
	closed := make(map[Point]struct{})
	fringe := NewPointQueue()

	highestSoFar := Point{X: state.X, Y: state.Y}
	bestValueSofar := 0

	fringe.Enqueue(highestSoFar)

	for !fringe.IsEmpty() {
		next := fringe.Next()

		for _, succ := range successor(next, state) {
			if _, ok := closed[succ]; ok {
				continue
			}

			value := getPutBombValue(succ, state)
			// only if actually bigger, so we find closest most valuable
			if value > bestValueSofar {
				highestSoFar = succ
			}

			closed[succ] = struct{}{}
		}
	}

	return highestSoFar
}

// find how to get there
func findPathTo(p Point, state *player.State) []player.Move {
	closed := make(map[Point]struct{})

	fringe := NewPlanQueue()
	now := NewPlan(Point{X: state.X, Y: state.Y})
	fringe.Enqueue(now)

	for !fringe.IsEmpty() {
		next := fringe.Next()

		if p.X == next.point.X && p.Y == next.point.Y {
			return next.path
		}

		for _, succ := range successorPlan(next, state) {
			if _, ok := closed[succ.point]; ok {
				continue
			}

			fringe.Enqueue(succ)
			closed[succ.point] = struct{}{}
		}
	}

	return []player.Move{}
}

// helpers
func successor(p Point, state *player.State) (succ []Point) {
	for _, next := range cellsAround(p, state) {
		name := state.Board[next.X][next.Y].Name
		if name[0] == 'p' {
			// can walk over players
			succ = append(succ, next)
		} else {
			switch name {
			case objects.Ground.String():
				succ = append(succ, next)
			}
		}

	}
	return
}

func successorPlan(p Plan, state *player.State) (succ []Plan) {
	x, y := p.point.X, p.point.Y
	for _, next := range []Plan{
		NewPlan(p.point).extend(Point{X: x + 1, Y: y}, player.Right),
		NewPlan(p.point).extend(Point{X: x - 1, Y: y}, player.Left),
		NewPlan(p.point).extend(Point{X: x, Y: y + 1}, player.Up),
		NewPlan(p.point).extend(Point{X: x, Y: y - 1}, player.Down),
	} {

		if next.point.X <= 0 || next.point.X >= len(state.Board) {
			continue
		}

		if next.point.Y <= 0 || next.point.Y >= len(state.Board[0]) {
			continue
		}
		succ = append(succ, next)
	}
	return
}

func getPutBombValue(p Point, state *player.State) int {
	rocksAround := 0
	for _, next := range cellsAround(p, state) {
		switch state.Board[next.X][next.Y].Name {
		case objects.Rock.String():
			rocksAround++
		}
	}
	return rocksAround
}

func cellsAround(p Point, state *player.State) (around []Point) {
	x, y := p.X, p.Y
	for _, next := range []Point{
		{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1},
	} {

		if next.X <= 0 || next.X >= len(state.Board) {
			continue
		}

		if next.Y <= 0 || next.Y >= len(state.Board[0]) {
			continue
		}
		around = append(around, next)
	}
	return
}
