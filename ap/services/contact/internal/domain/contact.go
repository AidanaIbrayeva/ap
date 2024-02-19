package domain

import (
	"errors"
	"fmt"
	"strconv"
)

type Contact struct {
	ID          int
	FullName    string
	PhoneNumber string
}

func NewContact(id int, firstName, lastName, middleName, phoneNumber string) (*Contact, error) {

	_, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return nil, errors.New("phone number must contain only digits")
	}

	fullName := fmt.Sprintf("%s %s %s", lastName, firstName, middleName)

	return &Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}, nil
}

func (c *Contact) ReadOnlyFullName() string {
	return c.FullName
}
