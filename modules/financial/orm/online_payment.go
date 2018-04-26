package orm

import (
	"encoding/json"
	"time"

	"fmt"

	"database/sql/driver"
	"strings"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
)

// VatPercent global payment vat percent
var VatPercent = config.RegisterInt("crab.modules.financial.vat", 9, "vat percent")

// OnlinePaymentStatus is bill payment model
// @Enum{
// }
type OnlinePaymentStatus string

const (
	// Init pay transaction
	Init OnlinePaymentStatus = "init"
	// BackToSite back to clickyab panel but not verified
	BackToSite OnlinePaymentStatus = "back_to_site"
	// Finalized verified and finalized payment successfully
	Finalized OnlinePaymentStatus = "finalized"
)

// BankReasonState is bank reason verify or redirect reason
// @Enum{
// }
type BankReasonState string

// Error return string
func (e BankReasonState) Error() string {
	return string(e)
}

// NullBankReason to make a nullable enum for InventoryState
type NullBankReason struct {
	Valid      bool
	BankReason BankReasonState
}

// MarshalJSON try to marshaling to json
func (nt NullBankReason) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.BankReason)
	}
	return []byte("null"), nil
}

// UnmarshalJSON try to unmarshal dae from input
func (nt *NullBankReason) UnmarshalJSON(b []byte) error {
	text := strings.ToLower(string(b))
	if text == "null" {
		nt.Valid = false
		nt.BankReason = BankReasonState("")
		return nil
	}

	err := json.Unmarshal(b, &nt.BankReason)
	if err != nil {
		return err
	}

	nt.Valid = true
	return nil
}

// Scan implements the Scanner interface.
func (nt *NullBankReason) Scan(value interface{}) error {
	var b []byte
	switch value.(type) {
	case []byte:
		b = value.([]byte)
	case string:
		b = []byte(value.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	nt.BankReason, nt.Valid = BankReasonState(b), BankReasonState(b) != ""
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullBankReason) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return string(nt.BankReason), nil
}

const (
	// VERIFYRequestData VERIFYRequestData
	VERIFYRequestData BankReasonState = "v_invalid_request_data"
	// VERIFYParams VERIFYParams
	VERIFYParams BankReasonState = "v_params_illegal_characters"
	// VERIFYMerchant VERIFYMerchant
	VERIFYMerchant BankReasonState = "v_merchant_auth_failed"
	// VERIFYTimeout VERIFYTimeout
	VERIFYTimeout BankReasonState = "v_timeout_reached"
	// VERIFYEmptyDigitalCode VERIFYEmptyDigitalCode
	VERIFYEmptyDigitalCode BankReasonState = "v_empty_digital_code"
	// VERIFYLongParamsLength VERIFYLongParamsLength
	VERIFYLongParamsLength BankReasonState = "v_params_too_long"
	// VERIFYIllegalCharsReturnAmount VERIFYIllegalCharsReturnAmount
	VERIFYIllegalCharsReturnAmount BankReasonState = "v_invalid_return_amount"
	// VERIFYDigitalCodeInvalid VERIFYDigitalCodeInvalid
	VERIFYDigitalCodeInvalid BankReasonState = "v_invalid_digital_code"
	// VERIFYLongParams VERIFYLongParams
	VERIFYLongParams BankReasonState = "v_params_too_short"
	// VERIFYNegativeReturnAmount VERIFYNegativeReturnAmount
	VERIFYNegativeReturnAmount BankReasonState = "v_negative_amount"
	// VERIFYReturnAmountMismatch VERIFYReturnAmountMismatch
	VERIFYReturnAmountMismatch BankReasonState = "v_amount_not_match"
	// VERIFYNotFoundTransaction VERIFYNotFoundTransaction
	VERIFYNotFoundTransaction BankReasonState = "v_transaction_not_found"
	// VERIFYFloatReturnAmount VERIFYFloatReturnAmount
	VERIFYFloatReturnAmount BankReasonState = "v_invalid_amount"
	// VERIFYInternalBankErr VERIFYInternalBankErr
	VERIFYInternalBankErr BankReasonState = "v_internal_bank"
	// VERIFYReturnSomeOfAmount VERIFYReturnSomeOfAmount
	VERIFYReturnSomeOfAmount BankReasonState = "v_multi_amount"
	// VERIFYIPPassReverseInvalid VERIFYIPPassReverseInvalid
	VERIFYIPPassReverseInvalid BankReasonState = "v_invalid_ip_or_pass"
	// VERIFYNotSupported VERIFYNotSupported
	VERIFYNotSupported BankReasonState = "v_not_supported"

	// PAYTransactionCancelled PAYTransactionCancelled
	PAYTransactionCancelled BankReasonState = "p_cancel_by_user"
	// PAYReturnMoreThanTransaction PAYReturnMoreThanTransaction
	PAYReturnMoreThanTransaction BankReasonState = "p_amount_not_match"
	// PAYEarlyVerifyCalled PAYEarlyVerifyCalled
	PAYEarlyVerifyCalled BankReasonState = "p_early_verify"
	// PAYWrongCardNumber PAYWrongCardNumber
	PAYWrongCardNumber BankReasonState = "p_invalid_card_num"
	// PAYCardNotFound PAYCardNotFound
	PAYCardNotFound BankReasonState = "p_invalid_card_owner"
	// PAYCardExpired PAYCardExpired
	PAYCardExpired BankReasonState = "p_card_expired"
	// PAYExceedWrongPassCard PAYExceedWrongPassCard
	PAYExceedWrongPassCard BankReasonState = "p_wrong_pass_3_times"
	// PAYPasswordCardWrong PAYPasswordCardWrong
	PAYPasswordCardWrong BankReasonState = "p_wrong_pass"
	// PAYAmountMoreThanValid PAYAmountMoreThanValid
	PAYAmountMoreThanValid BankReasonState = "p_amount_exceed"
	// PAYPanPinWrong PAYPanPinWrong
	PAYPanPinWrong BankReasonState = "p_pin_pan_error"
	// PAYResponseTimeout PAYResponseTimeout
	PAYResponseTimeout BankReasonState = "p_response_timeout"
	// PAYCCV2ExpDate PAYCCV2ExpDate
	PAYCCV2ExpDate BankReasonState = "p_invalid_cvv_or_expdate"
	// PAYSufficientFunds PAYSufficientFunds
	PAYSufficientFunds BankReasonState = "p_no_sufficient_funds"
	// PAYCardIssuerInvalid PAYCardIssuerInvalid
	PAYCardIssuerInvalid BankReasonState = "p_card_issuer_is_down"
	// PAYOtherBankErr PAYOtherBankErr
	PAYOtherBankErr BankReasonState = "p_bank_error"

	// PayNotSupported PayNotSupported
	PayNotSupported BankReasonState = "p_not_supported"

	// HashMismatchErr HashMismatchErr
	HashMismatchErr BankReasonState = "hash_mismatch"
	// MerchantMismatchErr MerchantMismatchErr
	MerchantMismatchErr BankReasonState = "merchant_mismatch"
	// RequestVerifyErr RequestVerifyErr
	RequestVerifyErr BankReasonState = "verify_request_failed"
	// RespVerifyErr RespVerifyErr
	RespVerifyErr BankReasonState = "verify_response_failed"
	// PriceMismatchErr PriceMismatchErr
	PriceMismatchErr BankReasonState = "amount_mismatch"
)

// OnlinePayment model in database
// @Model {
//		table = online_payments
//		primary = true, id
//		find_by = id
//		list = yes
// }
type OnlinePayment struct {
	ID          int64                  `json:"id" db:"id"`
	DomainID    int64                  `json:"domain_id" db:"domain_id"`
	UserID      int64                  `json:"user_id" db:"user_id"`
	GatewayID   int64                  `json:"gateway_id" db:"gateway_id"`
	Amount      int64                  `json:"amount" db:"amount"`
	Status      OnlinePaymentStatus    `json:"status" db:"status"`
	BankStatus  mysql.NullInt64        `json:"bank_status" db:"bank_status"`
	RefNum      mysql.NullString       `json:"ref_num" db:"ref_num"`
	ResNum      string                 `json:"res_num" db:"res_num"`
	CID         mysql.NullString       `json:"cid" db:"cid"`
	TraceNumber mysql.NullString       `json:"trace_number" db:"trace_number"`
	Attr        mysql.GenericJSONField `json:"attr" db:"attr"`
	ErrorReason NullBankReason         `json:"error_reason" db:"error_reason"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
}

// FindInitPaymentByResNum return the OnlinePayment base on its res num
func (m *Manager) FindInitPaymentByResNum(id string) (*OnlinePayment, error) {
	var res OnlinePayment
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE res_num=? AND status=?", getSelectFields(OnlinePaymentTableFull, ""), OnlinePaymentTableFull),
		id,
		Init,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ChargeUser charge user transaction
func (m *Manager) ChargeUser(payment *OnlinePayment, domainID, chargeAmount int64) error {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()
	err = m.UpdateOnlinePayment(payment)
	if err != nil {
		return err
	}

	billQ := fmt.Sprintf("SELECT COALESCE(SUM(amount),0) AS balance FROM %s WHERE user_id=?", BillingTableFull)
	oldBalance, err := m.GetRDbMap().SelectInt(billQ, payment.UserID)
	if err != nil {
		return err
	}

	newBalance := oldBalance + chargeAmount

	// create billing
	bill := &Billing{
		Amount:    chargeAmount,
		PayAmount: payment.Amount,
		VAT:       payment.Amount - chargeAmount,
		UserID:    payment.UserID,
		IncomeID:  payment.ID,
		DomainID:  domainID,
		PayModel:  OnlinePaymentModel,
		Balance:   newBalance,
	}
	err = m.CreateBilling(bill)
	if err != nil {
		return err
	}

	// update user balance
	userManager, err := aaa.NewAaaManagerFromTransaction(m.GetWDbMap())
	if err != nil {
		return err
	}

	userQ := fmt.Sprintf("UPDATE %s SET balance=? WHERE id=?", aaa.UserTableFull)
	_, err = userManager.GetWDbMap().Exec(userQ, newBalance, payment.UserID)
	return err
}
