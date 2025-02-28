package sdk

import "encoding/json"

type PasswordPolicyPasswordSettingsAge struct {
	ExpireWarnDays    int64  `json:"-"`
	ExpireWarnDaysPtr *int64 `json:"expireWarnDays,omitempty"`
	HistoryCount      int64  `json:"-"`
	HistoryCountPtr   *int64 `json:"historyCount,omitempty"`
	MaxAgeDays        int64  `json:"-"`
	MaxAgeDaysPtr     *int64 `json:"maxAgeDays,omitempty"`
	MinAgeMinutes     int64  `json:"-"`
	MinAgeMinutesPtr  *int64 `json:"minAgeMinutes,omitempty"`
}

func NewPasswordPolicyPasswordSettingsAge() *PasswordPolicyPasswordSettingsAge {
	return &PasswordPolicyPasswordSettingsAge{
		ExpireWarnDays:    0,
		ExpireWarnDaysPtr: Int64Ptr(0),
		HistoryCount:      0,
		HistoryCountPtr:   Int64Ptr(0),
		MaxAgeDays:        0,
		MaxAgeDaysPtr:     Int64Ptr(0),
		MinAgeMinutes:     0,
		MinAgeMinutesPtr:  Int64Ptr(0),
	}
}

func (a *PasswordPolicyPasswordSettingsAge) IsPolicyInstance() bool {
	return true
}

func (a *PasswordPolicyPasswordSettingsAge) MarshalJSON() ([]byte, error) {
	type Alias PasswordPolicyPasswordSettingsAge
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.ExpireWarnDays != 0 {
		result.ExpireWarnDaysPtr = Int64Ptr(a.ExpireWarnDays)
	}
	if a.HistoryCount != 0 {
		result.HistoryCountPtr = Int64Ptr(a.HistoryCount)
	}
	if a.MaxAgeDays != 0 {
		result.MaxAgeDaysPtr = Int64Ptr(a.MaxAgeDays)
	}
	if a.MinAgeMinutes != 0 {
		result.MinAgeMinutesPtr = Int64Ptr(a.MinAgeMinutes)
	}
	return json.Marshal(&result)
}

func (a *PasswordPolicyPasswordSettingsAge) UnmarshalJSON(data []byte) error {
	type Alias PasswordPolicyPasswordSettingsAge

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.ExpireWarnDaysPtr != nil {
		a.ExpireWarnDays = *result.ExpireWarnDaysPtr
		a.ExpireWarnDaysPtr = result.ExpireWarnDaysPtr
	}
	if result.HistoryCountPtr != nil {
		a.HistoryCount = *result.HistoryCountPtr
		a.HistoryCountPtr = result.HistoryCountPtr
	}
	if result.MaxAgeDaysPtr != nil {
		a.MaxAgeDays = *result.MaxAgeDaysPtr
		a.MaxAgeDaysPtr = result.MaxAgeDaysPtr
	}
	if result.MinAgeMinutesPtr != nil {
		a.MinAgeMinutes = *result.MinAgeMinutesPtr
		a.MinAgeMinutesPtr = result.MinAgeMinutesPtr
	}
	return nil
}
