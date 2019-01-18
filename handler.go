package zendeskintone

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	aelog "google.golang.org/appengine/log"

	"github.com/yoheimiyamoto/gcp/taskqueue"
)

func init() {
	err := godotenv.Load("kintone.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	taskqueue.HandleFuncs("/", "default", Handler)
	appengine.Main()
}

/*
zendeskからのwebhookを受け取ってKintoneにPostします。
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// handling
	err := handling(ctx, r)
	if err != nil {
		aelog.Errorf(ctx, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handling(ctx context.Context, r *http.Request) error {

	// body取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	aelog.Infof(ctx, "body: %s", string(body))

	// work
	return work(ctx, body)
}

func work(ctx context.Context, body []byte) error {

	// レコード作成
	rec, err := newRecord(body)
	if err != nil {
		return errors.Wrap(err, "レコード作成")
	}
	aelog.Infof(ctx, "record: %#v", rec.Fields)

	appID := fmt.Sprint(rec.Fields["kintone_app_id"])
	if appID == "" {
		return fmt.Errorf("body内にpost先のkintone_app_idが存在しません")
	}

	// kintoneへpost
	c := newClient(ctx, appID)
	_, err = c.AddRecord(rec)
	if err != nil {
		return errors.Wrap(err, "KintoneへのPost")
	}

	return nil
}
