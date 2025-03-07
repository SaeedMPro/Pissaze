-- ============================================================
-- 		DATASET SCRIPT FOR PISSAZE SYSTEM DATABASE
-- ============================================================

-- Connect to database (if not already connected)
-- \c pissaze_system

-- ============================
-- 1. Insert Products
-- ============================
INSERT INTO product (brand, model, current_price, stock_count, category) VALUES
('Intel', 'i7-9700K', 300, 10, 'CPU'),
('AMD', 'Ryzen 7 3700X', 280, 15, 'CPU'),
('Intel', 'i9-10900K', 500, 8, 'CPU'),
('AMD', 'Ryzen 9 3950X', 700, 6, 'CPU'),
('CoolerMaster', 'Hyper 212', 50, 15, 'Cooler'),
('Noctua', 'NH-D15', 90, 8, 'Cooler'),
('Deepcool', 'Gammaxx 400', 45, 20, 'Cooler'),
('Seagate', 'Barracuda 2TB', 60, 20, 'HDD'),
('Western Digital', 'WD Blue 1TB', 50, 25, 'HDD'),
('Toshiba', 'X300 4TB', 80, 12, 'HDD'),
('ASUS', 'Z390-A', 150, 5, 'Motherboard'),
('MSI', 'B450 Tomahawk', 120, 10, 'Motherboard'),
('Gigabyte', 'Aorus Elite', 140, 6, 'Motherboard'),
('ASRock', 'B550 Pro4', 130, 9, 'Motherboard'),
('MSI', 'X570-A Pro', 180, 4, 'Motherboard'),
('Corsair', 'Vengeance 16GB', 80, 25, 'RAM Stick'),
('G.Skill', 'Trident Z 16GB', 85, 30, 'RAM Stick'),
('Kingston', 'HyperX Fury 16GB', 75, 28, 'RAM Stick'),
('Patriot', 'Viper 16GB', 70, 30, 'RAM Stick'),
('NVIDIA', 'RTX 3080', 700, 5, 'GPU'),
('AMD', 'Radeon RX 6800', 650, 6, 'GPU'),
('NVIDIA', 'RTX 3070', 500, 7, 'GPU'),
('Samsung', '970 EVO 1TB', 120, 10, 'SSD'),
('Kingston', 'A2000 1TB', 110, 12, 'SSD'),
('Crucial', 'MX500 1TB', 100, 15, 'SSD'),
('EVGA', '750W Gold', 100, 8, 'Power Supply'),
('Corsair', 'RM850x', 130, 7, 'Power Supply'),
('Seasonic', 'Focus 650W', 90, 10, 'Power Supply'),
('NZXT', 'H510', 70, 12, 'Case'),
('Phanteks', 'Eclipse P400A', 80, 10, 'Case'),
('CoolerMaster', 'MasterCase H500', 90, 9, 'Case'),
('Lian Li', 'PC-O11 Dynamic', 150, 5, 'Case'),
('Thermaltake', 'Core V71', 140, 6, 'Case');

INSERT INTO product_cpu VALUES 
(1, 'نسل ۹', 'Coffee Lake', 8, 8, 3.6, 4.9, 128, 95),
(2, 'Zen 2', 'Matisse', 8, 16, 3.6, 4.4, 128, 65),
(3, 'نسل ۱۰', 'Comet Lake', 10, 10, 3.7, 5.3, 128, 125),
(4, 'Zen 2', 'Matisse', 8, 16, 3.8, 4.5, 128, 95);

INSERT INTO product_cooler VALUES 
(5, 'air', 120, 2000, 5, 12.0, 15.8, 8.5),
(6, 'air', 140, 1800, 5, 11.0, 14.8, 7.5),
(7, 'air', 110, 2100, 5, 11.5, 15.0, 8.0);

INSERT INTO product_hdd VALUES 
(8, 2.00, 7200, 6, 14.7, 2.6, 10.2),
(9, 1.00, 5400, 5, 13.0, 2.5, 9.0),
(10, 4.00, 7200, 8, 16.0, 3.0, 12.0);

INSERT INTO product_motherboard VALUES 
(11, 'Z390', 4, 2666.00, 30, 30.5, 24.4, 24.4),
(12, 'B450', 4, 2933.00, 25, 28.0, 20.0, 21.0),
(13, 'Aorus Elite', 4, 3200.00, 35, 32.0, 25.0, 26.0),
(14, 'B550 Pro4', 4, 3600.00, 30, 30.0, 22.0, 23.0),
(15, 'X570 Aorus Elite', 4, 4000.00, 4, 33.0, 26.0, 27.0),
(16, 'Z490 AORUS', 4, 3000.00, 30, 31.0, 25.0, 26.0);

INSERT INTO product_ram_stick VALUES 
(17, 'DDR4', 16.00, 3200.00, 3, 13.3, 3.1, 5.5),
(18, 'DDR4', 16.00, 3600.00, 3, 13.3, 3.1, 5.5),
(19, 'DDR4', 16.00, 3000.00, 3, 13.0, 3.0, 5.0);

INSERT INTO product_gpu VALUES 
(20, 10, 1.71, 3, 320, 28.5, 11.2, 5.0),
(21, 8, 1.60, 2, 300, 26.0, 10.5, 4.5),
(22, 12, 1.80, 4, 350, 30.0, 12.0, 6.0);

INSERT INTO product_ssd VALUES 
(23, 1.00, 5),
(24, 1.00, 5),
(25, 2.00, 5);

INSERT INTO product_power_supply VALUES 
(26, 750, 15.0, 8.6, 15.0),
(27, 850, 16.0, 9.0, 16.0),
(28, 650, 14.0, 8.0, 14.0);

INSERT INTO product_case VALUES 
(29, 'میان برج', 'مشکی', 'فولاد', 120, 2, 0, 43.5, 20.0, 40.0),
(30, 'کل برج', 'سفید', 'آلومینیوم', 140, 3, 0, 50.0, 25.0, 45.0),
(31, 'مایکروم', 'مشکی', 'فلز', 110, 2, 0, 40.0, 18.0, 38.0),
(32, 'پراید', 'طوسی', 'فلز', 130, 2, 0, 45.0, 22.0, 42.0),
(33, 'تِRAM Stickو', 'مشکی', 'فولاد', 125, 2, 0, 42.0, 21.0, 41.0);

INSERT INTO compatible_cc_socket VALUES (1, 5);
INSERT INTO compatible_cc_socket VALUES (2, 6);
INSERT INTO compatible_cc_socket VALUES (3, 7);
INSERT INTO compatible_cc_socket VALUES (4, 5);

INSERT INTO compatible_mc_socket VALUES (1, 11);
INSERT INTO compatible_mc_socket VALUES (2, 12);
INSERT INTO compatible_mc_socket VALUES (1, 13);
INSERT INTO compatible_mc_socket VALUES (2, 14);
INSERT INTO compatible_mc_socket VALUES (1, 15);
INSERT INTO compatible_mc_socket VALUES (2, 16);

INSERT INTO compatible_rm_slot VALUES (17, 11);
INSERT INTO compatible_rm_slot VALUES (17, 12);
INSERT INTO compatible_rm_slot VALUES (18, 13);
INSERT INTO compatible_rm_slot VALUES (19, 14);
INSERT INTO compatible_rm_slot VALUES (18, 15);
INSERT INTO compatible_rm_slot VALUES (19, 16);

INSERT INTO compatible_gp_connector VALUES (20, 26);
INSERT INTO compatible_gp_connector VALUES (21, 27);
INSERT INTO compatible_gp_connector VALUES (22, 27);

INSERT INTO compatible_sm_slot VALUES (23, 11);
INSERT INTO compatible_sm_slot VALUES (24, 12);
INSERT INTO compatible_sm_slot VALUES (25, 13);

INSERT INTO compatible_gm_slot VALUES (20, 13);
INSERT INTO compatible_gm_slot VALUES (21, 14);
INSERT INTO compatible_gm_slot VALUES (22, 15);

-- ============================
-- 4. Insert Clients
-- ============================
INSERT INTO client (phone_number, first_name, last_name, referral_code) VALUES
('09187997434', 'فرزین', 'همزه ئی', 'JOHN123'),
('09398113791', 'دانیال', 'کشاورز نژاد', 'JANE456'),
('09368815915', 'سعید', 'مظاهری', 'BOB789'),
('09185552233', 'علی', 'علیزاد', 'ALICE321'),
('09361112233', 'امیر', 'موسوی', 'TOM654'),
('09907888770', 'شایان', 'محمدیان', 'LISA987'),
('09156664422', 'محمد', 'غضنفرپور', 'MARK111'),
('05555555555', 'محمد', 'درویشی', 'ANNA222'),
('04444444444', 'رضا', 'محمدی', 'JAMES333'),
('03333333333', 'علیرضا', 'رضایی', 'SARA444'),
('07777777777', 'علی', 'امیری', 'KEVIN555'),
('08888888888', 'امیرمحمد', 'موسوی', 'EMILY666');

INSERT INTO vip_client (client_id, expiration_time) VALUES
(1, NOW() + INTERVAL '1 month'),
(3, NOW() + INTERVAL '1 months'),
(5, NOW() + INTERVAL '1 month'),
(7, NOW() + INTERVAL '1 days');

INSERT INTO address_of_client (client_id, province, remain_address) VALUES
(1, 'تهران', 'خیابان آزادی، پلاک ۱۲۳'),
(1, 'تهران', 'خیابان انقلاب، پلاک ۴۵۶'),
(1, 'اصفهان', 'خیابان نقش جهان، پلاک ۷۸۹'),
(1, 'شیراز', 'خیابان ولیعصر، پلاک ۱۰۱'),
(4, 'مشهد', 'خیابان رضوی، پلاک ۲۰۲'),
(5, 'تبریز', 'خیابان امام، پلاک ۳۰۳'),
(6, 'کرمان', 'خیابان ولیعصر، پلاک 404'),
(7, 'قم', 'خیابان معلم، پلاک 505'),
(8, 'رشت', 'خیابان دماوند، پلاک 606'),
(9, 'اصفهان', 'خیابان چهارباغ، پلاک 707'),
(10, 'تهران', 'خیابان انقلاب، پلاک 808'),
(11, 'شیراز', 'خیابان حافظ، پلاک 909'),
(12, 'مشهد', 'خیابان فردوسی، پلاک 1010');

-- ============================
-- 7. Insert Shopping Carts
-- ============================
INSERT INTO shopping_cart (client_id, cart_number, cart_status) VALUES
(1, 1, 'active'),
(2, 1,'active'),
(3, 1,'locked'),
(4, 1,'locked'),
(5, 1,'locked'),
(6, 1,'active'),
(7, 1,'active'),
(8, 1,'locked'),
(9, 1,'active'),
(10, 1,'active'),
(11, 1,'active'),
(12, 1,'active');

INSERT INTO shopping_cart (client_id, cart_number, cart_status) VALUES
(1, 2,'locked'),
(1, 3, 'locked'),
(1, 4, 'blocked'),
(3, 2,'active'),
(5, 2,'active'),
(7, 2,'locked');

INSERT INTO locked_shopping_cart (client_id, cart_number, locked_number) VALUES
(1, 2, 1),
(1, 3, 2),
(1, 3, 3),
(3, 1, 1),
(4, 1, 1),
(5, 1, 1),
(7, 2, 1),
(8, 1, 1);

-- ============================
-- 9. Insert Discount Codes
-- ============================
INSERT INTO discount_code (amount, discount_limit, expiration_time, code_type) VALUES
(0.10, 50000.00, NOW() + INTERVAL '1 week', 'public'),
(0.15, 60000.00, NOW() + INTERVAL '2 weeks', 'public'),
(0.20, 70000.00, NOW() + INTERVAL '3 weeks', 'public');

INSERT INTO discount_code (amount, discount_limit, expiration_time, code_type) VALUES
(50000.00, 50000.00, NOW() + INTERVAL '1 week', 'private'),
(60000.00, 60000.00, NOW() + INTERVAL '1 week', 'private'),
(70000.00, 70000.00, NOW() + INTERVAL '2 weeks', 'private');

INSERT INTO private_code (code, client_id) VALUES
(4, 1),
(6, 1),
(5, 2);

-- ============================
-- 11. Insert Transactions
-- ============================
INSERT INTO transaction (tracking_code, transaction_status, transaction_type) VALUES
(1001, 'successful', 'wallet'),
(1002, 'successful', 'bank'),
(1003, 'semi-successful', 'wallet'),
(1004, 'successful', 'wallet'),
(1005, 'unsuccessful', 'bank'),
(1006, 'successful', 'wallet'),
(1007, 'successful', 'wallet'),
(1008, 'successful', 'bank'),
(1009, 'successful', 'wallet'),
(1010, 'successful', 'wallet');

INSERT INTO bank_transaction (tracking_code, card_number) VALUES
(1002, '6037697479741990'),
(1008, '5859831217628220');

INSERT INTO deposit_wallet (tracking_code, client_id, amount) VALUES
(1001, 1, 100000.00),
(1004, 3, 50000.00),
(1006, 7, 75000.00),
(1007, 8, 60000.00),
(1009, 11, 85000.00),
(1010, 12, 90000.00);

-- ============================
-- 14. Insert Shopping Cart Contents (added_to)
-- ============================
INSERT INTO added_to (cart_number, client_id, locked_number, product_id, quantity, cart_price) VALUES
(2, 1, 1, 1, 1, 300.00),      -- John buys a CPU (Intel i7-9700K)
(2, 1, 1, 6, 1, 700.00),      -- John buys a GPU (NVIDIA RTX 3080)
(3, 1, 2, 2, 1, 280.00),      -- Jane buys a CPU (AMD Ryzen 7 3700X)
(3, 1, 2, 11, 1, 700.00),     -- Jane buys a GPU (NVIDIA RTX 3080)
(1, 3, 1, 4, 1, 90.00),       -- Bob buys a Cooler (Noctua NH-D15)
(1, 3, 1, 9, 2, 80.00);       -- Bob buys 2 RAM sticks (Corsair Vengeance 16GB)

INSERT INTO applied_to (cart_number, client_id, locked_number, discount_code, time_stamp) VALUES
(2, 1, 1, 1, NOW()),
(2, 1, 1, 2, NOW());

-- ============================
-- 16. Insert Issued Orders (issued_for)
-- ============================
INSERT INTO issued_for (tracking_code, cart_number, client_id, locked_number) VALUES
(1001, 2, 1, 1);

-- ============================
-- 17. Insert Referral Relationships (refers)
-- ============================
INSERT INTO refers (referee_id, referrer_id) VALUES
('JANE456', 'JOHN123'),
('BOB789', 'JANE456'),
('ALICE321', 'BOB789'),
('TOM654', 'ALICE321'),
('LISA987', 'TOM654'),
('MARK111', 'LISA987'),
('ANNA222', 'MARK111'),
('JAMES333', 'ANNA222'),
('SARA444', 'JAMES333'),
('KEVIN555', 'SARA444'),
('EMILY666', 'KEVIN555');

-- ============================================================
-- End of Dataset
-- ============================================================
