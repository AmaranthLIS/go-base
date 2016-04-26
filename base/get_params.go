package base

import (
	"github.com/asaskevich/govalidator"
	"github.com/mcuadros/go-defaults"
)

// This function fills in get parameter defaults and validates them.

// TODO: ensure this will set values in place
func ValidateGetParams(obj interface{}) error {
	defaults.SetDefaults(obj)
	_, err := govalidator.ValidateStruct(obj)
	return err
}
