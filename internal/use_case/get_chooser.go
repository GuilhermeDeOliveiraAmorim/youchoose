package usecase

import (
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetChooserInputDTO struct {
	ChooserID       string `json:"chooser_id"`
	ChooserIDToFind string `json:"chooser_id_to_find"`
}

type GetChooserUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
}

func NewGetChooserUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
) *GetChooserUseCase {
	return &GetChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (gc *GetChooserUseCase) Execute(input GetChooserInputDTO) (ChooserOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gc.ChooserRepository, input.ChooserID, "GetChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ChooserOutputDTO{}, chooserValidatorProblems
	}

	chooser, chooserValidatorProblems := chooserValidator(gc.ChooserRepository, input.ChooserIDToFind, "GetChooserUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ChooserOutputDTO{}, chooserValidatorProblems
	}

	output := NewChooserOutputDTO(chooser)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: []util.ProblemDetails{},
	}
}
