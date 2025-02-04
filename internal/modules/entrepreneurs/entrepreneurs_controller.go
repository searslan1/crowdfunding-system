package controller

import (
	"entrepreneur/model"
	"entrepreneur/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// EntrepreneurController girişimci işlemlerini yönetir
type EntrepreneurController struct {
	EntrepreneurService *service.EntrepreneurService
}

// CreateEntrepreneur girişimci profili oluşturur
func (ec *EntrepreneurController) CreateEntrepreneur(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// SPK Gerekliliği: E-Devlet onayı kontrolü
	if !ec.EntrepreneurService.IsEDevletVerified(userID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Kullanıcı e-Devlet doğrulamasından geçmemiş.",
		})
	}

	// Kullanıcının zaten bir girişimci profili olup olmadığı kontrol edilir
	existingProfile, _ := ec.EntrepreneurService.GetByUserID(userID)
	if existingProfile != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Kullanıcı zaten bir girişimci profiline sahip.",
		})
	}

	var entrepreneur model.Entrepreneur
	if err := c.BodyParser(&entrepreneur); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz girişimci verisi. Lütfen girişimci profili bilgilerini kontrol edin.",
		})
	}

	entrepreneur.UserID = userID
	entrepreneur.Status = "pending" // Admin onayı bekleniyor

	if err := ec.EntrepreneurService.Create(&entrepreneur); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Girişimci profili oluşturulamadı.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entrepreneur)
}

// AdminApproveEntrepreneur girişimci profilini admin tarafından onaylar
func (ec *EntrepreneurController) AdminApproveEntrepreneur(c *fiber.Ctx) error {
	entrepreneurID := c.Params("id")

	// Girişimci profili veritabanından alınıyor
	entrepreneurIDParsed, err := strconv.Atoi(entrepreneurID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz girişimci ID'si.",
		})
	}

	entrepreneur, err := ec.EntrepreneurService.GetByUserID(uint(entrepreneurIDParsed))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Girişimci profili bulunamadı.",
		})
	}

	// Eğer girişimci profili zaten onaylı ise, tekrar onay verilmesine gerek yok
	if entrepreneur.Status == "approved" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bu girişimci zaten onaylanmış.",
		})
	}

	// Girişimci profilinin statüsünü onaylı hale getir
	entrepreneur.Status = "approved"
	entrepreneur.IsAdminApproved = true

	// Girişimci profili veritabanında güncelleniyor
	if err := ec.EntrepreneurService.Update(entrepreneur); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Girişimci profili güncellenemedi.",
		})
	}

	// Girişimci profilini başarıyla onayladık ve updated profili geri döndürüyoruz
	return c.JSON(fiber.Map{
		"message": "Girişimci profili başarıyla onaylandı.",
		"data": entrepreneur,
	})
}
