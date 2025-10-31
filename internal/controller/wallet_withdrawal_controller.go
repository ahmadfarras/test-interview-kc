package controller

import (
	"errors"
	"test-interview-kc/internal/dto/request"
	"test-interview-kc/internal/usecase"
	"test-interview-kc/pkg/logger"

	walletError "test-interview-kc/internal/error"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WalletWithdrawalController struct {
	walletWithdrawalUseCase usecase.WalletWithdrawalUseCase
	validator               *validator.Validate
	log                     *zap.Logger
}

func NewWalletWithdrawalController(
	wu usecase.WalletWithdrawalUseCase,
	validator *validator.Validate,
	log *zap.Logger,
) *WalletWithdrawalController {
	return &WalletWithdrawalController{
		walletWithdrawalUseCase: wu,
		validator:               validator,
		log:                     log,
	}
}

func (wwc *WalletWithdrawalController) Withdraw(c *fiber.Ctx) error {
	ctx := c.UserContext()
	log := logger.FromContext(ctx, wwc.log)

	walletID := c.Params("wallet_id")
	xRequestID := c.Get("X-Request-ID")
	if xRequestID == "" {
		log.Info("X-Request-ID is missing in the request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing X-Request-ID header"})
	}

	var req request.WalletWithdrawalRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	req.WalletID = walletID
	req.RequestID = xRequestID

	if err := wwc.validator.Struct(&req); err != nil {
		log.Error("Validation failed for withdrawal request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "VALIDATION_FAILED",
			"message": err.Error(),
		})
	}

	err := wwc.walletWithdrawalUseCase.Withdraw(ctx, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("Wallet account not found", zap.String("wallet_id", walletID))
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "WALLET_NOT_FOUND",
				"message": "Wallet account not found",
			})
		}

		if errors.Is(err, walletError.ErrInsufficientFunds) {
			log.Info("Insufficient funds", zap.String("wallet_id", walletID))
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   "INSUFFICIENT_FUNDS",
				"message": "Insufficient funds",
			})
		}

		if errors.Is(err, walletError.ErrIsAlreadyProcessed) {
			log.Info("Withdrawal request already processed", zap.String("request_id", xRequestID))
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   "ALREADY_PROCESSED",
				"message": "Withdrawal request already processed",
			})
		}

		log.Error("Withdrawal failed", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "WITHDRAWAL_FAILED",
			"message": "Withdrawal failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "withdrawal successful"})

}
