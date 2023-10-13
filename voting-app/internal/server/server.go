package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/thyagofr/voting-app/internal/pubsub"
	"html/template"
	"net/http"
)

type HTTPServer struct {
	pub pubsub.Publisher
	tpl *template.Template
	id  string
}

func NewHTTPServer(pub pubsub.Publisher, htmlFolder string) *HTTPServer {
	return &HTTPServer{
		pub: pub,
		tpl: template.Must(
			template.ParseGlob(fmt.Sprintf("%s/*.gohtml", htmlFolder)),
		),
		id: uuid.NewString(),
	}
}

func (srv HTTPServer) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hostname": srv.id,
	}
	switch r.Method {
	case http.MethodPost:
		var (
			ctx  = context.Background()
			vote = r.FormValue("vote")
		)

		if err := srv.pub.Publish(ctx, pubsub.NewVote(vote)); err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		}
		_ = srv.tpl.ExecuteTemplate(w, "index.gohtml", data)
	case http.MethodGet:
		if err := srv.tpl.ExecuteTemplate(w, "index.gohtml", data); err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		}
	}
}
