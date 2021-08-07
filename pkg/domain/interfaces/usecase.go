package interfaces

import "github.com/m-mizutani/zenv/pkg/domain/model"

type Usecase interface {
	Exec(input *model.ExecInput) error
	List(input *model.ListInput) error
	Write(input *model.WriteInput) error

	SetConfig(config *model.Config)
}
