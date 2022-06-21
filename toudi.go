package toudi

import (
	"errors"
	"time"
)

type item struct {

    Task string
    Done bool
    CreatedAt time.Time
    CompletedAt time.Time
}

type Toudis []item

func (t *Toudis) Add(task string) {
    todo := item {
        Task: task,
        Done: false,
        CreatedAt: time.Now(),
        CompletedAt: time.Time{},
    }

    *t = append(*t, todo)
}

func (t *Toudis) Complete(index int) error {
    ls := *t
    if index <= 0 || index > len(ls) {
        return errors.New("invalid index")
    }

    ls[index - 1].CompletedAt = time.Now()
    ls[index - 1].Done = true

    return nil
}
