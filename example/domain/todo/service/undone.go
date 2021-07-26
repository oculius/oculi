package service

import (
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/request"
)

func (s *service) Undone(req request.Context, todoId uint64) error {
	t, err := s.repository.GetByID(req, todoId)
	if err != nil {
		return err
	}

	if !t.IsDone {
		return constants.ErrTodoAlreadyUndoneState
	}

	return s.repository.Update(req, todoId, map[string]interface{}{
		"is_done": false,
	})
}
