package validation

import (
	"errors"
	"estrepreneur/model"
)

// ValidateEntrepreneur girişimci verilerini doğrular
func ValidateEntrepreneur(e *model.Entrepreneur) error {
	if e.UserID == 0 {
		return errors.New("UserID boş olamaz")
	}
	if e.StartupName == "" {
		return errors.New("Startup adı boş olamaz")
	}
	if e.Industry == "" {
		return errors.New("Sektör bilgisi boş olamaz")
	}
	if e.FundingNeeded <= 0 {
		return errors.New("Yatırım ihtiyacı sıfırdan büyük olmalıdır")
	}
	return nil
}
