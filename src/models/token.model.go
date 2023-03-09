package model

type Token struct {
	AccessToken string `bson:"access_token" json:"accessToken"`
	ExpiredAt   string `bson:"expired_at" json:"expiredAt"`
	Disabled    bool   `bson:"disabled" json:"disabled"`
}
