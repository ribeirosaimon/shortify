package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ribeirosaimon/shortify/config/server"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/shortify/internal/usecase"
	"github.com/ribeirosaimon/tooltip/response"
)

type UrlRecord struct {
	urlRecord usecase.UrlRecord
}

func NewUrlRecord() *UrlRecord {
	return &UrlRecord{
		urlRecord: server.GetUrlRecordUsecase(),
	}
}

// NewUrlRecord
// @Summary Post url
// @Description Create one urlRecord in database
// @Produce  json
// @Consume  json
// @Router /url-record [post]
// @Success 201 {object} dto.UrlRecord "Success"
func (u *UrlRecord) NewUrlRecord(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var urlRecordDto dto.UrlRecord

	if err := json.NewDecoder(r.Body).Decode(&urlRecordDto); err != nil {
		response.BadRequest(w, err)
	}

	urlCreated, err := u.urlRecord.Create(ctx, &urlRecordDto)
	if err != nil {
		response.BadRequest(w, err)
	}

	response.Created(w, urlResponse{
		urlCreated.GetShortenedUrl().GetValue(),
	})
}

type urlResponse struct {
	HashUrl string `json:"hash_url"`
}
