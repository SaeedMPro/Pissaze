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
('CoolerMaster', 'Hyper 212', 50, 15, 'Cooler'),
('Noctua', 'NH-D15', 90, 8, 'Cooler'),
('Seagate', 'Barracuda 2TB', 60, 20, 'HDD'),
('Western Digital', 'WD Blue 1TB', 50, 25, 'HDD'),
('ASUS', 'Z390-A', 150, 5, 'Motherboard'),
('MSI', 'B450 Tomahawk', 120, 10, 'Motherboard'),
('Corsair', 'Vengeance 16GB', 80, 25, 'RAM'),
('G.Skill', 'Trident Z 16GB', 85, 30, 'RAM'),
('NVIDIA', 'RTX 3080', 700, 5, 'GPU'),
('AMD', 'Radeon RX 6800', 650, 6, 'GPU'),
('Samsung', '970 EVO 1TB', 120, 10, 'SSD'),
('Kingston', 'A2000 1TB', 110, 12, 'SSD'),
('EVGA', '750W Gold', 100, 8, 'Power Supply'),
('Corsair', 'RM850x', 130, 7, 'Power Supply'),
('NZXT', 'H510', 70, 12, 'Case'),
('Phanteks', 'Eclipse P400A', 80, 10, 'Case'),
('Gigabyte', 'Aorus Elite', 140, 6, 'Motherboard'),
('ASRock', 'B550 Pro4', 130, 9, 'Motherboard');

-- ============================
-- 2. Insert Product Details
-- ============================
-- CPUs
INSERT INTO product_cpu VALUES 
(1, '9th Gen', 'Coffee Lake', 8, 8, 3.6, 4.9, 128, 95),
(2, 'Zen 2', 'Matisse', 8, 16, 3.6, 4.4, 128, 65);

-- Coolers
INSERT INTO product_cooler VALUES 
(3, 'air', 120, 2000, 5, 12.0, 15.8, 8.5),
(4, 'air', 140, 1800, 5, 11.0, 14.8, 7.5);

-- HDDs
INSERT INTO product_hdd VALUES 
(5, 2.00, 7200, 6, 14.7, 2.6, 10.2),
(6, 1.00, 5400, 5, 13.0, 2.5, 9.0);

-- Motherboards
INSERT INTO product_motherboard VALUES 
(7, 'Z390', 4, 2666.00, 30, 30.5, 24.4, 24.4),
(8, 'B450', 4, 2933.00, 25, 28.0, 20.0, 21.0),
(19, 'Aorus Elite', 4, 3200.00, 35, 32.0, 25.0, 26.0),
(20, 'B550 Pro4', 4, 3600.00, 30, 30.0, 22.0, 23.0);

-- RAM sticks
INSERT INTO product_ram_stick VALUES 
(9, 'DDR4', 16.00, 3200.00, 3, 13.3, 3.1, 5.5),
(10, 'DDR4', 16.00, 3600.00, 3, 13.3, 3.1, 5.5);

-- GPUs
INSERT INTO product_gpu VALUES 
(11, 10, 1.71, 3, 320, 28.5, 11.2, 5.0),
(12, 8, 1.60, 2, 300, 26.0, 10.5, 4.5);

-- SSDs
INSERT INTO product_ssd VALUES 
(13, 1.00, 5),
(14, 1.00, 5);

-- Power Supplies
INSERT INTO product_power_supply VALUES 
(15, 750, 15.0, 8.6, 15.0),
(16, 850, 16.0, 9.0, 16.0);

-- Cases
INSERT INTO product_case VALUES 
(17, 'Mid Tower', 'Black', 'Steel', 120, 2, 0, 43.5, 20.0, 40.0),
(18, 'Full Tower', 'White', 'Aluminum', 140, 3, 0, 50.0, 25.0, 45.0);

-- ============================
-- 3. Insert Compatibility Records
-- ============================
-- CPU-Cooler
INSERT INTO compatible_cc_socket VALUES (1, 3);
INSERT INTO compatible_cc_socket VALUES (2, 4);

-- CPU-Motherboard
INSERT INTO compatible_mc_socket VALUES (1, 7);
INSERT INTO compatible_mc_socket VALUES (2, 8);
INSERT INTO compatible_mc_socket VALUES (1, 19);
INSERT INTO compatible_mc_socket VALUES (2, 20);

-- RAM-Motherboard
INSERT INTO compatible_rm_slot VALUES (9, 7);
INSERT INTO compatible_rm_slot VALUES (9, 8);
INSERT INTO compatible_rm_slot VALUES (10, 19);
INSERT INTO compatible_rm_slot VALUES (10, 20);

-- GPU-PowerSupply
INSERT INTO compatible_gp_connector VALUES (11, 15);
INSERT INTO compatible_gp_connector VALUES (12, 16);

-- SSD-Motherboard
INSERT INTO compatible_sm_slot VALUES (13, 7);
INSERT INTO compatible_sm_slot VALUES (14, 8);

-- GPU-Motherboard
INSERT INTO compatible_gm_slot VALUES (11, 19);
INSERT INTO compatible_gm_slot VALUES (12, 20);

-- ============================
-- 4. Insert Clients
-- ============================
INSERT INTO client (phone_number, first_name, last_name, referral_code) VALUES
('+1234567890', 'John', 'Doe', 'JOHN123'),
('+0987654321', 'Jane', 'Smith', 'JANE456'),
('+1122334455', 'Bob', 'Brown', 'BOB789'),
('+1029384756', 'Alice', 'Green', 'ALICE321'),
('+5647382910', 'Tom', 'White', 'TOM654'),
('+0192837465', 'Lisa', 'Black', 'LISA987'),
('+0918273645', 'Mark', 'Blue', 'MARK111'),
('+5555555555', 'Anna', 'Red', 'ANNA222'),
('+4444444444', 'James', 'Yellow', 'JAMES333'),
('+3333333333', 'Sara', 'Purple', 'SARA444'),
('+7777777777', 'Kevin', 'Orange', 'KEVIN555'),
('+8888888888', 'Emily', 'Pink', 'EMILY666');

-- ============================
-- 5. Insert VIP Clients
-- ============================
INSERT INTO vip_client (client_id, expiration_time) VALUES
(1, NOW() + INTERVAL '1 month'),
(3, NOW() + INTERVAL '1 months'),
(5, NOW() + INTERVAL '1 month'),
(7, NOW() + INTERVAL '1 days');

-- ============================
-- 6. Insert Addresses
-- ============================
INSERT INTO address_of_client (client_id, province, remain_address) VALUES
(1, 'California', '123 Tech Street'),
(1, 'asdas', '123 Tech Street'),
(1, 'sadads', 'asdasda'),
(1, 'sadadsd', 'asfasdfasf'),
(2, 'New York', '456 Main Ave'),
(3, 'Texas', '789 Commerce Blvd'),
(4, 'Florida', '321 Sunshine Road'),
(5, 'Washington', '654 Evergreen Terrace'),
(6, 'Oregon', '987 Pine Street'),
(7, 'Nevada', '159 Desert Drive'),
(8, 'Colorado', '753 Mountain Ave'),
(9, 'Arizona', '852 Cactus Rd'),
(10, 'Michigan', '951 Motor City Blvd'),
(11, 'Illinois', '333 Windy Lane'),
(12, 'Georgia', '444 Peach Street');

-- ============================
-- 7. Insert Shopping Carts
-- ============================
-- Each client gets at least one active cart.
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

-- Additional carts for some clients (VIPs and active users)
INSERT INTO shopping_cart (client_id, cart_number, cart_status) VALUES
(1, 2,'locked'),
(3, 2,'active'),
(5, 2,'active'),
(7, 2,'locked');

-- ============================
-- 8. Insert Locked Carts
-- ============================
-- Simulate some carts that are locked.
INSERT INTO locked_shopping_cart (client_id, cart_number, locked_number) VALUES
(1, 2, 1),
(1, 2, 2),
(3, 1, 1),
(4, 1, 1),
(5, 1, 1),
(7, 2, 1),
(8, 1, 1);

-- ============================
-- 9. Insert Discount Codes
-- ============================
-- Public discount codes (percentage-based)
INSERT INTO discount_code (amount, discount_limit, expiration_time, code_type) VALUES
(0.10, 50000.00, NOW() + INTERVAL '1 week', 'public'),
(0.15, 60000.00, NOW() + INTERVAL '2 weeks', 'public'),
(0.20, 70000.00, NOW() + INTERVAL '3 weeks', 'public');

-- Private discount codes (fixed-amount)
INSERT INTO discount_code (amount, discount_limit, expiration_time, code_type) VALUES
(50000.00, 50000.00, NOW() + INTERVAL '1 week', 'private'),
(60000.00, 60000.00, NOW() + INTERVAL '1 week', 'private'),
(70000.00, 70000.00, NOW() + INTERVAL '2 weeks', 'private');

-- ============================
-- 10. Insert Private Codes
-- ============================
INSERT INTO private_code (code, client_id) VALUES
(4, 1),
(3, 1),
(2, 1),
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

-- ============================
-- 12. Insert Bank Transactions
-- ============================
INSERT INTO bank_transaction (tracking_code, card_number) VALUES
(1002, '6037697479741990'),
(1008, '5859831217628220');

-- ============================
-- 13. Insert Deposit Wallets
-- ============================
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
(2, 1, 1, 2, 1, 280.00),      -- Jane buys a CPU (AMD Ryzen 7 3700X)
(2, 1, 1, 11, 1, 700.00),     -- Jane buys a GPU (NVIDIA RTX 3080)
(1, 3, 1, 4, 1, 90.00),       -- Bob buys a Cooler (Noctua NH-D15)
(1, 3, 1, 9, 2, 80.00);       -- Bob buys 2 RAM sticks (Corsair Vengeance 16GB)

-- ============================
-- 15. Insert Applied Discounts (applied_to)
-- ============================
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
