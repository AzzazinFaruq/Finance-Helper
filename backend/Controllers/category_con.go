package controllers

import (
	models "simple_crud/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type CategoryHandler struct {
	DB *bun.DB
}

func NewCategoryHandler(db *bun.DB) *CategoryHandler {
	return &CategoryHandler{DB: db}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {

	type CreateCategoryRequest struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
	}

	req := new(CreateCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}
	category := &models.Category{
		Name:   req.Name,
		UserID: req.UserID,
	}

	_, err := h.DB.NewInsert().Model(category).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Category Created Succesfully",
		"data":    category,
	})
}

func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	var categories []models.Category

	err := h.DB.NewSelect().
		Model(&categories).
		Relation("User").
		Scan(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(categories) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message": "Category not Found",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "Category Found",
			"data":    categories,
		})
	}

}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	type UpdateUserRequest struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
	}

	req := new(UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
	}

	category := &models.Category{
		Name:   req.Name,
		UserID: req.UserID,
	}

	// Query Builder Update
	q := h.DB.NewUpdate().
		Model(category).
		Column("name", "user_id").
		Where("id = ?", id)

	result, err := q.Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Cek apakah ada baris yang terupdate
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Category tidak ditemukan atau akses ditolak"})
	}

	return c.JSON(fiber.Map{"message": "Category berhasil diupdate"})
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.DB.NewDelete().
		Model((*models.Category)(nil)).
		Where("id = ?", id).
		Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Category tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "Category berhasil dihapus"})
}
