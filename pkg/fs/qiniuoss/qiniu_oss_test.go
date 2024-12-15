package qiniuoss

import (
	"github.com/google/uuid"
	"os"
	"testing"
)

func Test_qiniu(t *testing.T) {
	key := uuid.New().String()

	o := NewQiNiuOSS(QiniuOSSConfig{
		AccessKey: "UdafhaRtgYAS5DfUGXfK6j4x_AtucoskBiW8whki",
		SecretKey: "OdPkAd5dzCVO4-2XfGm_N6n92rSZH30QJwONHXOm",
		Domain:    "rqqg8wbdg.hd-bkt.clouddn.com",
		Bucket:    "vtrace-oss",
	})

	content, err := os.ReadFile("/home/jaylan/go-project/gitee.com/qciip/v-trace/LICENSE")
	if err != nil {
		t.Fatal(err)
	}

	ossKey, err := o.Store(content, key, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("---> %s\n", ossKey)

	url, err := o.Path(ossKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("---> %s\n", url)

	b, err := o.Read(url)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile("a.png", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	_, err = f.Write(b)
	if err != nil {
		t.Fatal(err)
	}
}
