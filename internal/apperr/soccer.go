package apperr

var (
	ErrTeamNameRequired = NewBadRequestError("team name is required")
)
