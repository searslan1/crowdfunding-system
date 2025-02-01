package validation

import (
	"entrepreneur/model"
	"regexp"
	"strings"
)

// ValidateEntrepreneur, girişimci profilinin geçerli olup olmadığını kontrol eder.
func ValidateEntrepreneur(entrepreneur *model.Entrepreneur) error {
	if err := validateStartupName(entrepreneur.StartupName); err != nil {
		return err
	}

	if err := validateIndustry(entrepreneur.Industry); err != nil {
		return err
	}

	if err := validateFundingNeeded(entrepreneur.FundingNeeded); err != nil {
		return err
	}

	if err := validatePitchDeckURL(entrepreneur.PitchDeckURL); err != nil {
		return err
	}

	return nil
}

// validateStartupName, girişimci adına dair geçerlilik kontrolü yapar.
func validateStartupName(startupName string) error {
	if strings.TrimSpace(startupName) == "" {
		return &model.AppError{
			Code:    "INVALID_STARTUP_NAME",
			Message: "Startup adı boş olamaz.",
		}
	}
	return nil
}

// validateIndustry, sektör adı geçerliliğini kontrol eder.
func validateIndustry(industry string) error {
	if strings.TrimSpace(industry) == "" {
		return &model.AppError{
			Code:    "INVALID_INDUSTRY",
			Message: "Sektör adı boş olamaz.",
		}
	}
	return nil
}

// validateFundingNeeded, gerekli fonlama tutarının geçerliliğini kontrol eder.
func validateFundingNeeded(fundingNeeded float64) error {
	if fundingNeeded <= 0 {
		return &model.AppError{
			Code:    "INVALID_FUNDING_NEEDED",
			Message: "Fonlama miktarı sıfırdan büyük olmalıdır.",
		}
	}
	return nil
}

// validatePitchDeckURL, pitch deck URL'sinin geçerliliğini kontrol eder.
func validatePitchDeckURL(pitchDeckURL *string) error {
	if pitchDeckURL != nil && !isValidURL(*pitchDeckURL) {
		return &model.AppError{
			Code:    "INVALID_PITCH_DECK_URL",
			Message: "Geçersiz Pitch Deck URL.",
		}
	}
	return nil
}

// isValidURL, geçerli bir URL olup olmadığını kontrol eder.
func isValidURL(url string) bool {
	re := regexp.MustCompile(`^(http|https):\/\/[^\s$.?#].[^\s]*$`)
	return re.MatchString(url)
}
