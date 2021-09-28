package book_files

import (
	"books_online_api/business/books"
	"books_online_api/helpers"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

type BookFilesLocal struct {
	Repo           books.Repository
	ContextTimeout time.Duration
}

func (b BookFilesLocal) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {

	srcFileBooks, err := domain.FileBook.Open()
	if err != nil {
		return books.Domain{}, errors.New("internal server error")
	}
	defer srcFileBooks.Close()

	fileNameBook := domain.FileBook.Filename
	splitFileNameBook := strings.Split(fileNameBook, ".")
	// Check page extension page book
	extentionFileBook := splitFileNameBook[len(splitFileNameBook)-1]
	validExtensionFileBook := []string{"pdf"}

	if !helpers.CheckItemInSlice(validExtensionFileBook, extentionFileBook) {
		return books.Domain{}, errors.New("input file book no pdf")
	}
	newFileNameBook := time.Now().Format(time.RFC850) + "_" + helpers.RandomString(10) + "_" + splitFileNameBook[0] + "." + extentionFileBook
	dstFileBook, err := os.Create("./public/files/file_books/" + newFileNameBook)

	if err != nil {
		return books.Domain{},  errors.New("internal server error")
	}
	defer  dstFileBook.Close()

	if _, err := io.Copy(dstFileBook, srcFileBooks); err != nil {
		return books.Domain{}, err
	}

	domain.UrlBook = newFileNameBook

	bookDb, err := b.Repo.CreateBook(ctx, domain)

	if err != nil {
		return books.Domain{}, err
	}

	return bookDb, nil
}

func NewBookFileLocal(repo books.Repository, timeout time.Duration) books.Localy {
	return &BookFilesLocal{
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

