package todo

import (
	"github.com/ravielze/oculi/common/functions"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Undone(req request.ReqContext) error {
	data, err := req.Get(consts.ParameterPrefix("id"))
	if err != nil {
		return err
	}
	id := functions.Atoi(data.(string), 0)
	return h.domain.Todo.Undone(req, id)
}
