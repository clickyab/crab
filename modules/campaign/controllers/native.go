package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"strings"

	"errors"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// @Validate{
//}
type getNativeDataPayload struct {
	URL string `json:"url" validate:"required,url"`
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
	res := getMetaTags(p.URL)
	if res == nil {
		c.BadResponse(w, errors.New("error fetching the link"))
	}
	c.OKResponse(w, res)
}

func getMetaTags(url string) *getNativeDataResp {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
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
