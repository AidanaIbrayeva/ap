package usecase

import (
	"errors"
	"log"
	"services/contact/internal/domain"
	"services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{
		repository: repository,
	}
}

func (c *ContactUseCaseImpl) CreateContact(contact *domain.Contact) (int, error) {
	if contact == nil {
		return 0, errors.New("contact is nil")
	}

	id, err := c.repository.CreateContact(contact)
	if err != nil {
		log.Printf("Failed to create contact: %v", err)
		return 0, err
	}

	return id, nil
}

func (c *ContactUseCaseImpl) GetContact(id int) (*domain.Contact, error) {
	contact, err := c.repository.GetContact(id)
	if err != nil {
		log.Printf("Failed to get contact: %v", err)
		return nil, err
	}

	return contact, nil
}

func (c *ContactUseCaseImpl) UpdateContact(contact *domain.Contact) error {
	if contact == nil {
		return errors.New("contact is nil")
	}

	err := c.repository.UpdateContact(contact)
	if err != nil {
		log.Printf("Failed to update contact: %v", err)
		return err
	}

	return nil
}

func (c *ContactUseCaseImpl) DeleteContact(id int) error {
	err := c.repository.DeleteContact(id)
	if err != nil {
		log.Printf("Failed to delete contact: %v", err)
		return err
	}

	return nil
}

func (c *ContactUseCaseImpl) CreateGroup(group *domain.Group) (int, error) {
	if group == nil {
		return 0, errors.New("group is nil")
	}

	id, err := c.repository.CreateGroup(group)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		return 0, err
	}

	return id, nil
}

func (c *ContactUseCaseImpl) GetGroup(id int) (*domain.Group, error) {
	group, err := c.repository.GetGroup(id)
	if err != nil {
		log.Printf("Failed to get group: %v", err)
		return nil, err
	}

	return group, nil
}

func (c *ContactUseCaseImpl) AddContactToGroup(contactID, groupID int) error {
	err := c.repository.AddContactToGroup(contactID, groupID)
	if err != nil {
		log.Printf("Failed to add contact to group: %v", err)
		return err
	}

	return nil
}
