package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

type pageToken struct {
	Page int
}

func marshalPageToken(t pageToken) string {
	raw, _ := json.Marshal(t)
	return base64.StdEncoding.EncodeToString(raw)
}

func unmarshalPageToken(t string) pageToken {
	var raw []byte
	if len(t) > 0 {
		raw, _ = base64.StdEncoding.DecodeString(t)
	}

	token := pageToken{}
	if len(raw) != 0 {
		_ = json.Unmarshal(raw, &token)
	}
	return token
}

func Paginate[T any](
	data []T,
	requestedSize *resources.PaginationSize,
	requestedPage *resources.PaginationPage,
	requestedNextToken *resources.PaginationNextToken,
	requestedNextPageToken *resources.PaginationNextTokenHeader,
) (*resources.PaginatedResponse[T], error) {
	if requestedNextPageToken != nil || requestedNextToken != nil {
		// TODO: we consider requestedNextToken and requestedNextPageToken are
		// the same things, but it might not be true.

		// For simplicity we assume that the token is a Base64 encoded value of
		// a simple JSON object like `{"page":0,"size":10}`.

		var raw string
		if requestedNextPageToken != nil {
			raw = *requestedNextPageToken
		} else {
			raw = *requestedNextToken
		}

		token := unmarshalPageToken(raw)
		subset, nextPage := getSubset(data, token.Page, 10)

		var nextPageToken *string
		if nextPage != nil {
			marshalled := marshalPageToken(pageToken{Page: *nextPage})
			nextPageToken = &marshalled
		}

		return &resources.PaginatedResponse[T]{
			Data: subset,
			Meta: resources.ResponseMeta{
				PageToken: &raw,
				Size:      len(subset),
			},
			Next: resources.Next{
				PageToken: nextPageToken,
			},
		}, nil
	}

	// Doing ordinary page/size pagination.
	page := 0
	if requestedPage != nil {
		page = *requestedPage
	}

	size := 0
	if requestedSize != nil {
		size = *requestedSize
	}

	if page < 0 {
		page = 0
	}
	if size <= 0 {
		size = 10
	}

	subset, nextPage := getSubset(data, page, size)
	return &resources.PaginatedResponse[T]{
		Data: subset,
		Meta: resources.ResponseMeta{
			Page: &page,
			Size: len(subset),
		},
		Next: resources.Next{
			Page: nextPage,
		},
	}, nil

}

func getSubset[T any](data []T, page, size int) ([]T, *int) {
	var subset []T
	ix1 := page * size
	ix2 := ix1 + size
	var nextPage *int
	if ix1 >= len(data) {
		subset = []T{}
	} else if ix2 >= len(data) {
		subset = data[ix1:]
	} else {
		subset = data[ix1:ix2]
		next := 1 + page
		nextPage = &next
	}
	return subset, nextPage
}