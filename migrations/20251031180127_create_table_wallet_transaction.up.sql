-- testdb.wallet_transaction definition

CREATE TABLE `wallet_transaction` (
  `id` varchar(100) NOT NULL,
  `amount` decimal(19,6) DEFAULT NULL,
  `wallet_account_id` varchar(100) NOT NULL,
  `request_id` varchar(255) DEFAULT NULL,
  `type` enum('CREDIT','DEBIT') DEFAULT NULL,
  `entry_type` enum('WITHDRAWAL','DEPOSIT','TRANSFER','PAYMENT') DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `transaction_date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `updated_at` timestamp NOT NULL,
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;