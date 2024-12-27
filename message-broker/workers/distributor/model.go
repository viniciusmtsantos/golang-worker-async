package distributor

type PayloadCreditUserPoints struct {
	ReferralID     int64 `json:"referral_id"`
	ReferrerUserID int64 `json:"referrer_user_id"`
}
