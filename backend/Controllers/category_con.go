package controllers

import (
	"simple_crud/Models"
	// "strconv"

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
		Name:    req.Name,
		UserID: req.UserID,
	}

	_, err := h.DB.NewInsert().Model(category).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Category Created Succesfully",
		"data":   category,
	})
}

// func (h *CategoryHandler) GetClinic(c *fiber.Ctx) error {
// 	var clinics []models.Clinic

// 	err := h.DB.NewSelect().
// 		Model(&clinics).
// 		Order("id ASC").
// 		Scan(c.Context())

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(clinics)
// }

// func (h *CategoryHandler) UpdateClinic(c *fiber.Ctx) error {
// 	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

// 	// Gunakan struct input terpisah agar fleksibel
// 	type UpdateUserRequest struct {
// 		Name    string `json:"name"`
// 		Address string `json:"address"`
// 		Phone   string `json:"phone"`
// 		Slug    string `json:"slug"`
// 	}

// 	req := new(UpdateUserRequest)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
// 	}

// 	// Siapkan user model untuk update
// 	clinic := &models.Clinic{
// 		ID:      id,
// 		Name:    req.Name,
// 		Address: req.Address,
// 		Phone:   req.Phone,
// 		Slug:    req.Slug,
// 	}

// 	// Query Builder Update
// 	q := h.DB.NewUpdate().
// 		Model(clinic).
// 		Column("name", "address", "phone", "slug").
// 		Where("id = ?", id)

// 	result, err := q.Exec(c.Context())
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	// Cek apakah ada baris yang terupdate
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return c.Status(404).JSON(fiber.Map{"error": "Klinik tidak ditemukan atau akses ditolak"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Klinik berhasil diupdate"})
// }

// func (h *CategoryHandler) DeleteClinic(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	result, err := h.DB.NewDelete().
// 		Model((*models.Clinic)(nil)).
// 		Where("id = ?", id).
// 		Exec(c.Context())

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	rows, _ := result.RowsAffected()
// 	if rows == 0 {
// 		return c.Status(404).JSON(fiber.Map{"error": "Klinik tidak ditemukan"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Klinik berhasil dihapus"})
// }
