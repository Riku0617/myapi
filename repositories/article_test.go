package repositories_test

import (
	"testing"

	"github.com/Riku0617/myapi/models"
	"github.com/Riku0617/myapi/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		want      models.Article
	}{
		{
			testTitle: "subtest1",
			want: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		},
		{
			testTitle: "subtest2",
			want: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.want.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.want.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.want.ID)
			}
			if got.Title != test.want.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.want.Title)
			}
			if got.Contents != test.want.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.want.Contents)
			}
			if got.UserName != test.want.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.want.UserName)
			}
			if got.NiceNum != test.want.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.want.NiceNum)
			}
		})
	}
}
