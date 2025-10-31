package controller

import (
	"errors"
	"test-interview-kc/internal/usecase"
	"test-interview-kc/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type walletAccountController struct {
	walletAccountUseCase usecase.WalletAccountUseCase
	log                  *zap.Logger
}

func NewWalletAccountController(
	walletAccountUseCase usecase.WalletAccountUseCase,
	log *zap.Logger,
) *walletAccountController {
	return &walletAccountController{
		walletAccountUseCase: walletAccountUseCase,
		log:                  log,
	}
}

// GetAccountDetails handles the request to get wallet account details.
func (w *walletAccountController) GetAccountDetails(c *fiber.Ctx) error {
	ctx := c.UserContext()
	log := logger.FromContext(ctx, w.log)

	accountID := c.Params("id")
	if accountID == "" {
		log.Info("Account ID is missing in the request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "ACCOUNT_ID_MISSING",
			"message": "Account ID is required",
		})
	}

	account, err := w.walletAccountUseCase.GetAccountDetails(ctx, accountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("Wallet account not found", zap.String("account_id", accountID))
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "WALLET_ACCOUNT_NOT_FOUND",
				"message": "Wallet account not found",
			})
		}

		log.Error("Failed to retrieve account details", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "FAILED_TO_RETRIEVE_ACCOUNT_DETAILS",
			"message": "Failed to retrieve account details",
		})
	}

	return c.JSON(account)
}
