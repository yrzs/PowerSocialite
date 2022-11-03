package weCom

import "github.com/yrzs/PowerSocialite/src/models"

type ResponseGetExternalContact struct {
	*ResponseWeCom
	ExternalContact *models.ExternalContact `json:"external_contact"`
	FollowUsers     []*models.FollowUser    `json:"follow_user"`
	NextCursor      string                  `json:"next_cursor"`
}
