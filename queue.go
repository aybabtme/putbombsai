package putbombsai

import (
	"container/list"
	"github.com/aybabtme/bomberman/player"
)

// Moves
type MoveQueue struct {
	l *list.List
}

func NewMoveQueue() *MoveQueue {
	return &MoveQueue{l: list.New()}
}

func (q *MoveQueue) Next() player.Move {
	next := q.l.Front()
	q.l.Remove(next)
	return next.Value.(player.Move)
}

func (q *MoveQueue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *MoveQueue) Enqueue(m player.Move) {
	q.l.PushBack(m)
}

// Points
type PointQueue struct {
	l *list.List
}

func NewPointQueue() *PointQueue {
	return &PointQueue{l: list.New()}
}

func (q *PointQueue) Next() Point {
	next := q.l.Front()
	q.l.Remove(next)
	return next.Value.(Point)
}

func (q *PointQueue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *PointQueue) Enqueue(p Point) {
	q.l.PushBack(p)
}

// Plans
type PlanQueue struct {
	l *list.List
}

func NewPlanQueue() *PlanQueue {
	return &PlanQueue{l: list.New()}
}

func (q *PlanQueue) Next() Plan {
	next := q.l.Front()
	q.l.Remove(next)
	return next.Value.(Plan)
}

func (q *PlanQueue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *PlanQueue) Enqueue(p Plan) {
	q.l.PushBack(p)
}
