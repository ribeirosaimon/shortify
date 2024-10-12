package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ribeirosaimon/shortify/config/di"
	"github.com/ribeirosaimon/shortify/internal/dto"
	"github.com/ribeirosaimon/shortify/internal/usecase"
	"github.com/ribeirosaimon/tooltip/response"
)

type UrlRecord struct {
	urlRecord usecase.UrlRecord
}

func NewUrlRecord() *UrlRecord {
	usecase.NewUrlRecord()
	return &UrlRecord{
		urlRecord: di.GetRegistry().Inject(di.UrlRecordUseCase).(usecase.UrlRecord),
	}
}

// NewUrlRecord
// @Summary Post UrlRecord
// @Description Create one urlRecord in database
// @Produce  json
// @Consume  json
// @Router /url-record [get]
// @Success 200 {UrlRecordDto} UrlRecordDto
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
