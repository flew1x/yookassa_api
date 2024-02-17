package yookassa

import "time"

const (
	StatusPending  = "pending"
	StatusWaiting  = "waiting_for_capture"
	SatusSucceeded = "succeeded"
	StatusCanceled = "canceled"
)

type PaymentRequest struct {
	Amount            Amount                 `json:"amount"` // Sum of the payment. Sometimes the sum of the payment is not equal to the sum of the order because of the commission.
	Description       string                 `json:"description,omitempty"` // A description of the payment. No longer than 128 characters. As example "Payment for order #1234"
	Receipt           *ReceiptRequestData    `json:"receipt,omitempty"` // Data for formatting the receipt
	Recipient         *RecipientRequest      `json:"recipient,omitempty"` // Consumer of the payment
	PaymentToken      string                 `json:"payment_token,omitempty"` // Payment token. Formated by Checkout.js or mobile sdk
	PaymentMethodID   string                 `json:"payment_method_id,omitempty"` // Payment method ID (saved method)
	PaymentMethodData *PaymentMethod         `json:"payment_method_data,omitempty"` // Payment method (you can dont input it and user can choose it from the list)
	Confirmation      RedirectConfirmation   `json:"confirmation"` // Scenario of payment confirmation
	SavePaymentMethod bool                   `json:"save_payment_method,omitempty"` // Bool to save payment method
	Capture           bool                   `json:"capture,omitempty"` // Automatically capture the payment
	ClientIP          string                 `json:"client_ip,omitempty"` // Client IP  if dont input heretofore in the request (tcp connect)
	Metadata          map[string]interface{} `json:"metadata,omitempty"` // Anything you want to store and help you
	AirlineTicket     *AirlineTicketData     `json:"airline,omitempty"` // For airline tickets
	Transfers         []TransferRequestData  `json:"transfers,omitempty"` // Data of transfers money
	Deal              *DealRequestData       `json:"deal,omitempty"` // Data of deal
	FraudData         *FraudData             `json:"fraud_data,omitempty"` // Fraud data for check fraud
	MerchantCustomerID string                 `json:"merchant_customer_id,omitempty"` // ID of customer as example email or phone. No more than 200 characters
}

type PaymentResponse struct {
	ID                   string                `json:"id"`
	Status               string                `json:"status"`
	Test                 bool                  `json:"test"`
	Paid                 bool                  `json:"paid"`
	Refundable           bool                  `json:"refundable"`
	Amount               Amount                `json:"amount"`
	IncomeAmount         *Amount               `json:"income_amount,omitempty"`
	RefundedAmount       *Amount               `json:"refunded_amount,omitempty"`
	Created              time.Time             `json:"created_at"`
	Captured             time.Time             `json:"captured_at"`
	Expires              time.Time             `json:"expires_at"`
	Description          string                `json:"description"`
	Recipient            Recipient             `json:"recipient"`
	PaymentMethod        PaymentMethod         `json:"payment_method"`
	ReceiptRegistration  string                `json:"receipt_registration,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
	Confirmation         ConfirmationInfo      `json:"confirmation"`
	Transfers            []TransferDetails     `json:"transfers,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type FraudData struct {
	ToppedUpPhone string `json:"topped_up_phone,omitempty"` // Topped up phone
	MerchantCustomerBankAccount string `json:"merchant_customer_bank_account,omitempty"` // Merchant customer bank account
}

type DealRequestData struct {
	ID string `json:"id,omitempty"` // ID of the deal
	Settlements 
}

type Settlements struct {
	Type string `json:"type,omitempty"` // Type of settlement. As example "payout"
	Amount Amount `json:"amount,omitempty"` // Amount of the settlement
}

// List of items. No more than 100 items.
type Items struct {
	Description string `json:"description"` // Name of the item (1-128 characters)
	Amount      Amount `json:"amount"` // Price of the item
	VatCode     int `json:"vat_code,omitempty"` // NDS code (1-6) see https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#vat-codes
	Quantity    string `json:"quantity"` // Quantity of the item. Float or integer number
	Measure     string `json:"measure,omitempty"` // Unit of measurement. As example "piece". Required parameter for Yookassa receipts
	MarkQuantity MarkQuantity `json:"mark_quantity,omitempty"` // Mark quantity. As example "piece". Required parameter for Yookassa receipts
	PaymentSubject string `json:"payment_subject,omitempty"` // Payment subject. As example "service or commodity" https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#payment-subject
	PaymentMode string `json:"payment_mode,omitempty"` // Payment mode. As example "full_payment" https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#payment-mode
	CountryOfOriginCode string `json:"country_of_origin_code,omitempty"` // Country of origin code. As example "RU" https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#country-of-origin
	CustomsDeclarationNumber string `json:"customs_declaration_number,omitempty"` // Customs declaration number (1-32). As example "10714040/140917/0090376"
	Excise string `json:"excise,omitempty"` // Sum of excise. As example "10.00"
	ProductCode string `json:"product_code,omitempty"` // Unique number which appropriate for the product. Format: a number with 16th formation. Max length: 32bytes.  00 00 00 01 00 21 FA 41 00 23 05 41 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 12 00 AB 00
	MarkCodeInfo MarkCodeInfo `json:"mark_code_info,omitempty"` // Mark code info. Must be filled one of fields
	MarkMode string `json:"mark_mode,omitempty"` // Mode of code processing. Must be filled "0"
	PaymentSubjectIndustryDetails PaymentSubjectIndustryDetails `json:"payment_subject_industry_details,omitempty"` // Industry calculation details
}

type PaymentSubjectIndustryDetails struct {
	FederalID string `json:"federal_id,omitempty"` // Federal`s ID
	DocumentDate string `json:"document_date,omitempty"` // The date of establishment
	DocumentNumber string `json:"document_number,omitempty"` // The number of normalized certificate
	Value string `json:"value,omitempty"` // The value of the certificate
}

type MarkCodeInfo struct {
	MarkCodeRaw string `json:"mark_code_raw,omitempty"` // Code of commodity with view which was parsed of scanner. As example "010460406000590021N4N57RTCBUZTQ\u001d2403054002410161218\u001d1424010191ffd0\u001g92tIAF/YVpU4roQS3M/m4z78yFq0nc/WsSmLeX6QkF/YVWwy5IMYAeiQ91Xa2m/fFSJcOkb2N+uUUtfr4n0mOX0Q=="
	Unknown string `json:"unknown,omitempty"` // Unknown code of commodity
	Ean8 string `json:"ean_8,omitempty"` // EAN8 code
	Ean13 string `json:"ean_13,omitempty"` // EAN13 code
	Itf14 string `json:"itf_14,omitempty"` // ITF14 code
	Gs10 string `json:"gs_10,omitempty"` // GS10 code
	Gs1m string `json:"gs_1m,omitempty"` // GS1M code
	Short string `json:"short,omitempty"` // Code of commodity in format - short code of mark
	Fur string `json:"fur,omitempty"` // Control-identifier sign of fur commodity
	Egails20 string `json:"egails_20,omitempty"` // Egails20 code
	Egails30 string `json:"egails_30,omitempty"` // Egails30 code
}

type MarkQuantity struct {
	Numerator string `json:"numerator"` // Count of sales of the one recived box. Required parameter for Yookassa receipts. Mustnt be bigger than denominator
	Denominator string `json:"denominator"` // All count of sales of the items. Required parameter for Yookassa receipts
}

// Sum of the payment. Sometimes the sum of the payment is not equal to the sum of the order because of the commission.
type Amount struct {
	Value string `json:"value"` //The sum of the chosen payment method. Always floating point number. As example "10.00".
	Currency string `json:"currency"` // The code of the currency. As example "RUB". 3 letter code  
}

// Information about the reciept
type RecipientRequest struct {
	GatewayID string `json:"gateway_id"` // ID of subaccount. Use for separate payments from one account. It is not required.
}

type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

type AuthorizationDetails struct {
	RRN      string `json:"rrn"`
	AuthCode string `json:"auth_code"`
}

type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

type TransferDetails struct {
	AccountID         string `json:"account_id"`
	Status            string `json:"status"`
	Amount            Amount `json:"amount"`
	PlatformFeeAmount Amount `json:"platform_fee_amount"`
}


// Data for formatting the receipt
type ReceiptRequestData struct {
	Customer Customer `json:"customer,omitempty"` // Information about the customer. At least an email
	Items    []Items      `json:"items,omitempty"`    // List of items. No more than 100 items.
	Phone    string       `json:"phone,omitempty"`    // Phone of the customer for send a receipt. It is obsolete - recommend input to receipt.customer.phone
	Email    string       `json:"email,omitempty"`    // Email of the customer for send a receipt. It is obsolete - recommend input to receipt.customer.email
	TaxSystemCode string `json:"tax_system_code,omitempty"` // Tax system code (1-6). For Yookassa it is not required
	ReceiptIndustryDetails ReceiptIndustryDetails `json:"receipt_industry_details,omitempty"` // Industry calculation details
	ReceiptOperatorDetails ReceiptOperatorDetails `json:"receipt_operator_details,omitempty"` // Operator calculation details
}

type ReceiptOperatorDetails struct {
	OperationID string `json:"operation_id,omitempty"` // ID of the operation (0-255)
	Value string `json:"value,omitempty"` // The data of the operation
	CreatedAt string `json:"created_at,omitempty"` // The date of operation
}

type ReceiptIndustryDetails struct {
	FederalID string `json:"federal_id,omitempty"` // Federal`s ID
	DocumentDate string `json:"document_date,omitempty"` // The date of establishment
	DocumentNumber string `json:"document_number,omitempty"` // The number of normalized certificate
	Value string `json:"value,omitempty"` // The value of the certificate
}

// Information about the customer. At least an email
type Customer struct {
	FullName string `json:"full_name,omitempty"` // Full name of the customer
	INN      string `json:"inn,omitempty"` // INN of the customer. If does not exist, passport data is required
	Email    string `json:"email,omitempty"` // Email of the customer. The required field for the payment if you dont use phone
	Phone    string `json:"phone,omitempty"` // Phone of the customer. The required field for the payment if you dont use email
}

type TransferRequestData struct {
	AccountID         string `json:"account_id"`
	Amount            Amount `json:"amount"`
	PlatformFeeAmount Amount `json:"platform_fee_amount"`
}

type AirlineTicketData struct {
	TicketNumber     string                       `json:"ticket_number,omitempty"`
	BookingReference string                       `json:"booking_reference,omitempty"`
	Passengers       []AirlineTocketPassengerData `json:"passengers,omitempty"`
	Legs             []AirlineTicketLegData       `json:"legs"`
}

type AirlineTocketPassengerData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AirlineTicketLegData struct {
	DepartureAirport   string    `json:"departure_airport"`
	DepartureDate      time.Time `json:"departure_date"`
	DestinationAirport string    `json:"destination_airport"`
	CarrierCode        string    `json:"carrier_code"`
}