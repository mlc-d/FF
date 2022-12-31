package topic

import "gitlab.com/mlc-d/ff/dto"

type Service interface {
	Create(t *dto.Topic)
}

type service struct {
}
