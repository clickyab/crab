package advalidator

import (
	"encoding/json"
	"sync"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

var (
	adValidation = config.RegisterString("", `{
  "web_vast":{"image":{"exact":[{"width":800,"height":440}],"ext":["image/png","image/gif","image/jpeg"]},"video":{"ext":["video/mp4"],"exact":[{"width":800,"height":440}],"size":2000}},
  "app_native": {
    "logo": {
      "rate": 1,
      "min": {
        "width": 627,
        "height": 627
      },
      "ext": [
        "image/png"
      ]
    },
    "icon": {
      "rate": 1,
      "min": {
        "width": 512,
        "height": 512
      },
      "ext": [
        "image/png"
      ]
    },
    "v_image": {
      "exact": [
        {
          "width": 320,
          "height": 480
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ],
      "size": 200
    },
    "h_image": {
      "rate": 1.5,
      "exact": [
        {
          "width": 480,
          "height": 320
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ]
    },
    "video": {
      "ext": [
        "video/mp4"
      ],
      "size": 20000
    }
  },
  "web_native": {
    "logo": {
      "rate": 1,
      "min": {
        "width": 627,
        "height": 627
      },
      "ext": [
        "image/png"
      ]
    },
    "icon": {
      "rate": 1,
      "min": {
        "width": 512,
        "height": 512
      },
      "ext": [
        "image/png"
      ]
    },
    "v_image": {
      "exact": [
        {
          "width": 320,
          "height": 480
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ],
      "size": 200
    },
    "h_image": {
      "rate": 1.5,
      "exact": [
        {
          "width": 480,
          "height": 320
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ]
    },
    "video": {
      "ext": [
        "video/mp4"
      ],
      "size": 20000
    }
  },
  "web_banner": {
    "video": {
      "ext": [
        "video/mp4"
      ],
      "size": 2000
    },
    "image": {
      "exact": [
        {
          "width": 120,
          "height": 600
        },
        {
          "width": 160,
          "height": 600
        },
        {
          "width": 300,
          "height": 250
        },
        {
          "width": 336,
          "height": 280
        },
        {
          "width": 468,
          "height": 60
        },
        {
          "width": 728,
          "height": 90
        },
        {
          "width": 120,
          "height": 240
        },
        {
          "width": 320,
          "height": 50
        },
        {
          "width": 800,
          "height": 440
        },
        {
          "width": 300,
          "height": 600
        },
        {
          "width": 970,
          "height": 90
        },
        {
          "width": 970,
          "height": 250
        },
        {
          "width": 250,
          "height": 250
        },
        {
          "width": 300,
          "height": 1050
        },
        {
          "width": 320,
          "height": 480
        },
        {
          "width": 480,
          "height": 320
        },
        {
          "width": 128,
          "height": 128
        }
      ],
      "ext": [
        "image/gif",
        "image/jpeg",
        "image/png"
      ]
    },
    "video_image": {
      "exact": [
        {
          "width": 640,
          "height": 360
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ],
      "size": 200
    }
  },
  "app_banner": {
    "video": {
      "ext": [
        "video/mp4"
      ],
      "size": 2000
    },
    "image": {
      "exact": [
        {
          "width": 120,
          "height": 600
        },
        {
          "width": 160,
          "height": 600
        },
        {
          "width": 300,
          "height": 250
        },
        {
          "width": 336,
          "height": 280
        },
        {
          "width": 468,
          "height": 60
        },
        {
          "width": 728,
          "height": 90
        },
        {
          "width": 120,
          "height": 240
        },
        {
          "width": 320,
          "height": 50
        },
        {
          "width": 800,
          "height": 440
        },
        {
          "width": 300,
          "height": 600
        },
        {
          "width": 970,
          "height": 90
        },
        {
          "width": 970,
          "height": 250
        },
        {
          "width": 250,
          "height": 250
        },
        {
          "width": 300,
          "height": 1050
        },
        {
          "width": 320,
          "height": 480
        },
        {
          "width": 480,
          "height": 320
        },
        {
          "width": 128,
          "height": 128
        }
      ],
      "ext": [
        "image/gif",
        "image/jpeg",
        "image/png"
      ]
    },
    "video_image": {
      "exact": [
        {
          "width": 640,
          "height": 360
        }
      ],
      "ext": [
        "image/png",
        "image/gif",
        "image/jpeg"
      ],
      "size": 200
    }
  }
}`, "")
	lock = sync.Mutex{}
	// AdValidationConf AdValidationConf
	AdValidationConf = AdTotalValidationRule{}
)

// RegisterAdValidationRules RegisterAdValidationRules
func RegisterAdValidationRules() {
	lock.Lock()
	defer lock.Unlock()
	assert.Nil(json.Unmarshal([]byte(adValidation.String()), &AdValidationConf))
}
