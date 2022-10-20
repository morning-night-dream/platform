package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/morning-night-dream/platform/app/core/database/store"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type Slack struct {
	secret string
	client *http.Client
	store  *store.Article
}

func NewSlack(secret string, store *store.Article) *Slack {
	return &Slack{
		secret: secret,
		client: http.DefaultClient,
		store:  store,
	}
}

func (s *Slack) Events(w http.ResponseWriter, r *http.Request) {
	// @see https://github.com/slack-go/slack/blob/master/examples/eventsapi/events.go
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := s.verify(r.Header, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		s.challenge(w, body)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		switch ev := innerEvent.Data.(type) {
		// @see https://api.slack.com/events/link_shared
		case *slackevents.LinkSharedEvent:
			for _, link := range ev.Links {
				u, err := url.Parse(link.URL)
				if err != nil {
					return
				}

				s.save(r.Context(), *u)
			}
		default:
			log.Printf("%+v", ev)
		}
	}
}

func (s *Slack) verify(header http.Header, body []byte) error {
	sv, err := slack.NewSecretsVerifier(header, s.secret)
	if err != nil {
		return errors.Wrap(err, "failed new secrets verify")
	}

	if _, err := sv.Write(body); err != nil {
		return errors.Wrap(err, "failed write body")
	}

	if err := sv.Ensure(); err != nil {
		return errors.Wrap(err, "failed ensure")
	}

	return nil
}

func (s *Slack) challenge(w http.ResponseWriter, body []byte) {
	var r *slackevents.ChallengeResponse

	if err := json.Unmarshal(body, &r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text")

	_, _ = w.Write([]byte(r.Challenge))
}

func (s *Slack) save(ctx context.Context, u url.URL) {
	gr, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	res, err := s.client.Do(gr.WithContext(ctx))
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	og := opengraph.NewOpenGraph()

	err = og.ProcessHTML(strings.NewReader(string(body)))
	if err != nil {
		return
	}

	thumbnail := ""
	if len(og.Images) > 0 {
		thumbnail = og.Images[0].URL
	}

	article := model.Article{
		URL:         og.URL,
		Title:       og.Title,
		Thumbnail:   thumbnail,
		Description: og.Description,
	}

	err = s.store.Save(ctx, article)

	if err != nil {
		log.Print(err)

		return
	}
}
