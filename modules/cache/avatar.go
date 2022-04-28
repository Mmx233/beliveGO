package cache

import (
	"encoding/json"
	"fmt"
	"github.com/Mmx233/tool"
	"os"
	"path"
	"time"
)

type avatarCache struct {
	URL     string
	Expires int64
}

type avatar struct {
	path func(uid uint) string
}

var Avatar = avatar{
	path: func(uid uint) string {
		return path.Join(os.TempDir(), "belive-avatar-"+fmt.Sprint(uid))
	},
}

func (a *avatar) Cache(uid uint, url string, valid time.Duration) error {
	d, _ := json.Marshal(&avatarCache{
		URL:     url,
		Expires: time.Now().Add(valid).Unix(),
	})
	f, e := os.OpenFile(a.path(uid), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = f.Write(d)
	return e
}

func (a *avatar) Read(uid uint) (string, error) {
	if !tool.File.Exists(a.path(uid)) {
		return "", Nil
	}
	f, e := os.OpenFile(a.path(uid), os.O_RDONLY, 0600)
	if e != nil {
		return "", e
	}
	defer f.Close()
	var data avatarCache
	e = json.NewDecoder(f).Decode(&data)
	if e != nil {
		return "", e
	}
	if data.Expires < time.Now().Unix() {
		return "", Nil
	}
	return data.URL, nil
}
