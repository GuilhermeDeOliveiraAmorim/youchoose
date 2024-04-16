package util

const (
	TypeValidationError     = "Validation Error"
	TypeInternalServerError = "Internal Server Error"
	TypeBadRequest          = "Bad Request"
	TypeNotFound            = "Not Found"
	TypeConflict            = "Conflict"
)

const (
	SharedErrorTitleInvalidName            = "Nome inválido"
	SharedErrorTitleInvalidBirthDate       = "Data de nascimento inválida"
	SharedErrorTitleInvalidNationality     = "Nacionalidade inválida"
	SharedErrorTitleInvalidImageID         = "ID de imagem inválido"
	SharedErrorTitleErrorChangingLogin     = "Erro ao alterar login"
	SharedErrorTitleErrorChangingAddress   = "Erro ao alterar endereço"
	SharedErrorTitleErrorChangingBirthDate = "Erro ao alterar data de aniversário"
	SharedErrorTitleErrorChangingImageID   = "Erro ao alterar ID da imagem"
	SharedErrorTitleErrorChangingName      = "Erro ao alterar nome"
	SharedErrorTitleErrorImageSize         = "Tamanho inválido"
	SharedErrorTitleInvalidDescription     = "Descrição inválida"
	SharedErrorTitleInvalidYear            = "Ano do filme inválido"
	SharedErrorTitleInvalidIDs             = "Existe um ou mais IDs inválidos"
)

const (
	ChooserErrorDetailEmptyName     = "O nome do Chooser não pode estar vazio"
	ChooserErrorDetailMaxLengthName = "O nome do Chooser não pode ter mais do que 100 caracteres"
)

const (
	ActorErrorDetailEmptyName          = "O nome do(a) ator(atriz) não pode estar vazio"
	ActorErrorDetailMaxLengthName      = "O nome do(a) ator(atriz) não pode ter mais do que 100 caracteres"
	ActorErrorDetailNotNullBirthDate   = "A data de nascimento do(a) ator(atriz) não pode ser nula"
	ActorErrorDetailNotNullNationality = "A nacionalidade do(a) ator(atriz) não pode ser nula"
	ActorErrorDetailEmptyImageID       = "O ID de imagem do(a) ator(atriz) não pode estar vazio"
)

const (
	DirectorErrorDetailEmptyName          = "O nome do(a) diretor(a) não pode estar vazio"
	DirectorErrorDetailMaxLengthName      = "O nome do(a) diretor(a) não pode ter mais do que 100 caracteres"
	DirectorErrorDetailNotNullBirthDate   = "A data de nascimento do(a) diretor(a) não pode ser nula"
	DirectorErrorDetailNotNullNationality = "A nacionalidade do(a) diretor(a) não pode ser nula"
	DirectorErrorDetailEmptyImageID       = "O ID de imagem do(a) diretor(a) não pode estar vazio"
)

const (
	GenreErrorDetailEmptyName     = "O nome do gênero não pode estar vazio"
	GenreErrorDetailMaxLengthName = "O nome do gênero não pode ter mais do que 100 caracteres"
	GenreErrorDetailEmptyImageID  = "O ID de imagem do gênero não pode estar vazio"
)

const (
	ImageErrorDetailEmptyName = "O nome da imagem não pode estar vazio"
	ImageErrorDetailImageSize = "O tamanho da imagem deve estar entre 1 e 100000 bytes"
)

const (
	ListErrorDetailEmptyName        = "O nome da lista não pode estar vazio e deve ter no máximo 100 caracteres"
	ListErrorDetailEmptyDescription = "A descrição da lista não pode estar vazia e deve ter no máximo 150 caracteres"
)

const (
	MovieErrorDetailEmptyImageID = "O ID do filme não pode estar vazio"
	MovieErrorDetailInvalidYear  = "O ano do filme deve estar entre 1800 e "
)

const (
	VoteErrorDetailEmptyIDs = "Para registrar uma votação os IDs não podem estar vazios"
)
