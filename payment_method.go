package yookassa

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	PaymentByBankCard = "bank_card"
	PaymentBySberbank = "sberbank"
	PaymentByYooMoney = "yoo_money"
	PaymentByQIWI     = "qiwi"
	PaymentByTinkoff  = "tinkoff_bank"
	PaymentBySBP      = "sbp"
	PaymentBySberLoan = "sber_loan"
	PaymentByInstallments = "installments"
	PaymentBySberBusinessOnline = "b2b_sberbank"
	PaymentByMobileBalance = "mobile_balance"
	PaymentByCash = "cash"
	VATUntaxed    = "untaxed"
	VATCalculated = "calculated"
	VATMixed      = "mixed"
)

type SimplePaymentRequestData struct {
	Type string `json:"type"`
}

type PhoneNumberPaymentRequestData struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
}

type BankCardPaymentRequestData struct {
	Card BankCardData
}

type bankCardPaymentRequestData struct {
	Type string       `json:"type"`
	Card BankCardData `json:"card"`
}

func (p *BankCardPaymentRequestData) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(&bankCardPaymentRequestData{
		Type: PaymentByBankCard,
		Card: p.Card,
	})
}

type BankCardData struct {
	Number      string `json:"number"`
	ExpiryYear  string `json:"expiry_year"`
	ExpiryMonth string `json:"expiry_month"`
	CSC         string `json:"csc,omitempty"`
	Cardholder  string `json:"cardholder,omitempty"`
}

type SberBusinessOnlinePaymentRequestDate struct {
	PaymentPurpose string
	VATData        VATData
}

type sberBusinessOnlinePaymentRequestDate struct {
	Type           string  `json:"type"`
	PaymentPurpose string  `json:"payment_purpose"`
	VATData        VATData `json:"vat_data"`
}

func (p *SberBusinessOnlinePaymentRequestDate) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(&sberBusinessOnlinePaymentRequestDate{
		Type:           PaymentBySberBusinessOnline,
		PaymentPurpose: p.PaymentPurpose,
		VATData:        p.VATData,
	})
}

// Response

type PaymentMethod struct {
	ID      string      `json:"id,omitempty"`
	Type    string      `json:"type"`
	Saved   bool        `json:"saved"`
	Details interface{} `json:"details"`
}

type paymentMethod struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Saved bool   `json:"saved"`
}

func (m *PaymentMethod) UnmarshalJSON(data []byte) error {
	var mm paymentMethod
	if err := jsoniter.Unmarshal(data, &mm); err != nil {
		return err
	}

	m.ID = mm.ID
	m.Type = mm.Type
	m.Saved = mm.Saved

	switch m.Type {
	case PaymentByBankCard:
		m.Details = &BankCardPaymentDetails{}
	case PaymentByYooMoney:
		m.Details = &YooMoneyPaymentDetails{}
	case PaymentBySberbank:
		m.Details = &SberOnlinePaymentDetails{}
	case PaymentBySberBusinessOnline:
		m.Details = &SberBusinessOnlinePaymentDetails{}
	case PaymentByQIWI,
		PaymentByTinkoff, 
		PaymentBySBP,
		PaymentByMobileBalance, 
		PaymentByCash, PaymentByInstallments:
		m.Details = nil
	default:
		m.Details = &map[string]interface{}{}
	}

	if m.Details != nil {
		return jsoniter.Unmarshal(data, m.Details)
	}
	return nil
}

type BankCardPaymentDetails struct {
	Card BankCardInfo `json:"card"`
}

type BankCardInfo struct {
	First6        string `json:"first_6,omitempty"`
	Last4         string `json:"last_4"`
	ExpiryYear    string `json:"expiry_year"`
	ExpiryMonth   string `json:"expiry_month"`
	Type          string `json:"type"`
	IssuerCountry string `json:"issuer_country"`
	IssuerName    string `json:"issuer_name"`
	Source        string `json:"source"`
}

type YooMoneyPaymentDetails struct {
	AccountNumber string `json:"account_number"`
}

type SberOnlinePaymentDetails struct {
	Phone string `json:"phone"`
}

type SberBusinessOnlinePaymentDetails struct {
	PaymentPurpose   string           `json:"payment_purpose"` // Max 210 symbols
	PayerBankDetails PayerBankDetails `json:"payer_bank_details"`
}

type PayerBankDetails struct {
	FullName   string `json:"full_name"`
	ShortName  string `json:"short_name"`
	Address    string `json:"address"`
	INN        string `json:"inn"`
	KPP        string `json:"kpp"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	BankBIK    string `json:"bank_bik"`
	Account    string `json:"account"`
}

type VATData struct {
	Type   string `json:"type"`
	Amount Amount `json:"amount"`
	Rate   string `json:"rate"`
}