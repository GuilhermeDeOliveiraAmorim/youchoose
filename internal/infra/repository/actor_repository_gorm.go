package gorm

import (
	"youchoose/internal/entity"
	valueobject "youchoose/internal/value_object"

	"gorm.io/gorm"
)

type ActorRepository struct {
	gorm *gorm.DB
}

// Create implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Create(actor *entity.Actor) error {
	panic("unimplemented")
}

// CreateMany implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) CreateMany(actors *[]entity.Actor) error {
	panic("unimplemented")
}

// Deactivate implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Deactivate(actor *entity.Actor) error {
	panic("unimplemented")
}

// DoTheseActorsAreIncludedInTheMovie implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) DoTheseActorsAreIncludedInTheMovie(movieID string, actorsIDs []string) (bool, []entity.Actor, error) {
	panic("unimplemented")
}

// DoTheseActorsExist implements repositoryinterface.ActorRepositoryInterface.
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

// GetAll implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) GetAll() ([]entity.Actor, error) {
	panic("unimplemented")
}

// GetByID implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) GetByID(actorID string) (entity.Actor, error) {
	panic("unimplemented")
}

// Update implements repositoryinterface.ActorRepositoryInterface.
func (a *ActorRepository) Update(actor *entity.Actor) error {
	panic("unimplemented")
}

func NewActorRepository(gorm *gorm.DB) *ActorRepository {
	return &ActorRepository{
		gorm: gorm,
	}
}
