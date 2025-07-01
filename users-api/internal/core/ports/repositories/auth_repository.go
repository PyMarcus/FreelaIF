package repositories

type IAuthRepository[T any] interface{
	GetUser(email string, password string) (T, error)
}


