package zendeskintone

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/yoheimiyamoto/kintone-sdk-go/kintone"
	"google.golang.org/appengine/urlfetch"
)

type rawRecord map[string]interface{}

func newRecord(body []byte) (*kintone.Record, error) {
	// rawレコード作成
	raw := make(rawRecord)
	err := json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	// kintone record 作成
	f := kintone.Fields{}
	for k, v := range raw {
		f[k] = kintone.SingleLineTextField(fmt.Sprint(v))
	}
	return kintone.NewRecord(f), nil
}

func newClient(ctx context.Context, appID string) *kintone.Client {
	var httpclient *http.Client
	if ctx != nil {
		httpclient = urlfetch.Client(ctx)
	}
	client, _ := kintone.NewClient(
		os.Getenv("KINTONE_DOMAIN"),
		os.Getenv("KINTONE_USER"),
		os.Getenv("KINTONE_PASSWORD"),
		appID,
		httpclient,
	)
	return client
}
