package payment

import (
	"net/http"
	"net/url"

	"clickyab.com/crab/libs"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
)

// frontRouteRedirect frontend route after payment is done (success,fail)
var frontRouteRedirect = config.RegisterString("crab.modules.financial.payment.front.url", "financial/verify", "saman bank name")

// InitPaymentResp init payment response
type InitPaymentResp struct {
	Params  map[string]interface{} `json:"params"`
	Method  string                 `json:"method"`
	BankURL string                 `json:"bank_url"`
}

// CommonPay for similar actions across payment gates
type CommonPay struct {
}

// FrontRedirect redirect to front route
func (s *CommonPay) FrontRedirect(w http.ResponseWriter, r *http.Request, code int, v url.Values) error {
	u := &url.URL{
		Scheme: framework.Scheme(r),
		Host:   r.Host,
		Path:   frontRouteRedirect.String(),
	}
	u.RawQuery = v.Encode()
	return libs.Redirect(w, code, u)
}

// Payable handle payment
type Payable interface {
	InitPayment(r *http.Request) *InitPaymentResp
	VerifyPayment(string, string, string) error
	FrontRedirect(http.ResponseWriter, *http.Request, int, url.Values) error

	MID() int64
	UserID() int64
	RedirectURL(r *http.Request) string
	Amount() int64
	ResNum() string
	VerifyURL() string
	BankName() string
	VerifyErr(int64) error
	PaymentErr(int64) error
	GetPaymentURL() string
	GetPaymentMethod() string
}
