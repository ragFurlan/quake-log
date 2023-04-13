package quakeLog

import (
	"errors"
	"quakelog/internal/entity"
	usecaseMock "quakelog/test/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var anyError = errors.New("Error")

func TestGroupByGame_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	usecase := usecaseMock.NewMockQuakeLogUsecase(controller)
	service := NewQuakeLog(usecase)

	quakeLog := entity.QuakeLog{
		TotalKills: 30,
		Players:    []string{"ana", "john"},
		Kills: map[string]int{
			"ana":  10,
			"john": 20,
		},
	}

	quakeLogs := map[string]entity.QuakeLog{
		"game_01": quakeLog,
	}

	usecase.EXPECT().GroupByGame().Return(quakeLogs, nil)
	_, err := service.GroupByGame()

	assert.NoError(t, err)

}

func TestGroupByGame_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	usecase := usecaseMock.NewMockQuakeLogUsecase(controller)
	service := NewQuakeLog(usecase)
	quakeLogs := map[string]entity.QuakeLog{}

	usecase.EXPECT().GroupByGame().Return(quakeLogs, anyError)
	_, err := service.GroupByGame()

	assert.Error(t, err)

}

func TestGroupedByTypeOfDeath_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	usecase := usecaseMock.NewMockQuakeLogUsecase(controller)
	service := NewQuakeLog(usecase)

	quakeLogKills := entity.QuakeLogKills{
		KillsByMeans: map[string]int{
			"MOD_SHOTGUN": 10,
		},
	}

	quakeLogs := map[string]entity.QuakeLogKills{
		"game_01": quakeLogKills,
	}

	usecase.EXPECT().GroupedByTypeOfDeath().Return(quakeLogs, nil)
	_, err := service.GroupedByTypeOfDeath()

	assert.NoError(t, err)

}

func TestGroupedByTypeOfDeath_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	usecase := usecaseMock.NewMockQuakeLogUsecase(controller)
	service := NewQuakeLog(usecase)

	quakeLogs := map[string]entity.QuakeLogKills{}

	usecase.EXPECT().GroupedByTypeOfDeath().Return(quakeLogs, anyError)

	_, err := service.GroupedByTypeOfDeath()

	assert.Error(t, err)
}
