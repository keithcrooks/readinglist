package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float32  `json:"rating"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BooksResponse struct {
	Books *[]Book `json:"books"`
}

type ReadingListModel struct {
	Endpoint string
}

func (m *ReadingListModel) GetAll() (*[]Book, error) {
	res, err := http.Get(m.Endpoint)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bookRes BooksResponse
	err = json.Unmarshal(data, &bookRes)
	if err != nil {
		return nil, err
	}

	return bookRes.Books, nil
}

func (m *ReadingListModel) Get(id int64) (*Book, error) {
	url := fmt.Sprintf("%s/%d", m.Endpoint, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var bookRes BookResponse
	err = json.Unmarshal(data, &bookRes)
	if err != nil {
		return nil, err
	}

	return bookRes.Book, nil
}
