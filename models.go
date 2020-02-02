package truverifi

import (
	"strings"
	"time"
)

const (
	ErrorNotAvailable           = "NOT_AVAILABLE_FOR_THIS_ACCOUNT"
	ErrorNoLineAssigned         = "NO_LINE_ASSIGNED"
	ErrorZipCodeRejected        = "ZIP_CODE_REJECTED"
	ErrorNumberChangeInProgress = "NUMBER_CHANGE_IN_PROGRESS"
	ErrorBalanceInsufficient    = "BALANCE_INSUFFICIENT"
	ErrorNoLineAvailable        = "NO_LINE_AVAILABLE"
	ErrorInternalError          = "INTERNAL_ERROR"
	StatusReady                 = "ready"
	StatusPending               = "pending"
	StatusError                 = "error"
	StatusDisconnected          = "disconnected"
	ServiceNotListed            = "SERVICE_NOT_LISTED"
)

var (
	Services = []string{ServiceNotListed, "adidas", "airbnb", "alibaba", "amasia", "amazon", "authy", "avail", "bank_of_america", "beforthright", "bitmo", "cash_app", "cdkeys_com", "chase_bank", "chowbus", "classpass", "coinbase", "coupons_com", "craigslist", "crowdtap", "cryptopay", "dabbl", "didi", "digit", "discord", "dollarclix", "doordash", "dosh", "doublelist", "earnhoney", "ebay", "elevacity", "empower", "etoro", "e_rewards", "facebook", "fedex", "fetlife", "fiverr", "freelancer", "g2g_com_offgamers", "gameflip", "gemini", "gobank_green_dot", "gobranded", "google_gmail_voice", "grabpoint", "hopper", "hotmail", "imoney", "inboxdollars", "instagc", "instagram", "juno", "linkedin", "localbitcoins", "lyft", "mail_com", "metalpay", "mezu", "microsoft_azure", "microsoft_rewards", "nerdwallet", "netflix", "nike", "offernation", "offerup", "oneopinion", "outlook", "paxful", "paypal", "paysend", "pcgamesupply", "perk_com", "plenty_of_fish_pof", "postmates", "prolific", "proton", "pruvit", "purse", "retailmenot", "sea_gamer_mall", "service_not_listed", "shopkick", "simple", "skout", "snapchat", "snap_kitchen", "societi_tv", "spotify", "ssi_opinion_outpost_quickthoughts_save_with_surveys_streetbees", "stripe", "superpay_me", "supreme", "surveytime", "survey_honey", "survey_junkie", "swagbucks", "target", "telegram", "telnyx", "ticketmaster", "tinder", "transferwise", "turo", "twitter", "uber", "valued_opinion", "venmo", "walmart", "wechat", "whatsapp", "wirecash", "yahoo", "yandex", "yodlee", "zelle", "zillow", "zoosk"}
)

type AccountInfo struct {
	Username     string         `json:"username"`
	Balance      int            `json:"balance"`
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	ID          int       `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
}

type LineStatus struct {
	Client          *Truverifi
	PhoneNumber     string    `json:"phoneNumber"`
	Status          string    `json:"status"`
	ExpirationTime  time.Time `json:"expirationTime"`
	CurrentServices []string  `json:"currentServices"`
	SMS             []*SMS    `json:"sms"`
}

func (l *LineStatus) GetStatus() string {
	return strings.ToLower(l.Status)
}

type SMS struct {
	ID          int       `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Type        string    `json:"type"`
	PhoneNumber string    `json:"phoneNumber"`
	Text        string    `json:"text"`
}

type CheckServiceRequest struct {
	Zip      string   `json:"zip"`
	Services []string `json:"services"`
}

type CheckServiceResponse struct {
	Available         bool     `json:"available"`
	AvailableServices []string `json:"availableServices"`
	AvailableZips     []string `json:"availableZips"`
}

type ChangeServiceRequest struct {
	Zip      string   `json:"zip"`
	Services []string `json:"services"`
}

type ChangeServiceResponse struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ExtendResponse struct {
	ExpirationTime time.Time `json:"expirationTime"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ApiError struct {
	Message string
	Code    string
}

func (a *ApiError) Error() string {
	return a.Message
}
