package pagination

import (
	"fmt"
	custom_error "songs/internal/common/custom_error"
)

type Params struct {
	Limit  int
	Offset int
}

type ParamsOptional struct {
	Limit  *int
	Offset *int
}

type Constraints struct {
	MinLimit  int
	MaxLimit  int
	MaxOffset int
}

type Defaults struct {
	Limit  int
	Offset int
}

func (p ParamsOptional) MergeDefaults(defaults Defaults) Params {
	var params Params
	if p.Limit == nil {
		params.Limit = defaults.Limit
	} else {
		params.Limit = *p.Limit
	}
	if p.Offset == nil {
		params.Offset = defaults.Offset
	} else {
		params.Offset = *p.Offset
	}
	return params
}

func (p *Params) CheckConstraints(constraints Constraints) error {
	if p.Limit < constraints.MinLimit {
		return custom_error.NewBadInputError(
			fmt.Errorf("limit must be greater than or equal to %d", constraints.MinLimit),
		)
	}
	if p.Limit > constraints.MaxLimit {
		return custom_error.NewBadInputError(
			fmt.Errorf("limit must be less than or equal to %d", constraints.MaxLimit),
		)
	}
	if p.Offset < 0 {
		return custom_error.NewBadInputError(
			fmt.Errorf("offset must be greater than or equal to 0"),
		)
	}
	if p.Offset > constraints.MaxOffset {
		return custom_error.NewBadInputError(
			fmt.Errorf("offset must be less than or equal to %d", constraints.MaxOffset),
		)
	}
	return nil
}

func NewPaginationParams(limitOrNil, offsetOrNil *int, defaultLimit int) Params {
	limit := defaultLimit
	if limitOrNil != nil {
		limit = *limitOrNil
	}

	offset := 0
	if offsetOrNil != nil {
		offset = *offsetOrNil
	}

	return Params{limit, offset}
}

// func (p *PaginationParams) Validate() error {
// 	return validator.Validate.Struct(p)
// }

func StripSlice[S ~[]E, E any](s S, pagParams Params) S {
	start := min(pagParams.Offset, len(s))
	end := min(start+pagParams.Limit, len(s))
	return s[start:end]
}
