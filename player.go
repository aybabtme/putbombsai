package putbombsai

import (
	"github.com/aybabtme/bomberman/player"
)

var _ player.Player = &BombsPlayer{}

type BombsPlayer struct {
	state   player.State
	update  chan player.State
	outMove chan player.Move
}

func NewBombsPlayer(state player.State, seed int64) player.Player {
	h := &BombsPlayer{
		state:   state,
		update:  make(chan player.State, 1),
		outMove: make(chan player.Move, 1),
	}

	go func() {

		plan := planPutBombThenHide(&state)

		for state := range h.update {
			_ = state

			if !plan.IsEmpty() {
				h.outMove <- plan.Next()
				continue
			}

			plan = planPutBombThenHide(&state)
		}
	}()

	return h
}

func (h *BombsPlayer) Name() string {
	return h.state.Name
}

func (h *BombsPlayer) Move() <-chan player.Move {
	return h.outMove
}

func (h *BombsPlayer) Update() chan<- player.State {
	return h.update
}

type Point struct {
	X, Y int
}

func planPutBombThenHide(state *player.State) *MoveQueue {
	plan := NewMoveQueue()
	bombs := planToPutBomb(plan, state)
	return planAwayFromBomb(bombs, plan, state)
}

type Plan struct {
	point Point
	path  []player.Move
}

func NewPlan(p Point) Plan {
	return Plan{
		point: p,
		path:  make([]player.Move, 0),
	}
}

func (p Plan) extend(newPoint Point, next player.Move) Plan {
	newPath := make([]player.Move, 0, len(p.path)+1)
	for _, m := range p.path {
		newPath = append(newPath, m)
	}
	newPath = append(newPath, next)
	return Plan{
		point: newPoint,
		path:  newPath,
	}
}
