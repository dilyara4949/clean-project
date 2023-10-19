package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/clean-project/domain"
)

type articleUsecase struct {
	articleRepository domain.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleRepository(articleRepository domain.ArticleRepository, timeout time.Duration) domain.ArticleUserCase {
	return &articleUsecase{
		articleRepository: articleRepository,
		contextTimeout: timeout,
	}
}

func (tu *articleUsecase) Create(c context.Context, task *domain.Article) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.articleRepository.Create(ctx, task)
}

func (tu *articleUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.articleRepository.FetchByUserID(ctx, userID)
}