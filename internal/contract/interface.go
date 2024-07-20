package contract

// import "pokemon/internal/domain"

// type Constraint interface {
// 	domain.Player | domain.Pokemon
// }

type GlobalInterface[T any] interface {
	GetAll[T]
	GetById[T]
	Save[T]
	Update[T]
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

type Update[T any] interface {
	Update(*T) error
}

type DeleteById interface {
	DeleteById(int) error
}
