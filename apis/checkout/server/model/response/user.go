package response

/*****
struct for posts
******/

type PostUser struct {
	Token string `json:"token"`
}

type PostUserException struct {
	Field   string `json:"field" enums:"COUNTRY,EMAIL,PASSWORD,FORM"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

/*****
struct for posts
******/
