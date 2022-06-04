package position

import (
	"errors"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	GetPositions() ([]model.Position, error)
	GetPosition(uuid.UUID) (model.Position, error)
}

type service struct {
	positionRepository Repository
}

func NewService(positionRepository Repository) *service {
	return &service{positionRepository: positionRepository}
}

func (service *service) GetPositions() ([]model.Position, error) {
	positions, err := service.positionRepository.GetPositions()

	if err != nil {
		return []model.Position{}, err
	}

	return positions, nil
}

func (service *service) GetPosition(guid uuid.UUID) (model.Position, error) {
	for _, position := range positions {
		if position.Guid == guid {
			return position, nil
		}
	}

	return model.Position{}, errors.New("positionNotFound")
}
