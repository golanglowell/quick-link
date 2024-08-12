package memory

import (
	"errors"
	"fmt"

	"github.com/golanglowell/quick-link/internal/domain"
)

type URLRepository struct {
	urls     map[string]*domain.URL
	commands chan command
}

type command interface {
	execute(map[string]*domain.URL)
}

type saveCommand struct {
	url    *domain.URL
	result chan<- error
}

type findCommand struct {
	shortCode string
	result    chan<- findResult
}

type findResult struct {
	url *domain.URL
	err error
}

func (c saveCommand) execute(urls map[string]*domain.URL) {
	urls[c.url.ShortCode] = c.url
	c.result <- nil
}

func (c findCommand) execute(urls map[string]*domain.URL) {
	url, ok := urls[c.shortCode]
	if !ok {
		c.result <- findResult{nil, errors.New("URL not found")}
	} else {
		c.result <- findResult{url, nil}
	}
}

func NewURLRepository() *URLRepository {

	repo := &URLRepository{
		urls:     make(map[string]*domain.URL),
		commands: make(chan command),
	}

	go repo.commandLoop()
	return repo
}

func (u *URLRepository) commandLoop() {
	for cmd := range u.commands {
		cmd.execute(u.urls)
	}
}

func (u *URLRepository) Save(url *domain.URL) error {
	if url == nil {
		return fmt.Errorf("invalid input")
	}

	result := make(chan error)
	u.commands <- saveCommand{url: url, result: result}
	return <-result
}

func (u *URLRepository) FindByShortCode(shortCode string) (*domain.URL, error) {
	result := make(chan findResult)
	u.commands <- findCommand{shortCode: shortCode, result: result}
	res := <-result
	return res.url, res.err
}

func (r *URLRepository) Close() {
	close(r.commands)
}
