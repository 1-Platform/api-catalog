package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"pid"`
	Name  string             `json:"name"`

	Oauth2Uid          string    `bson:"oauth2_uid"`
	Oauth2Provider     string    `bson:"oauth2_provider"`
	Oauth2AccessToken  string    `bson:"oauth2_access_token"`
	Oauth2RefreshToken string    `bson:"oauth2_refresh_token"`
	Oauth2Expiry       time.Time `bson:"oauth2_expiry"`
}

func (u *User) GetPID() string {
	return u.Email
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) PutArbitoryData(data map[string]any) {
	if name, ok := data["name"]; ok {
		if n, ok := name.(string); ok {
			u.Name = n
		}
	}
}

func (u *User) GetOAuth2UID() string {
	return u.Oauth2Uid
}

func (u *User) PutPID(pid string) {
	u.Email = pid
}

func (u *User) PutOAuth2UID(uid string) {
	u.Oauth2Uid = uid
}

func (u *User) PutOAuth2Provider(provider string) {
	u.Oauth2Provider = provider
}

func (u *User) PutOAuth2AccessToken(token string) {
	u.Oauth2AccessToken = token
}

func (u *User) PutOAuth2RefreshToken(refreshToken string) {
	u.Oauth2RefreshToken = refreshToken
}

func (u *User) PutOAuth2Expiry(expiry time.Time) {
	u.Oauth2Expiry = expiry
}
