package daos

type UserLoginParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UserData string `json:"userData,omitempty"`
}

func (z *Context) userLogin(params UserLoginParams) (string, int, error) {

	var result string

	status, err := z.request("user.login", params, &result)
	if err != nil {
		return "", status, err
	}

	return result, status, nil
}
