/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for details
-- ----------------------------
DROP TABLE IF EXISTS `details`;
CREATE TABLE `details` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `price` double DEFAULT NULL,
                           `created_time` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of details
-- ----------------------------
BEGIN;
INSERT INTO `details` VALUES (1, 'apple', 1, '2023-05-18 19:43:10');
INSERT INTO `details` VALUES (2, 'pear', 1, '2023-05-18 19:43:45');
INSERT INTO `details` VALUES (3, 'banana', 0.5, '2023-05-18 19:44:08');
COMMIT;

-- ----------------------------
-- Table structure for ratings
-- ----------------------------
DROP TABLE IF EXISTS `ratings`;
CREATE TABLE `ratings` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `product_id` bigint(20) unsigned DEFAULT NULL,
                           `score` int(10) unsigned DEFAULT NULL,
                           `updated_time` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of ratings
-- ----------------------------
BEGIN;
INSERT INTO `ratings` VALUES (1, 1, 5, '2023-05-18 19:48:04');
INSERT INTO `ratings` VALUES (2, 2, 4, '2023-05-18 19:48:21');
INSERT INTO `ratings` VALUES (3, 3, 5, '2023-05-18 19:48:34');
COMMIT;

-- ----------------------------
-- Table structure for reviews
-- ----------------------------
DROP TABLE IF EXISTS `reviews`;
CREATE TABLE `reviews` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `product_id` bigint(20) unsigned DEFAULT NULL,
                           `message` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `created_time` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of reviews
-- ----------------------------
BEGIN;
INSERT INTO `reviews` VALUES (1, 1, 'good', '2023-05-18 19:48:47');
INSERT INTO `reviews` VALUES (2, 2, 'bad', '2023-05-18 19:49:12');
INSERT INTO `reviews` VALUES (3, 3, 'good', '2023-05-18 20:15:56');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
