package controllers


import (

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/pkg/configs"
	"backend/pkg/models"

	"context"
	"net/http"
	"time"
)

var expenseCollection *mongo.Collection = configs.GetCollection("expenses")
var expenseCategoryCollection *mongo.Collection = configs.GetCollection("expenseCategories")
var validate *validator.Validate = validator.New()

// Get all expense categories
func GetExpenseCategories(c *fiber.Ctx) error {
	collection := configs.GetCollection("expenseCategories")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense categories",
			"error": err,
		})
	}
	var expenseCategories []models.ExpenseCategory
	if err = cur.All(ctx, &expenseCategories); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense categories",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": expenseCategories,
	})
}

// Get expense category by id
func GetExpenseCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var expenseCategory models.ExpenseCategory
	err := expenseCategoryCollection.FindOne(context.Background(), filter).Decode(&expenseCategory)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense category",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": expenseCategory,
	})
}

// Create expense category
func CreateExpenseCategory(c *fiber.Ctx) error {
	var expenseCategory models.ExpenseCategory
	if err := c.BodyParser(&expenseCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse expense category",
			"error": err,
		})
	}
	if err := validate.Struct(expenseCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Expense category is invalid",
			"error": err,
		})
	}

	result, err := expenseCategoryCollection.InsertOne(context.Background(), expenseCategory)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot create expense category",
			"error": err,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}

// Update expense category
func UpdateExpenseCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var expenseCategory models.ExpenseCategory
	if err := c.BodyParser(&expenseCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse expense category",
			"error": err,
		})
	}
	if err := validate.Struct(expenseCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Expense category is invalid",
			"error": err,
		})
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": expenseCategory}
	result, err := expenseCategoryCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update expense category",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}

// Delete expense category
func DeleteExpenseCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	result, err := expenseCategoryCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete expense category",
			"error": err,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}
