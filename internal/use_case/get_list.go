package usecase

import (
	repositoryinterface "youchoose/internal/repository_interface"
	"youchoose/internal/util"
)

type GetListInputDTO struct {
	ChooserID string `json:"chooser_id"`
	ListID    string `json:"id"`
}

type GetListUseCase struct {
	ChooserRepository repositoryinterface.ChooserRepositoryInterface
	ListRepository    repositoryinterface.ListRepositoryInterface
}

func NewGetListUseCase(
	ChooserRepository repositoryinterface.ChooserRepositoryInterface,
	ListRepository repositoryinterface.ListRepositoryInterface,
) *GetListUseCase {
	return &GetListUseCase{
		ChooserRepository: ChooserRepository,
		ListRepository:    ListRepository,
	}
}

func (gl *GetListUseCase) Execute(input GetListInputDTO) (ListOutputDTO, util.ProblemDetailsOutputDTO) {
	_, chooserValidatorProblems := chooserValidator(gl.ChooserRepository, input.ChooserID, "GetListUseCase")
	if len(chooserValidatorProblems.ProblemDetails) > 0 {
		return ListOutputDTO{}, chooserValidatorProblems
	}

	list, listValidatorProblems := listValidator(gl.ListRepository, input.ListID, "GetListUseCase")
	if len(listValidatorProblems.ProblemDetails) > 0 {
		return ListOutputDTO{}, listValidatorProblems
	}

	problemsDetails := []util.ProblemDetails{}

	output := NewListOutputDTO(list)

	return output, util.ProblemDetailsOutputDTO{
		ProblemDetails: problemsDetails,
	}
}
