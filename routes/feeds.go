// feeds.go
package rest

import (
	"fmt"
	"github.com/gorilla/feeds"
	"gorsi/model"
	"net/http"
	"strconv"
	"time"
)

func Atom(w http.ResponseWriter, r *http.Request) {

	feed := &feeds.Feed{
		Title:       "Slo Prevajalci",
		Link:        &feeds.Link{Href: "http://ubuntu.si/slo_prevajalci"},
		Description: "Sporočila iz #ubuntu-si",
		Author:      &feeds.Author{"dz0ny", "dz0ny@freenode"},
		Created:     time.Now(),
	}
	_, res := sporocilo.NajdiSporocila(0)

	for _, v := range res {
		feed.Add(&feeds.Item{
			Title:       v.Nick + ": " + v.Title,
			Link:        &feeds.Link{Href: "http://ubuntu.si/slo_prevajalci/sporocilo/" + strconv.FormatInt(v.Id, 36)},
			Description: v.Msg,
			Created:     v.Created,
		})
	}

	atom, _ := feed.ToAtom()
	fmt.Fprintf(w, atom)
}
