package main

import (
	"os"

	"encoding/json"

	"crypto/sha1"

	"fmt"

	"path/filepath"

	"github.com/clickyab/services/assert"
	"github.com/ogier/pflag"
)

var (
	path   = pflag.String("path", "index.json", "the json to split")
	target = pflag.String("target", "target", "target folder WARNING all files inside it is removed")
)

func main() {
	pflag.Parse()
	assert.Nil(os.MkdirAll(*target, 0777))
	in, err := os.OpenFile(*path, os.O_RDONLY, 0666)
	assert.Nil(err)
	defer func() {
		_ = in.Close()
	}()
	dec := json.NewDecoder(in)
	var data = make(map[string]interface{})
	assert.Nil(dec.Decode(&data))
	res, err := split("root", data, 0)
	assert.Nil(err)
	assert.Nil(save(filepath.Join(*target, "index.json"), res))
}

func mkFileName(in string) string {
	h := sha1.New()
	c := []byte(in)
	_, _ = h.Write(c)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func split(prefix string, data map[string]interface{}, depth int) (map[string]interface{}, error) {
	var err error
	for i := range data {

		if d, ok := data[i].(map[string]interface{}); ok {
			if depth < 1 {
				data[i], err = split(prefix+"_"+i, d, depth+1)
				if err != nil {
					return nil, err
				}
				continue
			}
			path := mkFileName(prefix + "_" + i)
			if err = save(filepath.Join(*target, path), data[i]); err != nil {
				return nil, err
			}
			data[i] = "file://" + path
			continue
		}
	}

	return data, nil
}

func save(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(data)
}
