package request

/*****
struct for posts
******/

type PostUser struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

/*****
struct for posts
******/
