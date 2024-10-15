package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ribeirosaimon/shortify/config/db"
	"github.com/ribeirosaimon/shortify/internal/entity"
	"github.com/ribeirosaimon/tooltip/storage/pgsql"
	"github.com/ribeirosaimon/tooltip/tlog"
)

var urlTable = "url_records"

type UrlRecordRepository interface {
	GetUrlRecord(ctx context.Context, id string) (*entity.UrlRecord, error)
	InsertUrlRecord(ctx context.Context, urlRecord *entity.UrlRecord) (*entity.UrlRecord, error)
}

// NewUrl is once open function
func NewUrl() *urlRepositoryImpl {
	return &urlRepositoryImpl{
		conn: db.NewPgsqlConnection(),
	}
}

type urlRepositoryImpl struct {
	conn pgsql.PConnInterface
}

func (u *urlRepositoryImpl) InsertUrlRecord(ctx context.Context, urlRecord *entity.UrlRecord) (*entity.UrlRecord, error) {
	query := fmt.Sprintf("INSERT INTO %s (created_at, id, original_url, shortened_url) VALUES ($1, $2, $3, $4)", urlTable)
	tlog.Debug("CreateUrlRecord", query)

	exec, err := u.conn.GetConnection().Exec(query, urlRecord.GetCreatedAt(), urlRecord.GetId().GetValue(), urlRecord.GetOriginalUrl().GetValue(), urlRecord.GetShortenedUrl().GetValue())
	if err != nil {
		tlog.Error("CreateUrlRecord", "insert url error", err)
		return nil, fmt.Errorf("insert url error: %w", err)
	}
	if _, err = exec.RowsAffected(); err != nil {
		tlog.Error("CreateUrlRecord", "rows affected error", err)
	}
	return urlRecord, nil
}

func (u *urlRepositoryImpl) GetUrlRecord(ctx context.Context, id string) (*entity.UrlRecord, error) {
	query := "SELECT u.id, u.original_url, u.shortened_url, u.created_at  FROM url_records u WHERE id = $1"
	tlog.Debug("GetUrlRecord", query)
	var urlDao struct {
		Id           string    `json:"id"`
		OriginalUrl  string    `json:"original_url"`
		CreatedAt    time.Time `json:"created_at"`
		ShortenedUrl string    `json:"shortened_url"`
	}

	if err := u.conn.GetConnection().QueryRow(query, id).
		Scan(&urlDao.Id, &urlDao.OriginalUrl, &urlDao.ShortenedUrl, &urlDao.CreatedAt); err != nil {
		tlog.Error("GetUrlRecord", "find url error", err)
		return nil, err
	}

	return entity.TranformInUrlRecord(urlDao.Id, urlDao.OriginalUrl, urlDao.ShortenedUrl, urlDao.CreatedAt), nil
}
