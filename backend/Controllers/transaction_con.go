package controllers

import (
	models "simple_crud/Models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type TransactionHandler struct {
	DB *bun.DB
}

func NewTransactionHandler(db *bun.DB) *TransactionHandler {
	return &TransactionHandler{DB: db}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {

	type CreateTransactionRequest struct {
		UserID      int64   `json:"user_id"`
		Amount      float64 `json:"amount"`
		Type        string  `json:"type"`
		Date        string  `json:"date"` 
		Description string  `json:"description"`
		CategoryID  int64   `json:"category_id"`
	}

	req := new(CreateTransactionRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal salah, gunakan YYYY-MM-DD"})
	}

	transaction := &models.Transaction{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Type:        req.Type,
		Date:        parsedDate,
		Description: req.Description,
		CategoryID:  req.CategoryID,
	}

	_, err = h.DB.NewInsert().Model(transaction).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Transaction Created Succesfully",
		"data":    transaction,
	})
}

func (h *TransactionHandler) GetTransaction(c *fiber.Ctx) error {
	var transactions []models.Transaction

	err := h.DB.NewSelect().
		Model(&transactions).
		Relation("User").
		Relation("Category").
		Scan(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(transactions) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message": "Transaction not Found",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"data":    transactions,
		})
	}

}

func (h *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	type UpdateTransactionRequest struct {
		UserID      int64   `json:"user_id"`
		Amount      float64 `json:"amount"`
		Type        string  `json:"type"`
		Date        string  `json:"date"` 
		Description string  `json:"description"`
		CategoryID  int64   `json:"category_id"`
	}

	req := new(UpdateTransactionRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal salah, gunakan YYYY-MM-DD"})
	}

	category := &models.Transaction{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Type:        req.Type,
		Date:        parsedDate,
		Description: req.Description,
		CategoryID:  req.CategoryID,
	}

	// Query Builder Update
	q := h.DB.NewUpdate().
		Model(category).
		Column("user_id", "amount", "type", "date", "description", "category_id").
		Where("id = ?", id)

	result, err := q.Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Cek apakah ada baris yang terupdate
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Transaction tidak ditemukan atau akses ditolak"})
	}

	return c.JSON(fiber.Map{"message": "Transaction berhasil diupdate"})
}

func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.DB.NewDelete().
		Model((*models.Transaction)(nil)).
		Where("id = ?", id).
		Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Transaction tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "Transaction berhasil dihapus"})
}
