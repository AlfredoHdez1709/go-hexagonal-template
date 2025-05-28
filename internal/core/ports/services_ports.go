package ports

type ServiceHub struct {
	UserService IUserService
}

func NewServiceHub(user IUserService) ServiceHub {
	return ServiceHub{
		UserService: user,
	}
}
