package main

import (
	"net/http"
	"photo-service/handler"
	"photo-service/repository"
	"photo-service/repository/source"
	"photo-service/usecase"
)

func main() {
	mux := http.NewServeMux()

	s3Source := source.NewS3Source("ap-southeast-2")
	s3Repository := repository.NewS3Repository(s3Source)
	uploadPhotoHandler := handler.NewUploadPhotoHandler(usecase.NewUploadPhotoUseCase(s3Repository))

	mux.HandleFunc("/upload", uploadPhotoHandler.HandleUploadPhoto)

	http.ListenAndServe(":8080", mux)
}
