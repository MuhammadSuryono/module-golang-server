package validate

import (
	"fmt"
	"testing"
)

type person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

func TestValidate(t *testing.T) {
	validate := NewValidate()

	p := person{
		Name:                "Joel",
		Email:               "dummyemail@email.com",
		Age:                 32,
		DriverLicenseNumber: "",
	}
	err := validate.ValidationStruct(p).Error
	fmt.Println(err)
}
