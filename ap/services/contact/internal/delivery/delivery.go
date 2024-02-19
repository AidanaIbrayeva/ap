package delivery

import (
	"encoding/json"
	"net/http"
	"services/contact/internal/domain"
	"services/contact/internal/usecase"
	"strconv"
)

type ContactDeliveryImpl struct {
	usecase usecase.ContactUseCase
}

func NewContactDelivery(usecase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{
		usecase: usecase,
	}
}

type ContactDelivery interface {
	CreateContactHandler(w http.ResponseWriter, r *http.Request)
	GetContactHandler(w http.ResponseWriter, r *http.Request)
}

func (cd *ContactDeliveryImpl) CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение данных из запроса
	var contact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	contactID, err := cd.usecase.CreateContact(&contact)
	if err != nil {
		http.Error(w, "Failed to create contact", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{ ContactID int }{ContactID: contactID})
}

func (cd *ContactDeliveryImpl) GetContactHandler(w http.ResponseWriter, r *http.Request) {
	contactID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(contactID)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := cd.usecase.GetContact(id)
	if err != nil {
		http.Error(w, "Failed to get contact", http.StatusInternalServerError)
		return
	}

	// Отправка ответа с данными контакта
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}
