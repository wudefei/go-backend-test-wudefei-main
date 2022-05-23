package service

type SmsService struct {
	logID string
}

func NewSmsService(logId string) *SmsService {
	return &SmsService{logID: logId}
}

func (s *SmsService) SendSms(phone string, verifiCode string) error {
	//---TODO
	return nil
}
