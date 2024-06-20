package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const defaultMaxBodyCharacters = 32

// keyMaxBodyCharacters is the max mail body characters
var keyMaxBodyCharacters = []byte("MaxBodyCharacters")

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(maxBodyCharacters uint32) Params {
	return Params{
		MaxBodyCharacters: maxBodyCharacters,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(defaultMaxBodyCharacters)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			keyMaxBodyCharacters,
			&p.MaxBodyCharacters,
			validateMaxBodyCharacters,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxBodyCharacters(p.MaxBodyCharacters); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMaxBodyCharacters(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("%s: %T", "invalid parameter type", i)
	}

	if v <= 0 {
		return fmt.Errorf("%s: %d", "max body characters should be more than zero", v)
	}

	return nil
}
