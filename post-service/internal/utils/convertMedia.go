package utils

import (
	post "github.com/billzayy/social-media/post-service/api"
	"github.com/billzayy/social-media/post-service/internal/models"
)

func ConvertMedia(m []models.PostMedia) []*post.PostMedia {
	var mediaList []*post.PostMedia
	for _, item := range m {
		mediaList = append(mediaList, &post.PostMedia{
			PostId: item.PostId.String(),
			Url:    item.Url,
			Type:   item.Type,
		})
	}
	return mediaList
}
