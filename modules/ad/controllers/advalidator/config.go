package advalidator

import (
	"encoding/json"
	"sync"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

var (
	adValidation = config.RegisterString("", `{"app_native":{"logo":{"rate":1,"min":{"width":627,"height":627},"ext":["image/png"]},"icon":{"rate":1,"min":{"width":512,"height":512},"ext":["image/png"]},"v_image":{"exact":[{"width":320,"height":480}],"ext":["image/png","image/gif","image/jpeg"],"size":200},"h_image":{"rate":1.5,"exact":[{"width":480,"height":320}],"ext":["image/png","image/gif","image/jpeg"]},"video":{"ext":["video/mp4"],"size":20000}},"web_native":{"logo":{"rate":1,"min":{"width":627,"height":627},"ext":["image/png"]},"icon":{"rate":1,"min":{"width":512,"height":512},"ext":["image/png"]},"v_image":{"exact":[{"width":320,"height":480}],"ext":["image/png","image/gif","image/jpeg"],"size":200},"h_image":{"rate":1.5,"exact":[{"width":480,"height":320}],"ext":["image/png","image/gif","image/jpeg"]},"video":{"ext":["video/mp4"],"size":20000}}}`, "")
	lock         = sync.Mutex{}
	// AdValidationConf AdValidationConf
	AdValidationConf = AdTotalValidationRule{}
)

// RegisterAdValidationRules RegisterAdValidationRules
func RegisterAdValidationRules() {
	lock.Lock()
	defer lock.Unlock()
	assert.Nil(json.Unmarshal([]byte(adValidation.String()), &AdValidationConf))
}
