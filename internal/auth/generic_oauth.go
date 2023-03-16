package auth

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmespath/go-jmespath"
)

func extractInfo(searchPath string, data map[string]any) (string, error) {
	result, err := jmespath.Search(searchPath, data)
	if err != nil {
		return "", err
	}
	val, ok := result.(string)
	if !ok {
		newVal, ok := result.(int)
		if !ok {
			return "", errors.New("Invalid value")
		}
		val = strconv.Itoa(int(newVal))
	}

	return val, nil
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
		if email, err = extractInfo(emailPath, claim); err != nil {
			return nil, err
		}

		if uid, err = extractInfo(userIdPath, claim); err != nil {
			return nil, err
		}

		if name, err = extractInfo(displayNamePath, claim); err != nil {
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

	if email, err = extractInfo(emailPath, userInfo); err != nil {
		return nil, err
	}

	if uid, err = extractInfo(userIdPath, userInfo); err != nil {
		return nil, err
	}

	if name, err = extractInfo(displayNamePath, userInfo); err != nil {
		return nil, err
	}

	return map[string]any{
		"uid":   uid,
		"email": email,
		"name":  name,
	}, nil
}
