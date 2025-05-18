-- Product detail table
CREATE TABLE IF NOT EXISTS `shopdevgo`.`product_detail` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `product_id` BIGINT(20) NOT NULL COMMENT 'Reference to product',
    `type` VARCHAR(50) NOT NULL COMMENT 'Detail type (specification, feature, warranty, etc.)',
    `name` VARCHAR(100) NOT NULL COMMENT 'Detail name',
    `value` VARCHAR(255) NOT NULL COMMENT 'Detail value',
    `display_order` INT(11) NOT NULL DEFAULT 0 COMMENT 'Display order',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Status: 1-active, 0-inactive, -1-deleted',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`),
    KEY `idx_product_id` (`product_id`),
    KEY `idx_type` (`type`),
    KEY `idx_status` (`status`),
    CONSTRAINT `fk_product_detail_product` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Product detail table';

-- Insert sample product detail data
INSERT INTO `shopdevgo`.`product_detail` (`product_id`, `type`, `name`, `value`, `display_order`)
VALUES
    (1, 'specification', 'Processor', 'A16 Bionic', 1),
    (1, 'specification', 'Display', '6.7-inch Super Retina XDR', 2),
    (1, 'specification', 'Camera', '48MP main camera', 3),
    (1, 'feature', 'Water Resistance', 'IP68 rated', 1),
    (1, 'warranty', 'Standard Warranty', '1 year manufacturer warranty', 1),
    
    (2, 'specification', 'Processor', 'Snapdragon 8 Gen 2', 1),
    (2, 'specification', 'Display', '6.8-inch Dynamic AMOLED', 2),
    (2, 'specification', 'Camera', '200MP main camera', 3),
    (2, 'feature', 'S Pen Support', 'Built-in S Pen', 1),
    (2, 'warranty', 'Standard Warranty', '1 year manufacturer warranty', 1),
    
    (3, 'specification', 'Processor', 'M2 Pro/Max', 1),
    (3, 'specification', 'Display', '16-inch Liquid Retina XDR', 2),
    (3, 'specification', 'Memory', 'Up to 64GB unified memory', 3),
    (3, 'feature', 'Battery Life', 'Up to 22 hours', 1),
    (3, 'warranty', 'AppleCare', 'Optional extended warranty', 1);