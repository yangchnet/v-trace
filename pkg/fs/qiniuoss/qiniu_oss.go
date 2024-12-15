package qiniuoss

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitee.com/qciip-icp/v-trace/pkg/fs"
	"github.com/google/wire"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiniuOSS struct {
	QiniuOSSConfig
}

type QiniuOSSConfig struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Domain    string `mapstructure:"domain"`
	Bucket    string `mapstructure:"bucket"`
}

var OssProvider = wire.NewSet(NewQiNiuOSS)

func NewQiNiuOSS(conf QiniuOSSConfig) *QiniuOSS {
	fmt.Println(conf)
	return &QiniuOSS{
		QiniuOSSConfig: conf,
	}
}

var _ fs.Interface = (*QiniuOSS)(nil)

// Store a file to fs, and return a key to find file
func (o *QiniuOSS) Store(fileContent []byte, key string, meta map[string]string, callbacks []fs.FsCallBack) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: o.Bucket,
	}
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	region, err := storage.GetRegion(o.AccessKey, o.Bucket)
	if err != nil {
		return "", err
	}

	cfg := storage.Config{
		Region:        region,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)

	resp := storage.PutRet{}
	dataLen := int64(len(fileContent))
	err = formUploader.Put(
		context.Background(),
		&resp,
		upToken,
		key,
		bytes.NewReader(fileContent),
		dataLen,
		&storage.PutExtra{
			Params: meta,
		},
	)
	if err != nil {
		return "", err
	}

	if len(callbacks) > 0 {
		for _, cb := range callbacks {
			cb(key)
		}
	}

	return resp.Key, nil
}

// Path return the file path that find by key
func (o *QiniuOSS) Path(key string) (string, error) {
	return fmt.Sprintf("http://%s", storage.MakePublicURL(o.Domain, key)), nil // 公有空间
}

// Read file from fs by path
func (o *QiniuOSS) Read(path string) ([]byte, error) {
	t := &http.Transport{}
	t.RegisterProtocol("tcp", http.NewFileTransport(http.Dir("/")))
	c := &http.Client{Transport: t}
	res, err := c.Get(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}
