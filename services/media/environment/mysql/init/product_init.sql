CREATE DATABASE IF NOT EXISTS shopdevgo
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

-- Product table
CREATE TABLE IF NOT EXISTS `shopdevgo`.`product` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(255) NOT NULL COMMENT 'Product name',
    `description` TEXT COMMENT 'Product description',
    `price` DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Product price',
    `discount_price` DECIMAL(15,2) DEFAULT NULL COMMENT 'Discounted price',
    `quantity` INT(11) NOT NULL DEFAULT 0 COMMENT 'Available quantity',
    `category_id` BIGINT(20) DEFAULT NULL COMMENT 'Category ID',
    `thumbnail` VARCHAR(255) DEFAULT NULL COMMENT 'Thumbnail image URL',
    `images` JSON DEFAULT NULL COMMENT 'Additional product images',
    `attributes` JSON DEFAULT NULL COMMENT 'Product attributes (color, size, etc.)',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Status: 1-active, 0-inactive, -1-deleted',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`),
    KEY `idx_category` (`category_id`),
    KEY `idx_status` (`status`),
    KEY `idx_price` (`price`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Product table';

-- Insert sample product data
INSERT INTO `shopdevgo`.`product` (`name`, `description`, `price`, `discount_price`, `quantity`, `category_id`, `thumbnail`, `attributes`, `status`)
VALUES
    ('iPhone 15 Pro', 'Latest iPhone model with advanced features', 1299.99, 1199.99, 100, 1, 'https://example.com/iphone15.jpg', '{"color": ["Black", "Silver", "Gold"], "storage": ["128GB", "256GB", "512GB"]}', 1),
    ('Samsung Galaxy S23', 'Flagship Android smartphone', 999.99, 899.99, 150, 1, 'https://example.com/galaxys23.jpg', '{"color": ["Black", "White", "Green"], "storage": ["128GB", "256GB"]}', 1),
    ('MacBook Pro 16"', 'Powerful laptop for professionals', 2499.99, 2299.99, 50, 2, 'https://example.com/macbookpro.jpg', '{"color": ["Space Gray", "Silver"], "ram": ["16GB", "32GB"], "storage": ["512GB", "1TB", "2TB"]}', 1),
    ('Sony WH-1000XM5', 'Premium noise-cancelling headphones', 399.99, 349.99, 200, 3, 'https://example.com/sonywh1000xm5.jpg', '{"color": ["Black", "Silver"]}', 1),
    ('Nintendo Switch OLED', 'Gaming console with enhanced display', 349.99, NULL, 75, 4, 'https://example.com/switcholed.jpg', '{"color": ["White", "Neon"]}', 1);