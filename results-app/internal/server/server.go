package server

import (
	"fmt"
	"github.com/thyagofr/results-app/internal/repository"
	"html/template"
	"log/slog"
	"net/http"
)

type HTTPServer struct {
	repo   repository.VoteRepository
	logger *slog.Logger
	tpl    *template.Template
}

func NewHTTPServer(repo repository.VoteRepository, htmlFolder string) *HTTPServer {
	return &HTTPServer{
		repo: repo,
		tpl: template.Must(
			template.ParseGlob(fmt.Sprintf("%s/*.gohtml", htmlFolder)),
		),
		logger: slog.Default(),
	}
}

func (srv HTTPServer) Index(w http.ResponseWriter, r *http.Request) {
	report, err := srv.repo.GetReport(r.Context())
	if err != nil {
		srv.logger.Error(err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	_ = srv.tpl.ExecuteTemplate(w, "index.gohtml", report)
}
