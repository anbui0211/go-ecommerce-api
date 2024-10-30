SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pre_go_acc_user_9999
-- ----------------------------
DROP TABLE IF EXISTS `pre_go_acc_user_9999`;
CREATE TABLE `pre_go_acc_user_9999` (
    `user_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
    `user_account` VARCHAR(255) NOT NULL COMMENT 'User account',
    `user_nickname` VARCHAR(255) NULL DEFAULT 'DEFAULT NULL' COMMENT 'User nickname',
    `user_avatar` VARCHAR(255) NULL DEFAULT 'DEFAULT NULL' COMMENT 'User avatar',
    `user_state` TINYINT UNSIGNED NOT NULL COMMENT 'User state: 0-Locked ,  1-Activated ,  2-Not Activated',
    `user_mobile` VARCHAR(20) NULL DEFAULT 'DEFAULT NULL' COMMENT 'Mobile phone number',
    `user_gender` TINYINT UNSIGNED NULL DEFAULT 'DEFAULT NULL' COMMENT 'User gender: 0-Secret ,  1-Male ,  2-Female',
    `user_birthday` DATE NULL DEFAULT 'DEFAULT NULL' COMMENT 'User birthday',
    `user_email` VARCHAR(255) NULL DEFAULT 'DEFAULT NULL' COMMENT 'User email address',
    `user_is_authentication` TINYINT UNSIGNED NOT NULL COMMENT 'Authentication status: 0-Not Authenticated ,  1-Pending ,  2-Authenticated ,  3-Failed',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP() COMMENT 'Record creation time',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP() COMMENT 'Record update time',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `unique_user_account` (`user_account`),
    KEY `idx_user mobile` (`user_mobile`),
    KEY `idx_user_email` (`user_email`),
    KEY `idx_user_state` (`user_state`),
    KEY `idx_user_is_authentication` (`user_is_authentication`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARACTER = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'pre_go_acc_user_9999';

-- ----------------------------
-- Table structure for pre_go_acc_user_base_9999
-- ----------------------------
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
CREATE TABLE `pre_go_acc_user_base_9999` (
    `user_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_account` VARCHAR(255) NOT NULL,
    `user_password` VARCHAR(255) NOT NULL,
    `user_salt` VARCHAR(255) NOT NULL,
    `user_login_time` TIMESTAMP NULL DEFAULT 'DEFAULT NULL',
    `user_logout_time` TIMESTAMP NULL DEFAULT 'DEFAULT NULL',
    `user_login_ip` VARCHAR(45) NULL DEFAULT 'DEFAULT NULL',
    `user_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
    `user_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `unique_user_account` (`user_account`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARACTER = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'pre_go_acc_user_base_9999';

-- ----------------------------
-- Table structure for pre_go_acc_user_verify_9999
-- ----------------------------
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
CREATE TABLE `pre_go_acc_user_verify_9999` (
    `verify_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `verify_otp` VARCHAR(6) NOT NULL,
    `verify_key` VARCHAR(255) NOT NULL,
    `verify_key_hash` VARCHAR(255) NOT NULL,
    `verify_type` INT NULL DEFAULT '1',
    `is_verified` INT NULL,
    `is_deleted` INT NULL,
    `verify_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
    `verify_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(),
    PRIMARY KEY(`verify_id`),
    UNIQUE KEY `unique_verify_key` (`verify_key`)
    KEY `index_verify_otp` (`verify_otp`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARACTER=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='pre_go_acc_user_verify_9999';

SET FOREIGN_KEY_CHECK = 1;

