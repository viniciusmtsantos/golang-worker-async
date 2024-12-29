package distributors

type PayloadCreditUserPoints struct {
	IndicatorID int64 `json:"indicator_id"`
	WantRetry   bool  `json:"want_retry"`
}
