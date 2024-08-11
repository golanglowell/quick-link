package application

import "github.com/golanglowell/quick-link/internal/domain"

type GetLinkUseCase struct {
	urlRepo domain.URLRepository
}

func NewGetLink(urlRepo domain.URLRepository) *GetLinkUseCase {
	return &GetLinkUseCase{urlRepo: urlRepo}
}

func (uc *GetLinkUseCase) Execute(shortCode string) (*domain.URL, error) {
	return uc.urlRepo.FindByShortCode(shortCode)
}
