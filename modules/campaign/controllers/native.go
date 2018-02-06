package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"strings"

	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"clickyab.com/crab/modules/upload/controllers"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/random"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// @Validate{
//}
type getNativeDataPayload struct {
	URL string `json:"url" validate:"required,url"`
}

// getNativeData getNativeData
// @Rest {
// 		url = /native/fetch
//		protected = true
// 		method = post
// }
func (c Controller) getNativeData(ctx context.Context, r *http.Request, p *getNativeDataPayload) (*getNativeDataResp, error) {
	u := authz.MustGetUser(ctx)
	res := getMetaTags(p.URL)
	if res == nil {
		return res, t9e.G("error fetching the link")
	}
	//upload if image exists
	if res.Image != "" {
		extension := strings.ToLower(filepath.Ext(res.Image))
		now := time.Now()
		fp := filepath.Join(controllers.UPath.String(), "uploads", "temp", now.Format("2006/01/02"))
		err := os.MkdirAll(fp, os.FileMode(controllers.Perm.Int64()))
		assert.Nil(err)
		fn := func() string {
			for {
				tmp := fmt.Sprintf("%d_%s%s", u.ID, <-random.ID, extension)
				if _, err := os.Stat(fp + tmp); os.IsNotExist(err) {
					return tmp
				}
			}
		}()
		f, err := os.OpenFile(filepath.Join(fp, fn), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(controllers.Perm.Int64()))
		assert.Nil(err)
		defer func() { assert.Nil(f.Close()) }()

		resp, err := http.Get(res.Image)
		if err != nil {
			return res, err
		}
		defer func() { _ = resp.Body.Close() }()

		_, err = io.Copy(f, resp.Body)
		assert.Nil(err)
		finalPath := filepath.Join("temp", now.Format("2006/01/02"), fn)
		res.Image = finalPath
	}

	if res.URL == "" {
		res.URL = p.URL
	}

	return res, nil
}

func getMetaTags(url string) *getNativeDataResp {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer func() { _ = resp.Body.Close() }()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var res = &getNativeDataResp{}
	z := html.NewTokenizer(strings.NewReader(string(bytes)))
bigLoop:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			break bigLoop
		case html.StartTagToken, html.SelfClosingTagToken, html.EndTagToken:
			name, hasAttr := z.TagName()
			if atom.Lookup(name) == atom.Body {
				break bigLoop
			}
			if atom.Lookup(name) != atom.Meta || !hasAttr {
				continue
			}
			m := make(map[string]string)
			var key, val []byte
			for hasAttr {
				key, val, hasAttr = z.TagAttr()
				m[atom.String(key)] = string(val)
			}
			processMeta(res, m)
		}
	}

	return res
}

type getNativeDataResp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	SiteName    string `json:"site_name"`
}

func processMeta(r *getNativeDataResp, attrs map[string]string) {
	switch attrs["property"] {
	case "og:description":
		r.Description = attrs["content"]
	case "og:title":
		r.Title = attrs["content"]
	case "og:url":
		r.URL = attrs["content"]
	case "og:image":
		r.Image = attrs["content"]
	case "og:site_name":
		r.SiteName = attrs["content"]
	}
}
