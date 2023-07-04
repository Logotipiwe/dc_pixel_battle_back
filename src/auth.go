package main

import (
	"github.com/Logotipiwe/dc_go_auth_lib/auth"
	"net/http"
	"os"
)

/*type User struct {
	Sub        string `json:"sub"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}*/

/*func GetAccessTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	var accessToken string
	if cookie != nil {
		accessToken = cookie.Value
	} else {
		accessToken = ""
	}
	return accessToken, nil
}*/

func GetUserData(r *http.Request) (*auth.User, error) {
	println("Getting user data...")
	if os.Getenv("autoAuth") == "1" {
		println("Auto auth enabled")
		return &auth.User{
			Sub:  os.Getenv("LOGOTIPIWE_GMAIL_ID"),
			Name: "Reman Gerus",
		}, nil
	}
	user, err := auth.GetUserData(r)
	return &user, err
	//accessToken, err := GetAccessTokenFromCookie(r)
	//if err != nil {
	//	return nil, err
	//}
	//println("Access token is : " + accessToken)
	//user, err := GetUserDataFromToken(accessToken)
	//if err == nil {
	//	println(fmt.Sprintf("User got with id %s and name %s", user.Sub, user.Name))
	//}
	//return user, err
}

/*func GetUserDataFromToken(accessToken string) (*User, error) {
	bearer := "Bearer " + accessToken
	getUrl := "https://www.googleapis.com/oauth2/v3/userinfo"
	request, _ := http.NewRequest("GET", getUrl, nil)
	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, _ := client.Do(request)
	defer res.Body.Close()
	var answer User
	err := json.NewDecoder(res.Body).Decode(&answer)
	if err != nil {
		return nil, err
	}
	if answer.Sub != "" {
		return &answer, nil
	} else {
		return &answer, errors.New("WTF HUH")
	}
}*/

/*func GetLoginUrl() string {
	loginUrl, _ := url.Parse(env.GetCurrUrl() + "/oauth2/auth")
	q := loginUrl.Query()
	q.Set("redirect", env.GetCurrUrl()+env.GetSubpath())
	loginUrl.RawQuery = q.Encode()

	return loginUrl.String()
}
*/
