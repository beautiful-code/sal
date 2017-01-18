package messages

type (
	UserMessage struct {
		Data UserRecord `json:"data"`
	}

	UserRecord struct {
		ID        uint   `json:"id"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	AuthUserMessage struct {
		Token string `json:"token"`
	}

	ApplicationMessage struct {
		Data ApplicationRecord `json:"data"`
	}

	ApplicationRecord struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)
