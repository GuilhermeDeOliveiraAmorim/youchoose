package usecase

import (
	"net/http"
	"youchoose/internal/entity"
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
	valueobject "youchoose/internal/value_object"
)

type CreateChooserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Day      int    `json:"day"`
	Month    int    `json:"month"`
	Year     int    `json:"year"`
	ImageID  string `json:"image_id"`
}

type CreateChooserOutputDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
	ImageID string `json:"image_id"`
}

type CreateChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewCreateChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *CreateChooserUseCase {
	return &CreateChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (cc *CreateChooserUseCase) Execute(input CreateChooserInputDTO) (CreateChooserOutputDTO, []util.ProblemDetails) {
	problemsDetails := []util.ProblemDetails{}

	_, error := cc.ChooserRepository.GetByEmail(input.Email)
	if error == nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "ValidationError",
			Title:    "E-mail já está em uso",
			Status:   http.StatusConflict,
			Detail:   "O e-mail fornecido já está sendo utilizado por outro chooser.",
			Instance: util.RFC409,
		})

		return CreateChooserOutputDTO{}, problemsDetails
	}

	newLogin, newLoginProblems := valueobject.NewLogin(input.Email, input.Password)
	if len(newLoginProblems) > 0 {
		problemsDetails = append(problemsDetails, newLoginProblems...)
	}

	newAddress, newAddressProblems := valueobject.NewAddress(input.City, input.State, input.Country)
	if len(newAddressProblems) > 0 {
		problemsDetails = append(problemsDetails, newAddressProblems...)
	}

	newBirthdate, newBirthdateProblems := valueobject.NewBirthDate(input.Day, input.Month, input.Year)
	if len(newBirthdateProblems) > 0 {
		problemsDetails = append(problemsDetails, newBirthdateProblems...)
	}

	newChooser, newChooserProblems := entity.NewChooser(input.Name, newLogin, newAddress, newBirthdate, input.ImageID)
	if len(newChooserProblems) > 0 {
		problemsDetails = append(problemsDetails, newChooserProblems...)
	}

	if len(problemsDetails) > 0 {
		return CreateChooserOutputDTO{}, problemsDetails
	}

	error = cc.ChooserRepository.Create(newChooser)
	if error != nil {
		problemsDetails = append(problemsDetails, util.ProblemDetails{
			Type:     "Internal Server Error",
			Title:    "Erro ao persistir um chooser",
			Status:   http.StatusInternalServerError,
			Detail:   error.Error(),
			Instance: util.RFC500,
		})
	}

	output := CreateChooserOutputDTO{
		ID:      newChooser.ID,
		Name:    newChooser.Name,
		Email:   newChooser.Login.Email,
		City:    newChooser.Address.City,
		State:   newChooser.Address.State,
		Country: newChooser.Address.Country,
		Day:     newChooser.BirthDate.Day,
		Month:   newChooser.BirthDate.Month,
		Year:    newChooser.BirthDate.Year,
		ImageID: newChooser.ImageID,
	}

	return output, problemsDetails
}
