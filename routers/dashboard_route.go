package routers

import (
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/controllers"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/database"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/middlewares"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/repositories"
	"github.com/RomuloNd/Projeto-CruzeiroDoSul/services"
	"github.com/gofiber/fiber/v2"
)

type DashboardRouter struct {
	dashboardController controllers.DashboardController
}

func NewDashboardRouter() *DashboardRouter {
	// Initialize repositories
	dashboardRepository := repositories.NewDashboardRepository(database.DB)

	// Initialize services with repositories
	dashboardService := services.NewDashboardService(dashboardRepository)

	// Initialize controllers with services
	dashboardController := controllers.NewDashboardController(dashboardService)

	return &DashboardRouter{
		dashboardController: dashboardController,
	}
}

func (r *DashboardRouter) InstallRouters(app *fiber.App) {
	dashboard := app.Group("/dashboard").Use(middlewares.LoginAndStaffRequired())
	dashboard.Get("/", r.dashboardController.RenderDashboard)
}
