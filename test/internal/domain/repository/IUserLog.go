package repository

type UserLogger interface {
	Create(Modeler) error
}
