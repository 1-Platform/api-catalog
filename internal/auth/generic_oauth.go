package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmespath/go-jmespath"
)

func extractInfo[V any](searchPath string, data map[string]any) (V, error) {
	result, err := jmespath.Search(searchPath, data)
	if err != nil {
		var t V
		return t, err
	}
	t, ok := result.(V)
	if !ok {
		return t, errors.New("Invalid value")
	}
	return t, nil
}

func GenericOauthUserInfo(token string, userInfoUrl string,
	emailPath, userIdPath, displayNamePath string) (map[string]any, error) {
	var email, uid, name string

	req, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.Header.Get("Content-Type") == "application/jwt" {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		claim := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(string(data), claim, nil)
		fmt.Println(claim["email"])
		if email, err = extractInfo[string](emailPath, claim); err != nil {
			return nil, err
		}

		if uid, err = extractInfo[string](userIdPath, claim); err != nil {
			return nil, err
		}

		if name, err = extractInfo[string](displayNamePath, claim); err != nil {
			return nil, err
		}
		return map[string]any{
			"uid":   uid,
			"email": email,
			"name":  name,
		}, nil
	}

	var userInfo map[string]any
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	if email, err = extractInfo[string](emailPath, userInfo); err != nil {
		return nil, err
	}

	if uid, err = extractInfo[string](userIdPath, userInfo); err != nil {
		return nil, err
	}

	if name, err = extractInfo[string](displayNamePath, userInfo); err != nil {
		return nil, err
	}

	return map[string]any{
		"uid":   uid,
		"email": email,
		"name":  name,
	}, nil
}
