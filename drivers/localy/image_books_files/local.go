package image_books_files

import (
	"books_online_api/business/books"
	"books_online_api/helpers"
	"context"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type ImageBooksLocal struct {
	Repo books.ImageBooksRepository
	ContextTime time.Duration
}

func (i ImageBooksLocal) UploadImages(ctx context.Context, domain books.Domain) (books.Domain, error) {
	var err error

	domain.UrlCover, err,_ = UploadImages(domain.FileCover)

	resultRepo, err := i.Repo.UploadImages(ctx, domain)

	return resultRepo, err
}

func NewImageBookLocal(repo books.ImageBooksRepository, timeout time.Duration) books.ImageBooksLocaly {
	return &ImageBooksLocal{
		Repo:        repo,
		ContextTime: timeout,
	}
}

func UploadImages(images []*multipart.FileHeader) ([]string, error, string) {
	result := []string{}
	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return result, err, "Internet Server Error"
		}
		defer src.Close()

		fileName := image.Filename
		splitFileName := strings.Split(fileName, ".")

		// Check image extention
		extention := splitFileName[len(splitFileName)-1]
		validExtensionImage := []string{"jpg", "jpeg", "png"}
		if !helpers.CheckItemInSlice(validExtensionImage, extention) {
			return result, nil, "No input image"
		}
		newFileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + helpers.RandomString(10) + "." + extention

		dst, err := os.Create("./public/img/img_books/" + newFileName)
		if err != nil {
			return result, err, "Internal server error"
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return result, err, "Internal server error"
		}

		result = append(result, newFileName)
	}
	return result, nil, "Success Upload Gambar"
}
