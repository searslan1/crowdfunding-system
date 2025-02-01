package routes

import (
	"entrepreneur/controller"

	"github.com/gin-gonic/gin"
)

// RegisterEntrepreneurRoutes, girişimci profilleri için API endpointlerini tanımlar.
func RegisterEntrepreneurRoutes(router *gin.Engine, entrepreneurController *controller.EntrepreneurController) {
	entrepreneurRoutes := router.Group("/entrepreneurs")
	{
		entrepreneurRoutes.POST("/", entrepreneurController.CreateEntrepreneur)       // Yeni girişimci profili oluştur
		entrepreneurRoutes.GET("/:id", entrepreneurController.GetEntrepreneurByID)     // ID'ye göre girişimci profili getir
		entrepreneurRoutes.GET("/user/:user_id", entrepreneurController.GetEntrepreneurByUserID) // Kullanıcı ID'ye göre girişimci profili getir
		entrepreneurRoutes.PUT("/", entrepreneurController.UpdateEntrepreneur)        // Girişimci profilini güncelle
		entrepreneurRoutes.DELETE("/:id", entrepreneurController.DeleteEntrepreneur)  // Girişimci profilini sil
	}
}
