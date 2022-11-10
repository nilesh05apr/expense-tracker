package controllers

import (
	"fmt"

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
	UpdateCategoryTotalAmount(c , expense.ExpenseCategoryID)
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
	UpdateCategoryTotalAmount(c , expense.ExpenseCategoryID)
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
	var expense models.Expense
	UpdateCategoryTotalAmount(c , expense.ExpenseCategoryID)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": result,
	})
}

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
	// update amount for each category
	for i := 0; i < len(expenseCategories); i++ {
		UpdateCategoryTotalAmount(c , expenseCategories[i].ID)
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

// Update Category total amount
func UpdateCategoryTotalAmount(c *fiber.Ctx, objID primitive.ObjectID) error {
	// id := c.Params("id")
	// objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"expenseCategoryId": objID}

	// Get total amount of expense category
	cur, err := expenseCollection.Find(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense category total amount",
			"error": err,
		})
	}
	var expenses []models.Expense
	if err = cur.All(context.Background(), &expenses); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot get expense category total amount",
			"error": err,
		})
	}
	var totalAmount float64
	for _, expense := range expenses {
		totalAmount += expense.Amount
	}
	
	// Update expense category total amount
	filter = bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"Amount": totalAmount}}
	result, err := expenseCategoryCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update expense category total amount",
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

