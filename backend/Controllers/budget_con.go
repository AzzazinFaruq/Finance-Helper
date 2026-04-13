package controllers

import (
	models "simple_crud/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type BudgetHandler struct {
	DB *bun.DB
}

func NewBudgetHandler(db *bun.DB) *BudgetHandler {
	return &BudgetHandler{DB: db}
}

func (h *BudgetHandler) CreateBudget(c *fiber.Ctx) error {

	type CreateBudgetRequest struct {
		UserID     int64   `json:"user_id"`
		CategoryID int64   `json:"category_id"`
		Amount     float64 `json:"amount"`
		Month      string  `json:"month"` // format: 2026-04
	}

	req := new(CreateBudgetRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	budget := &models.Budget{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Month:      req.Month,
	}

	_, err := h.DB.NewInsert().Model(budget).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Budget Created Succesfully",
		"data":    budget,
	})
}

func (h *BudgetHandler) GetBudget(c *fiber.Ctx) error {
	var budgets []models.Budget

	err := h.DB.NewSelect().
		Model(&budgets).
		Relation("User").
		Relation("Category").
		Scan(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(budgets) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message": "Budget not Found",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"data": budgets,
		})
	}
}

func (h *BudgetHandler) UpdateBudget(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	type UpdateBudgetRequest struct {
		UserID     int64   `json:"user_id"`
		CategoryID int64   `json:"category_id"`
		Amount     float64 `json:"amount"`
		Month      string  `json:"month"`
	}

	req := new(UpdateBudgetRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
	}

	budget := &models.Budget{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Month:      req.Month,
	}

	q := h.DB.NewUpdate().
		Model(budget).
		Column("user_id", "category_id", "amount", "month").
		Where("id = ?", id)

	result, err := q.Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Budget tidak ditemukan atau akses ditolak"})
	}

	return c.JSON(fiber.Map{"message": "Budget berhasil diupdate"})
}

func (h *BudgetHandler) DeleteBudget(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.DB.NewDelete().
		Model((*models.Budget)(nil)).
		Where("id = ?", id).
		Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Budget tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "Budget berhasil dihapus"})
}
