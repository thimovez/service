package usecase

type (
	User interface {
		Login(username, password string) error
	}

	UserRepo interface {
		SaveUser(id int64, username, password string) error
	}
)
