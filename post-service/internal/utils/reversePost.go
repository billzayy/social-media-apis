package utils

import "github.com/billzayy/social-media/post-service/internal/models"

// ? Using Merge Sort on from : https://blog.boot.dev/golang/merge-sort-golang/
func SortPostWithCreatedTime(items []models.PostResp) []models.PostResp {
	if len(items) < 2 {
		return items
	}
	first := SortPostWithCreatedTime(items[:len(items)/2])
	second := SortPostWithCreatedTime(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []models.PostResp, b []models.PostResp) []models.PostResp {
	final := []models.PostResp{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i].CreatedAt > b[j].CreatedAt { // Reverse The algorithms
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
