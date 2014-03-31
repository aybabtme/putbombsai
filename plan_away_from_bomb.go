package putbombsai

import (
	"github.com/aybabtme/bomberman/player"
)

func planAwayFromBomb(bombs []Point, plan *MoveQueue, state *player.State) *MoveQueue {

	destination := findWhereToRunAway(bombs, state)
	for _, move := range findPathTo(destination, state) {
		plan.Enqueue(move)
	}

	return plan
}

func findWhereToRunAway(bombs []Point, state *player.State) Point {
	closed := make(map[Point]struct{})
	fringe := NewPointQueue()

	here := Point{X: state.X, Y: state.Y}

	fringe.Enqueue(here)

	for !fringe.IsEmpty() {
		next := fringe.Next()

		for _, succ := range successor(next, state) {
			if _, ok := closed[succ]; ok {
				continue
			}

			linedWithBomb := false
			for _, bomb := range bombs {
				if bomb.X == succ.X || bomb.Y == succ.Y {
					linedWithBomb = true
				}
			}

			if !linedWithBomb {
				return succ
			}

			closed[succ] = struct{}{}
		}
	}

	return here
}
