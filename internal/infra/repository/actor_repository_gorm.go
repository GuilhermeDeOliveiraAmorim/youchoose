package gorm

import (
	"errors"
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"

	"gorm.io/gorm"
)

type ActorRepository struct {
	gorm *gorm.DB
}

func NewActorRepository(gorm *gorm.DB) *ActorRepository {
	return &ActorRepository{
		gorm: gorm,
	}
}

func (a *ActorRepository) Create(actor *entity.Actor) error {
	if err := a.gorm.Create(&Actors{
		ID:            actor.ID,
		Active:        actor.Active,
		CreatedAt:     actor.CreatedAt,
		UpdatedAt:     actor.UpdatedAt,
		DeactivatedAt: actor.DeactivatedAt,
		Name:          actor.Name,
		Day:           actor.BirthDate.Day,
		Month:         actor.BirthDate.Month,
		Year:          actor.BirthDate.Year,
		CountryName:   actor.Nationality.CountryName,
		Flag:          actor.Nationality.Flag,
		ImageID:       actor.ImageID,
	}).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (a *ActorRepository) CreateMany(actors *[]entity.Actor) error {
	var actorsModel []Actors

	for _, actor := range *actors {
		actorsModel = append(actorsModel, Actors{
			ID:            actor.ID,
			Active:        actor.Active,
			CreatedAt:     actor.CreatedAt,
			UpdatedAt:     actor.UpdatedAt,
			DeactivatedAt: actor.DeactivatedAt,
			Name:          actor.Name,
			Day:           actor.BirthDate.Day,
			Month:         actor.BirthDate.Month,
			Year:          actor.BirthDate.Year,
			CountryName:   actor.Nationality.CountryName,
			Flag:          actor.Nationality.Flag,
			ImageID:       actor.ImageID,
		})
	}

	if err := a.gorm.Create(actorsModel).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (a *ActorRepository) Deactivate(actor *entity.Actor) error {
	panic("unimplemented")
}

func (a *ActorRepository) DoTheseActorsAreIncludedInTheMovie(movieID string, actorsIDs []string) (bool, []entity.Actor, error) {
	panic("unimplemented")
}

func (a *ActorRepository) DoTheseActorsExist(actorIDs []string) (bool, []entity.Actor, error) {
	var actorsModel []Actors
	result := a.gorm.Where("id IN ?", actorIDs).Find(&actorsModel)

	if result.Error != nil {
		return false, nil, result.Error
	}

	var actors []entity.Actor

	if result.RowsAffected != int64(len(actorIDs)) {
		return false, actors, nil
	}

	for _, actorModel := range actorsModel {
		actors = append(actors, entity.Actor{
			SharedEntity: entity.SharedEntity{
				ID:            actorModel.ID,
				Active:        actorModel.Active,
				CreatedAt:     actorModel.CreatedAt,
				UpdatedAt:     actorModel.UpdatedAt,
				DeactivatedAt: actorModel.DeactivatedAt,
			},
			Name:    actorModel.Name,
			ImageID: actorModel.ImageID,
			BirthDate: &valueobject.BirthDate{
				Day:   actorModel.Day,
				Month: actorModel.Month,
				Year:  actorModel.Year,
			},
			Nationality: &valueobject.Nationality{
				CountryName: actorModel.CountryName,
				Flag:        actorModel.Flag,
			},
		})
	}

	return true, actors, nil
}

func (a *ActorRepository) GetAll() ([]entity.Actor, error) {
	panic("unimplemented")
}

func (a *ActorRepository) GetByID(actorID string) (entity.Actor, error) {
	panic("unimplemented")
}

func (a *ActorRepository) Update(actor *entity.Actor) error {
	panic("unimplemented")
}
