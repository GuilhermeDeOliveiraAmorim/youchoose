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
)

const (
	ActorErrorDetailEmptyName          = "O nome do(a) ator(atriz) não pode estar vazio"
	ActorErrorDetailMaxLengthName      = "O nome do(a) ator(atriz) não pode ter mais do que 100 caracteres"
	ActorErrorDetailNotNullBirthDate   = "A data de nascimento do(a) ator(atriz) não pode ser nula"
	ActorErrorDetailNotNullNationality = "A nacionalidade do(a) ator(atriz) não pode ser nula"
	ActorErrorDetailEmptyImageID       = "O ID de imagem do(a) ator(atriz) não pode estar vazio"
)

const (
	ChooserErrorDetailEmptyName     = "O nome do Chooser não pode estar vazio"
	ChooserErrorDetailMaxLengthName = "O nome do Chooser não pode ter mais do que 100 caracteres"
)
