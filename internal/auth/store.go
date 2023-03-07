package auth

import (
	"context"

	"github.com/akhilmhdh/authy"
	oauth2Core "github.com/akhilmhdh/authy/modules/oauth2/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthStore struct {
	mongoDriver    *mongo.Database
	collection     *mongo.Collection
	collectionName string
}

func NewAuthStore(db *mongo.Database) *AuthStore {
	return &AuthStore{
		mongoDriver:    db,
		collectionName: "auth",
		collection:     db.Collection("auth"),
	}
}

func (a *AuthStore) GetUserByPID(ctx context.Context, pid string) (authy.User, error) {
	var user User
	provider, uid, err := oauth2Core.ParseOAuth2PID(pid)

	if err == nil {
		filter := bson.M{"oauth2_uid": uid, "oauth2_provider": provider}
		err := a.collection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return nil, authy.ErrUserNotFound
			}
			return nil, err
		}

		return &user, nil
	}

	return nil, authy.ErrUserNotFound
}

func (a *AuthStore) NewOauth2User(ctx context.Context,
	provider string, userInfo map[string]any) (oauth2Core.Oauth2User, error) {
	if userInfo == nil {
		return &User{}, nil
	}

	return &User{
		Oauth2Provider: provider,
		Name:           userInfo["name"].(string),
		Oauth2Uid:      userInfo["uid"].(string),
		Email:          userInfo["email"].(string),
	}, nil
}

func (a *AuthStore) SaveOauth2User(ctx context.Context, us oauth2Core.Oauth2User) error {
	user := us.(*User)

	filter := bson.M{"oauth2_uid": user.Oauth2Uid, "oauth2_provider": user.Oauth2Provider}
	update := bson.M{"$set": user}
	opts := options.Update().SetUpsert(true)

	if _, err := a.collection.UpdateOne(ctx, filter, update, opts); err != nil {
		return err
	}

	return nil
}
