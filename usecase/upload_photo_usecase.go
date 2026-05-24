package usecase

import "photo-service/usecase/repository"

type UploadPhotoUseCase interface {
	Execute(photoData []byte) (string, error)
}

type uploadPhotoUseCase struct {
	s3Repo repository.S3Repository
}

func NewUploadPhotoUseCase(s3Repo repository.S3Repository) UploadPhotoUseCase {
	return &uploadPhotoUseCase{
		s3Repo: s3Repo,
	}
}

func (u *uploadPhotoUseCase) Execute(photoData []byte) (string, error) {
	return u.s3Repo.UploadPhoto(photoData, "temporary_photo.jpg")
}
