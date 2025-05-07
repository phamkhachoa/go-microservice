-- Create the inventory database if it doesn't exist
CREATE DATABASE IF NOT EXISTS shopdevgo_inventory
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

-- Use the inventory database
USE shopdevgo_inventory;

-- Create the product_inventory table
CREATE TABLE IF NOT EXISTS `product_inventory` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `product_id` BIGINT(20) NOT NULL COMMENT 'Foreign key referencing the product table',
    `quantity` INT(11) NOT NULL DEFAULT 0 COMMENT 'Available quantity',
    `reserved_quantity` INT(11) NOT NULL DEFAULT 0 COMMENT 'Quantity reserved for ongoing transactions',
    `reorder_point` INT(11) DEFAULT NULL COMMENT 'Quantity at which to reorder stock',
    `reorder_quantity` INT(11) DEFAULT NULL COMMENT 'Quantity to reorder when stock is low',
    `last_restock_date` DATETIME DEFAULT NULL COMMENT 'Date of last restock',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_product_id` (`product_id`),
    KEY `idx_quantity` (`quantity`),
    KEY `idx_reorder_point` (`reorder_point`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Product inventory table';

-- Insert sample data (optional)
INSERT INTO `product_inventory` (`product_id`, `quantity`, `reserved_quantity`, `reorder_point`, `reorder_quantity`)
VALUES
    (1, 100, 0, 20, 50),  -- iPhone 15 Pro
    (2, 150, 0, 30, 75),  -- Samsung Galaxy S23
    (3, 50, 0, 10, 25),   -- MacBook Pro 16"
    (4, 200, 0, 40, 100), -- Sony WH-1000XM5
    (5, 75, 0, 15, 30);   -- Nintendo Switch OLED