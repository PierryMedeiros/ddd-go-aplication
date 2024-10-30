package valueobject

import (
	"errors"
	"fmt"
)

type Address struct {
	Street string
	Number int
	Zip    string
	City   string
}

func NewAddress(street string, number int, zip string, city string) (*Address, error) {
	address := &Address{
		Street: street,
		Number: number,
		Zip:    zip,
		City:   city,
	}

	err := address.Validate()
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (a *Address) GetStreet() string {
    return a.Street
}

func (a *Address) GetNumber() int {
    return a.Number
}

func (a *Address) GetZip() string {
    return a.Zip
}

func (a *Address) GetCity() string {
    return a.City
}

func (a *Address) Validate() error {
	if len(a.Street) == 0 {
		return errors.New("street is required")
	}
	if a.Number == 0 {
		return errors.New("number is required")
	}
	if len(a.Zip) == 0 {
		return errors.New("zip is required")
	}
	if len(a.City) == 0 {
		return errors.New("city is required")
	}
	return nil
}

func (a *Address) ToString() string {
	return fmt.Sprintf("%s, %d, %s %s", a.Street, a.Number, a.Zip, a.City)
}
