package payload

type (
	UserGet struct {
		ID       *string `query:"id"`
		Name     *string `query:"name"`
		Email    *string `query:"email"`
		Password *string
		Token    *string
	}

	UserCreate struct {
		ID       *string
		Name     string `json:"name" form:"name" validate:"required"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	Login struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	UserUpdate struct {
		ID       string
		Name     *string `json:"name" form:"name"`
		Email    *string `json:"email" form:"email"`
		Token    *string
		Password *string `json:"password" form:"password"`
	}
)
