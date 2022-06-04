package position

import (
	"errors"
	"math/big"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetPositions() ([]model.Position, error)
	GetPosition(uuid.UUID) (model.Position, error)
}

var positions = []model.Position{
	{
		Guid:         uuid.New(),
		PositionType: model.Long,
		Coin:         "BTC",
		Amount:       *big.NewInt(100),
	},
	{
		Guid:         uuid.New(),
		PositionType: model.Short,
		Coin:         "BTC",
		Amount:       *big.NewInt(100),
	},
}

type repository struct{}

func NewRepository() Repository {
	return repository{}
}

func (repository repository) GetPositions() ([]model.Position, error) {
	return positions, nil
}

func (repository repository) GetPosition(guid uuid.UUID) (model.Position, error) {
	for _, position := range positions {
		if position.Guid == guid {
			return position, nil
		}
	}

	return model.Position{}, errors.New("positionNotFound")
}
