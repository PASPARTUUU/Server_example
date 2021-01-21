package handler

import (
	"github.com/PASPARTUUU/Server_example/service/store"
	"github.com/sirupsen/logrus"
)

// Handler -
type Handler struct {
	UserHandler *UserController
}

// New -
func New(storage *store.Store, log *logrus.Logger) *Handler {
	return &Handler{
		UserHandler: NewUsers(storage, log),
	}
}
