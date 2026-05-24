package repository

type S3Repository interface {
	UploadPhoto(photoData []byte, photoName string) (string, error)
	GetPresignedURL(photoName string) (string, error)
}