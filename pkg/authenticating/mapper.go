package authenticating

func ToUser(userDTO UserDTO) User {
	return User{
		UserName: userDTO.UserName,
		FirstName: userDTO.FirstName,
		LastName: userDTO.LastName,
		Password: userDTO.Password,
	}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{
		ID: user.ID,
		UserName: user.UserName,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Password: user.Password,
	}
}
