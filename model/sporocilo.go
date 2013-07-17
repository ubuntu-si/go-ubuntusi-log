// rest
package sporocilo

import (
	"fmt"
	"github.com/coocood/qbs"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type Sporocilo struct {
	Id      int64
	Nick    string
	Msg     string
	Title   string
	Created time.Time
}

func ShraniSporocilo(nick string, title string, msg string) (error, *Sporocilo) {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	sp := new(Sporocilo)
	sp.Nick = nick
	sp.Title = title
	sp.Msg = msg
	_, err2 := q.Save(sp)
	return err2, sp
}

func NajdiSporocila(offset int) (error, []*Sporocilo) {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	var sporocila []*Sporocilo
	err2 := q.Limit(25).Offset(offset).FindAll(&sporocila)
	return err2, sporocila
}

func UstvariTabeloSporocila() error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}
	defer migration.Close()
	return migration.CreateTableIfNotExists(new(Sporocilo))
}

func RegisterDb() {
	err := os.MkdirAll("./data", 777)
	if err != nil {
		fmt.Println(err)

	}
	qbs.Register("sqlite3", "./data/ubuntu_log.sqlite", "./data/ubuntu_log.sqlite", qbs.NewSqlite3())
}
