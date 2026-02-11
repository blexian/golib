package comdatatype

import (
	"log"
	"os"
	"testing"
	"text/template"
	"time"
)

var tmpl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

var searchResultDemo = IssuesSearchResult{
	TotalCount: 2,
	Items: []*Issue{
		{
			Number:    1,
			User:      &User{Login: "user1"},
			Title:     "issue1",
			CreatedAt: time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Number:    2,
			User:      &User{Login: "user2"},
			Title:     "issue2",
			CreatedAt: time.Date(2024, 6, 2, 0, 0, 0, 0, time.UTC),
		},
	},
}

// !+daysAgo
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func TestTxtTemplate(t *testing.T) {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	result := searchResultDemo
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
