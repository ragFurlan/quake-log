package quakeLog

import (
	entity "quakelog/internal/entity"
	usecase "quakelog/internal/usecase/quakeLog"

)

type QuakeLog interface {
	GroupedMatch() (map[string]entity.QuakeLog, error)
	GroupedByTypeOfDeath() (map[string]entity.QuakeLogKills, error)
}

type quakeLog struct {
	quakeLogUseCase usecase.QuakeLogUsecase
}

func NewQuakeLog(quakeLogUseCase usecase.QuakeLogUsecase) *quakeLog {
	return &quakeLog{quakeLogUseCase: quakeLogUseCase}
}

func (ref quakeLog) GroupedMatch() (map[string]entity.QuakeLog, error) {
	quake, err := ref.quakeLogUseCase.GroupedMatch()
	if err != nil {
		return map[string]entity.QuakeLog{}, err
	}

	return quake, nil
}

func (ref quakeLog) GroupedByTypeOfDeath() (map[string]entity.QuakeLogKills, error) {
	quake, err := ref.quakeLogUseCase.GroupedByTypeOfDeath()
	if err != nil {
		return map[string]entity.QuakeLogKills{}, err
	}

	return quake, nil
}
