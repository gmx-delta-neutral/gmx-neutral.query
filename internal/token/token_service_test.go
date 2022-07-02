package token

import (
	"math/big"
	"testing"

	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

type mockPositionRepository struct{}

var mockPosition = model.Position{
	Guid:         uuid.New(),
	PositionType: model.Long,
	Coin:         "BTC",
	Amount:       *big.NewInt(100),
}

var mockPositions = []model.Position{
	mockPosition,
}

func (m *mockPositionRepository) GetPositions() ([]model.Position, error) {
	return mockPositions, nil
}

func (m *mockPositionRepository) GetPosition(uuid.UUID) (model.Position, error) {
	return mockPosition, nil
}

func setupService() Service {
	service := NewService(&mockPositionRepository{})
	return service
}

func TestGetPositions(t *testing.T) {
	service := setupService()

	positions, err := service.GetPositions()

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(positions, mockPositions, cmp.Comparer(func(x model.Position, y model.Position) bool {
		return x.		
	})) {
		t.Error(cmp.Diff(positions, mockPositions))
	}
}

func compare_position(a model.Position, b model.Position) bool {
	return a.Coin == b.Coin && a.Amount.Cmp(b.Amount) && a.Guid == b.Guid && a.
}
