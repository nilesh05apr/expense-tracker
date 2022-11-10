package controllers

import (
	"fmt"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"

	"backend/pkg/configs"
	"backend/pkg/models"

	"context"
	"net/http"
	"time"
)


// var expenseCollection *mongo.Collection = configs.GetCollection("expenses")
// var expenseCategoryCollection *mongo.Collection = configs.GetCollection("expenseCategories")
// var validate *validator.Validate = validator.New()

// Get all expenses
func GetExpenses(c *fiber.Ctx) error {
	collection := configs.GetCollection("expenses")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expenses",
			"error": err,
		})
	}
	var expenses []models.Expense
	if err = cur.All(ctx, &expenses); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expenses",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": expenses,
	})
}

// Get expense by id
func GetExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var expense models.Expense
	err := expenseCollection.FindOne(context.Background(), filter).Decode(&expense)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": expense,
	})
}

// Get all expense by category
func GetExpenseByCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"expenseCategoryId": objID}
	var expenses []models.Expense
	cur, err := expenseCollection.Find(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense by category",
			"error": err,
		})
	}
	if err = cur.All(context.Background(), &expenses); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense by category",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": expenses,
	})
}


// Create expense
func CreateExpense(c *fiber.Ctx) error {
	var expense models.Expense
	if err := c.BodyParser(&expense); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse expense",
			"error": err,
		})
	}
	if err := validate.Struct(expense); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Expense is invalid",
			"error": err,
		})
	}
	expense.CreatedAt = time.Now()
	result, err := expenseCollection.InsertOne(context.Background(), expense)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot create expense",
			"error": err,
		})
	}
	fmt.Println(expense.ExpenseCategoryID)
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}

// Update expense
func UpdateExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var expense models.Expense
	if err := c.BodyParser(&expense); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse expense",
			"error": err,
		})
	}
	if err := validate.Struct(expense); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Expense is invalid",
			"error": err,
		})
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": expense}
	result, err := expenseCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update expense",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}	

// Delete expense
func DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	result, err := expenseCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete expense",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}


