package service

import (
	"entrepreneur/model"
	"entrepreneur/repository"

	"golang.org/x/crypto/bcrypt"
)

// EntrepreneurService, girişimci profili işlemlerini yönetir.
type EntrepreneurService struct {
	repo *repository.EntrepreneurRepository
	
}

// NewEntrepreneurService, yeni bir EntrepreneurService örneği oluşturur.
func NewEntrepreneurService(repo *repository.EntrepreneurRepository) *EntrepreneurService {
	return &EntrepreneurService{repo: repo}
}

// HashPassword, verilen şifreyi bcrypt ile hash'ler.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CreateEntrepreneur, yeni bir girişimci profili oluşturur.
func (s *EntrepreneurService) CreateEntrepreneur(entrepreneur *model.Entrepreneur, rawPassword string) error {
	// Aynı kullanıcıya ait girişimci profili olup olmadığını kontrol et.
	existing, _ := s.repo.GetByUserID(entrepreneur.UserID)
	if existing != nil {
		return ErrEntrepreneurExists
	}

	// Şifreyi hash'leyelim.
	hashedPassword, err := HashPassword(rawPassword)
	if err != nil {
		return err
	}

	// Şifreyi girişimci profiline ekleyelim.
	entrepreneur.BusinessModel = &hashedPassword // Örnek olarak BusinessModel alanına ekledik, uygun bir alan eklemelisin.

	// Profili oluştur.
	return s.repo.Create(entrepreneur)
}
