// rest
package rest

import (
	"encoding/json"
	"github.com/coocood/jas"
	"gorsi/model"
	"log"
)

type Sporocilo struct{}

func (*Sporocilo) Post(ctx *jas.Context) {
	nick := ctx.RequireString("nick")
	msg := ctx.RequireString("msg")
	title := ctx.RequireString("title")
	if nick != "" && msg != "" {
		err, res := sporocilo.ShraniSporocilo(nick, title, msg)
		if err == nil {
			b, _ := json.Marshal(res)
			log.Println(string(b))
			Notify(string(b))
			ctx.Data = res
		}
	}
}

type Sporocila struct{}

func (*Sporocila) Last(ctx *jas.Context) { // `GET /sporocila/:offset`
	err, res := sporocilo.NajdiSporocila(0)
	if err == nil {
		ctx.Data = res
	}
}
func (*Sporocila) Get(ctx *jas.Context) { // `GET /sporocila/:offset`
	offset := ctx.RequireInt("offset")
	err, res := sporocilo.NajdiSporocila(int(offset))
	if err == nil {
		ctx.Data = res
	}
}
