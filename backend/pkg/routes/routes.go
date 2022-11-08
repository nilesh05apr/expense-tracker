package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/pkg/controllers"

)

func Routes(app *fiber.App) {
	app.Get("/api/expenses", controllers.GetExpenses)
	app.Get("/api/expenses/:id", controllers.GetExpense)
	app.Get("/api/expenses/category/:id", controllers.GetExpenseByCategory)
	app.Post("/api/expenses", controllers.CreateExpense)
	app.Put("/api/expenses/:id", controllers.UpdateExpense)
	app.Delete("/api/expenses/:id", controllers.DeleteExpense)
	app.Get("/api/expenseCategories", controllers.GetExpenseCategories)
	app.Get("/api/expenseCategories/:id", controllers.GetExpenseCategory)
	app.Post("/api/expenseCategories", controllers.CreateExpenseCategory)
	app.Put("/api/expenseCategories/:id", controllers.UpdateExpenseCategory)
	app.Delete("/api/expenseCategories/:id", controllers.DeleteExpenseCategory)
}