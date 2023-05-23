package WordService

import (
	Models "WordMixBack/src/Model"
	"context"
)

type WordRepository interface {
	AddNewWords(ctx context.Context, words []Models.Word) ([]Models.Word, error)
	GetWords(ctx context.Context) ([]Models.Word, error)
}

type WordService struct {
	wordRepository WordRepository
}

func NewWordService(wordRepository WordRepository) *WordService {
	return &WordService{
		wordRepository: wordRepository,
	}
}

func (s *WordService) AddNewWords(ctx context.Context, words []Models.Word) ([]Models.Word, error) {
	return s.wordRepository.AddNewWords(ctx, words)
}

func (s *WordService) GetWords(ctx context.Context) ([]Models.Word, error) {
	return s.wordRepository.GetWords(ctx)
}
