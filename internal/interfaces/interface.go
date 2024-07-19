package interfaces

// import "pokemon/internal/domain"

// type Constraint interface {
// 	domain.Player | domain.Pokemon
// }

type InterfaceRepository[T any] interface {
	GetAll[T]
	GetById[T]
	Save[T]
	UpdateById[T]
	DeleteById
}

type GetAll[T any] interface {
	GetAll() ([]T, error)
}

type GetById[T any] interface {
	GetById(int) (T, error)
}

type Save[T any] interface {
	Save(*T) error
}

type UpdateById[T any] interface {
	UpdateById(*T) error
}

type DeleteById interface {
	DeleteById(int) error
}
