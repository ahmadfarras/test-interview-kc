package main

import (
	"fmt"
	"test-interview-kc/internal/controller"
	"test-interview-kc/internal/middleware"
	"test-interview-kc/internal/repository"
	"test-interview-kc/internal/usecase"
	"test-interview-kc/pkg/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Load config
	config := utils.MustLoadConfig("env.yaml")

	logger := utils.InitLogger(config.App.Env)
	defer logger.Sync()

	// Connect to MySQL using GORM
	db, err := utils.NewMySQLConnection(config)
	if err != nil {
		logger.Fatal("failed to connect to MySQL: ", zap.Error(err))
	}

	// Initialize Fiber
	app := fiber.New()
	app.Use(middleware.LoggerMiddleware(logger))

	// Health route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Initialize Validator
	requestValidator := validator.New()

	// --- Initialize Repositories
	walletAccountRepo := repository.NewWalletAccountRepository(db, logger)
	walletTransactionRepo := repository.NewWalletTransactionRepository(db, logger)

	// --- Initialize UseCases
	walletWithdrawalUsecase := usecase.NewWalletWithdrawalUseCase(walletAccountRepo, walletTransactionRepo, logger)
	walletAccountUsecase := usecase.NewWalletAccountUseCase(walletAccountRepo, logger)

	// --- Initialize Controllers
	walletWithdrawalController := controller.NewWalletWithdrawalController(walletWithdrawalUsecase, requestValidator, logger)
	walletAccountController := controller.NewWalletAccountController(walletAccountUsecase, logger)

	// --- Define Routes
	app.Get("/wallets/:id", walletAccountController.GetAccountDetails)
	app.Post("/wallets/:wallet_id/withdraw", walletWithdrawalController.Withdraw)

	// Start server
	addr := ":" + fiberPort(config.App.Port)
	logger.Info("Starting server on " + addr)
	if err := app.Listen(addr); err != nil {
		logger.Fatal("failed to start server: ", zap.Error(err))
	}
}

func fiberPort(port int) string {
	if port == 0 {
		return "8080"
	}
	return fmt.Sprintf("%d", port)
}
