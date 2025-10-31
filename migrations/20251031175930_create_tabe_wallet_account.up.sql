-- testdb.wallet_account definition

CREATE TABLE `wallet_account` (
  `id` varchar(100) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `balance` decimal(19,6) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NOT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;