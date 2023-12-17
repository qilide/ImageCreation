package account

type LoginBinder struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterBinder struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type MailBinder struct {
	Email string `json:"email"`
}

type ModifyBinder struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsSuperuser string `json:"is_superuser"`
}

type ModifyInformationBinder struct {
	Nickname    string `json:"nickname"`
	Age         string `json:"age"`
	Sex         string `json:"sex"`
	BrithDate   string `json:"brith_date"`
	Avatar      string `json:"avatar"`
	Biography   string `json:"biography"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Style       string `json:"style"`
	Posts       string `json:"posts"`
}
