package routers

import (
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/controllers"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/database"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/middlewares"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/repositories"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/services"
	"github.com/gofiber/fiber/v2"
)

type ProfileRouter struct {
	profileController controllers.ProfileController
}

func NewProfileRouter() *ProfileRouter {
	// Initialize repositories
	profileRepository := repositories.NewProfileRepository(database.DB)

	// Initialize services with repositories
	profileService := services.NewProfileService(profileRepository)

	// Initialize controllers with services
	profileController := controllers.NewProfileController(profileService)

	return &ProfileRouter{
		profileController: profileController,
	}
}

func (r *ProfileRouter) InstallRouters(app *fiber.App) {
	profile := app.Group("/profile").Use(middlewares.LoginRequired())

	profile.Get("/:id", r.profileController.RenderProfile)
	profile.Post("/update/:id", r.profileController.Update)
}
