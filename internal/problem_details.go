package internal

import "errors"

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance,omitempty"`
}

func NewProblemDetails(t string, title string, status int, detail string, instance string) (*ProblemDetails, error) {
	pd := ProblemDetails{
		Type:     t,
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: instance,
	}

	if err := pd.validate(); err != nil {
		return nil, err
	}

	return &pd, nil
}

func (pd *ProblemDetails) validate() error {
	if pd.Type == "" || len(pd.Type) > 100 {
		return errors.New("type deve ser não vazio e ter no máximo 100 caracteres")
	}

	if pd.Title == "" || len(pd.Title) > 100 {
		return errors.New("title deve ser não vazio e ter no máximo 100 caracteres")
	}

	if pd.Status < 100 || pd.Status >= 600 {
		return errors.New("status deve ser um código HTTP válido")
	}

	if pd.Detail == "" || len(pd.Detail) > 255 {
		return errors.New("detail deve ser não vazio e ter no máximo 255 caracteres")
	}

	if len(pd.Instance) > 255 {
		return errors.New("instance não deve ter mais do que 255 caracteres")
	}

	return nil
}