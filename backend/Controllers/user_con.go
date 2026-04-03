package controllers

import (
	models "simple_crud/Models"
	utils "simple_crud/Utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	DB *bun.DB
}

func NewUserHandler(db *bun.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	req := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal enkripsi password"})
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPwd,
	}

	if _, err := h.DB.NewInsert().Model(user).Exec(c.Context()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data":    user,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username dan password wajib diisi"})
	}

	user := new(models.User)
	err := h.DB.NewSelect().
		Model(user).
		Where("username = ?", req.Username).
		Limit(1).
		Scan(c.Context())
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username atau password salah"})
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username atau password salah"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat token"})
	}

	exp := time.Now().Add(24 * time.Hour)
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  exp,
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"status": true,
		"token":  token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Logout berhasil",
	})
}

func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {

	user := &models.User{}

	userID := c.Locals("user")

	err := h.DB.NewSelect().
		Model(user).
		Where("id = ?", userID).
		Scan(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}
	return c.JSON(fiber.Map{
		"user": user,
	})

}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Locals("user")

	type UpdateUserRequest struct {
		Username string `json:"username"`
	}

	req := new(UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
	}

	user := &models.User{
		Username:  req.Username,
		UpdatedAt: time.Now(),
	}

	q := h.DB.NewUpdate().
		Model(user).
		Column("username", "updated_at").
		Where("id = ?", id)

	result, err := q.Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan atau akses ditolak"})
	}

	return c.JSON(fiber.Map{"message": "User berhasil diupdate"})
}

