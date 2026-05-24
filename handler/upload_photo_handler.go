package handler

import (
	"io"
	"net/http"
	"photo-service/usecase"
)

type UploadPhotoHandler struct {
	UploadPhotoUseCase usecase.UploadPhotoUseCase
}

func NewUploadPhotoHandler(uploadPhotoUseCase usecase.UploadPhotoUseCase) *UploadPhotoHandler {
	return &UploadPhotoHandler{
		UploadPhotoUseCase: uploadPhotoUseCase,
	}
}

func (h *UploadPhotoHandler) HandleUploadPhoto(w http.ResponseWriter, r *http.Request) {
	imageData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read image data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	str, err := h.UploadPhotoUseCase.Execute(imageData)
	if err != nil {
		println("Error occurred while uploading photo - error: ", err.Error())
		http.Error(w, "Failed to upload photo", http.StatusInternalServerError)
		w.Write([]byte("Error: Failed to upload photo"))
		return
	}
	w.Write([]byte(str))
}