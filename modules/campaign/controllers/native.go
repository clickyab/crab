package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"strings"

	"io"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// @Validate{
//}
type getNativeDataPayload struct {
	URL string `json:"url" validate:"required,gt=5,url"`
}

// getNativeData getNativeData
// @Route {
// 		url = /native/fetch
//		method = post
//		payload = getNativeDataPayload
//		middleware = authz.Authenticate
//		200 = getNativeDataResp
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (c Controller) getNativeData(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	p := c.MustGetPayload(ctx).(*getNativeDataPayload)
	link := p.URL
	resp, err := http.Get(link)
	if err != nil {
		logrus.Warn(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warn(err)
	}
	resp.Body.Close()
	res := getMetaTags(strings.NewReader(string(bytes)))
	c.OKResponse(w, res)
}

func getMetaTags(reader io.Reader) *getNativeDataResp {
	var res = &getNativeDataResp{}
	z := html.NewTokenizer(reader)
bigLoop:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				break bigLoop
			}
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
			res = processMeta(res, m)
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

func processMeta(r *getNativeDataResp, attrs map[string]string) *getNativeDataResp {
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
	return r
}
