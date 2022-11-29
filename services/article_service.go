package services

import (
	"database/sql"
	"sync"

	"github.com/Riku0617/myapi/models"
	"github.com/Riku0617/myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var err error

	var amn sync.Mutex
	var cmn sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func(db *sql.DB, articleID int){
		defer wg.Done()
		amn.Lock()
		article, err = repositories.SelectArticleDetail(db, articleID)
		amn.Unlock()
	}(s.db,articleID)
	if err != nil {
		return models.Article{}, err
	}

	go func(db *sql.DB, articleID int){
		defer wg.Done()
		cmn.Lock()
		commentList, err = repositories.SelectCommentList(db, articleID)
		cmn.Unlock()
	}(s.db,articleID)
	
	wg.Wait()
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articles, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
