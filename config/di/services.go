package di

type Service string

const (
	// UrlRecordServices UrlRecord
	UrlRecordUseCase    Service = "urlRecordUsecase"
	UrlRecordRepository Service = "UrlRecordRepository"
	UrlRecordCache      Service = "UrlRecordCache"
)
