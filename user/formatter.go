package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		ImageUrl:   user.AvatarFilename,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}
