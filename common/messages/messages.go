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

	ApplicationListResponseMessage struct {
		Data []ApplicationRecord `json:"data"`
	}

	FeedbackMessage struct {
		Data FeedbackRecord `json:"data"`
	}

	FeedbackListRequestMessage struct {
		Data ApplicationRecord `json:"data"`
	}

	FeedbackListResponseMessage struct {
		Data []FeedbackRecord `json:"data"`
	}

	FeedbackRecord struct {
		ID            uint   `json:"id"`
		Desc          string `json:"desc"`
		ApplicationId uint   `json:"appid"`
		Email         string `json:"email"`
	}
)
