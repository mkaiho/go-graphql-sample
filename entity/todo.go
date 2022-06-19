package entity

import (
	"errors"
	"fmt"
)

type (
	TodoID string
)

func ParseTodoID(value string) (TodoID, error) {
	id := TodoID(value)
	if err := id.Validate(); err != nil {
		return id, fmt.Errorf("invalid todo ID: %w", err)
	}
	return id, nil
}

func (i TodoID) Validate() error {
	if i.isEmpty() {
		return errors.New("ID is empty")
	}
	return nil
}

func (i TodoID) String() string {
	return string(i)
}

func (i TodoID) isEmpty() bool {
	return len(i) == 0
}

type Todo struct {
	id   TodoID
	text string
	done bool
}

func NewTodo(
	id TodoID,
	text string,
	done bool,
) (*Todo, error) {
	todo := Todo{
		id:   id,
		text: text,
		done: done,
	}
	if err := todo.Validate(); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Todo) Validate() error {
	if err := t.ID().Validate(); err != nil {
		return fmt.Errorf("invalid todo")
	}
	if len(t.Text()) == 0 {
		return fmt.Errorf("invalid text")
	}
	return nil
}

func (t *Todo) ID() TodoID {
	if t == nil {
		return ""
	}
	return t.id
}

func (t *Todo) Text() string {
	if t == nil {
		return ""
	}
	return t.text
}

func (t *Todo) Done() bool {
	if t == nil {
		return false
	}
	return t.done
}

type Todos []*Todo
