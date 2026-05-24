package repository

import "photo-service/repository/source"

type s3Repository struct {
	s3Source source.S3Source
}

func NewS3Repository(s3Source source.S3Source) *s3Repository {
	return &s3Repository{
		s3Source: s3Source,
	}
}

func (r *s3Repository) UploadPhoto(photoData []byte, photoName string) (string, error) {
	return r.s3Source.UploadPhoto(photoData, photoName)
}

func (r *s3Repository) GetPresignedURL(photoName string) (string, error) {
	return r.s3Source.GetPresignedURL(photoName)
}
