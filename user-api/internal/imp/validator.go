package imp

import (
	"errors"
	"regexp"
	"user-api/internal/data"
)

type Validator struct {
	logID string
}

func NewValidator(logId string) *Validator {
	return &Validator{logID: logId}
}

func (v *Validator) ValidatorSendSms(req *data.SendSmsReq) error {
	regRuler := "^1[345789]{1}\\d{9}$"
	rege := regexp.MustCompile(regRuler)
	res := rege.MatchString(req.Phone)
	if res {
		return nil
	}
	return errors.New("phone invalid")
}
