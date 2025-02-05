--POSTGRESQL--
--PISSAZE SYSTEM DATABASE--

-- Create the database
CREATE DATABASE pissaze_system;

-- Connect to the database
\c pissaze_system

-- Create ENUM types
CREATE TYPE cooling_method_enum AS ENUM ('liquid', 'air');
CREATE TYPE discount_enum AS ENUM ('public', 'private');
CREATE TYPE transaction_status_enum AS ENUM ('Successful', 'semi-successful', 'unsuccessful');
CREATE TYPE transaction_type_enum AS ENUM ('bank', 'wallet');
CREATE TYPE cart_status_enum AS ENUM ('locked', 'blocked', 'active');

-- Table Definitions --

CREATE TABLE product (
    id              INT PRIMARY KEY, 
    brand           VARCHAR(50) NOT NULL,
    model           VARCHAR(50) NOT NULL,
    current_price   INT,
    stock_count     SMALLINT,
    category        VARCHAR(50),
    product_image   BYTEA
);

CREATE TABLE product_hdd (
    product_id          SERIAL PRIMARY KEY, 
    capacity            DECIMAL(5, 2),          
    rotational_speed    INT,  
    wattage             INT,           
    depth               DECIMAL(5, 2), 
    height              DECIMAL(5, 2),    
    width               DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_cooler (
    product_id              SERIAL PRIMARY KEY, 
    cooling_method          cooling_method_enum,
    fan_size                INT,              
    max_rotational_speed    INT,  
    wattage                 INT,               
    depth                   DECIMAL(5, 2), 
    height                  DECIMAL(5, 2),    
    width                   DECIMAL(5, 2),    
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_cpu (
    product_id          SERIAL PRIMARY KEY, 
    generation          VARCHAR(50),
    microarchitecture   VARCHAR(50),
    num_cores           SMALLINT,
    num_threads         SMALLINT,
    base_frequency      DECIMAL(5, 2), 
    boost_frequency     DECIMAL(5, 2),
    max_memory_limit    INT,         
    wattage             INT,                
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_ram_stick (
    product_id  SERIAL PRIMARY KEY, 
    generation  VARCHAR(50),
    capacity    DECIMAL(5, 2),    
    frequency   DECIMAL(5, 2),   
    wattage     INT, 
    depth       DECIMAL(5, 2), 
    height      DECIMAL(5, 2),    
    width       DECIMAL(5, 2),   
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_case (
    product_id      SERIAL PRIMARY KEY, 
    product_type    VARCHAR(50),
    color           VARCHAR(50),
    material        VARCHAR(50),
    fan_size        INT,         
    num_fans        SMALLINT,
    wattage         INT,
    depth           DECIMAL(5, 2), 
    height          DECIMAL(5, 2),    
    width           DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_power_supply (
    product_id          SERIAL PRIMARY KEY, 
    supported_wattage   INT,
    depth               DECIMAL(5, 2), 
    height              DECIMAL(5, 2),    
    width               DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_gpu (
    product_id  SERIAL PRIMARY KEY, 
    ram_size    INT,         
    clock_speed DECIMAL(5, 2), 
    num_fans    SMALLINT,
    wattage     INT,
    depth       DECIMAL(5, 2), 
    height      DECIMAL(5, 2),    
    width       DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_ssd (
    product_id  SERIAL PRIMARY KEY, 
    capacity    DECIMAL(5, 2), 
    wattage     INT,
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_motherboard (
    product_id          SERIAL PRIMARY KEY, 
    chipset_name        VARCHAR(50),
    num_memory_slots    SMALLINT,
    memory_speed_range  DECIMAL(5, 2),
    wattage             INT,
    depth               DECIMAL(5, 2), 
    height              DECIMAL(5, 2),    
    width               DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_cc_socket (
    cpu_id      INT NOT NULL, 
    cooler_id   INT NOT NULL, 
    FOREIGN KEY (cooler_id) REFERENCES product_cooler (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (cpu_id) REFERENCES product_cpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_mc_socket (
    cpu_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (cpu_id) REFERENCES product_cpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_rm_slot (
    ram_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ram_id) REFERENCES product_ram_stick (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_gp_connector (
    gpu_id          INT NOT NULL, 
    power_supply_id INT NOT NULL, 
    FOREIGN KEY (gpu_id) REFERENCES product_gpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (power_supply_id) REFERENCES product_power_supply (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_sm_slot (
    ssd_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ssd_id) REFERENCES product_ssd (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_gm_slot (
    gpu_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (gpu_id) REFERENCES product_gpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE client (
    client_id          SERIAL PRIMARY KEY, 
    phone_number       VARCHAR(15) NOT NULL UNIQUE,
    first_name         VARCHAR(50) NOT NULL,
    last_name          VARCHAR(50) NOT NULL,
    wallet_balance     DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    time_stamp         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    referral_code      VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE vip_client (
    client_id       INT PRIMARY KEY, 
    expiration_time TIMESTAMP NOT NULL,
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE address_of_client (
    client_id       INT NOT NULL, 
    province        VARCHAR(20) NOT NULL,
    remain_address  VARCHAR(255) NOT NULL,
    PRIMARY KEY (client_id, province, remain_address),
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE shopping_cart (
    cart_number    SERIAL NOT NULL, 
    client_id      INT NOT NULL,
    cart_status    cart_status_enum NOT NULL,
    PRIMARY KEY (client_id, cart_number),
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE discount_code (
    code            INT PRIMARY KEY, 
    amount          DECIMAL(10, 2) CHECK (amount > 0),
    discount_limit  DECIMAL(5, 2) CHECK (discount_limit > 0),
    usage_limit     SMALLINT DEFAULT 1 CHECK (usage_limit >= 0) , 
    expiration_time TIMESTAMP,
    code_type       discount_enum NOT NULL
);

--I decide to delete public_code and instead add `type` in discount_code

CREATE TABLE private_code (
    code        INT PRIMARY KEY, 
    client_id   INT NOT NULL,
    time_stamp  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (code) REFERENCES discount_code (code) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE transaction (
    tracking_code       INT PRIMARY KEY, 
    transaction_status  transaction_status_enum NOT NULL,
    transaction_type    transaction_type_enum NOT NULL,
    time_stamp          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--I decide to delete wallet_transaction and instead add `type` in transaction

CREATE TABLE bank_transaction (
    tracking_code   INT PRIMARY KEY, 
    card_number     INT NOT NULL,
    FOREIGN KEY (tracking_code) REFERENCES transaction (tracking_code) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE locked_shopping_cart (
    cart_number     INT NOT NULL, 
    client_id       INT NOT NULL,
    locked_number   SERIAL NOT NULL,
    PRIMARY KEY (client_id, cart_number, locked_number),
    FOREIGN KEY (client_id, cart_number) REFERENCES shopping_cart (client_id, cart_number) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE deposit_wallet (
    tracking_code   INT PRIMARY KEY, 
    client_id       INT NOT NULL,
    amount          DECIMAL(12, 2) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (tracking_code) REFERENCES transaction (tracking_code) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE subscribes (
    tracking_code   INT PRIMARY KEY, 
    client_id       INT NOT NULL,
    FOREIGN KEY (client_id) REFERENCES client (client_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (tracking_code) REFERENCES transaction (tracking_code) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE refers (
    referee_id  VARCHAR(20) PRIMARY KEY, 
    referrer_id VARCHAR(20) NOT NULL,
    FOREIGN KEY (referee_id) REFERENCES client (referral_code) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (referrer_id) REFERENCES client (referral_code) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE added_to (
    cart_number     INT NOT NULL, 
    client_id       INT NOT NULL,
    locked_number   INT NOT NULL,
    product_id      INT NOT NULL, 
    quantity        SMALLINT CHECK (quantity > 0),
    cart_price      DECIMAL(12, 2) CHECK (cart_price >= 0),
    PRIMARY KEY (client_id, cart_number, locked_number, product_id),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (client_id, cart_number, locked_number) REFERENCES locked_shopping_cart (client_id, cart_number, locked_number) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE applied_to (
    cart_number     INT NOT NULL, 
    client_id       INT NOT NULL,
    locked_number   INT NOT NULL,
    discount_code   INT NOT NULL, 
    time_stamp      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (client_id, cart_number, locked_number, discount_code),
    FOREIGN KEY (discount_code) REFERENCES discount_code (code) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (client_id, cart_number, locked_number) REFERENCES locked_shopping_cart (client_id, cart_number, locked_number) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE issued_for (
    tracking_code   INT PRIMARY KEY,
    cart_number     INT NOT NULL, 
    client_id       INT NOT NULL,
    locked_number   INT NOT NULL,
    FOREIGN KEY (client_id, cart_number, locked_number) REFERENCES locked_shopping_cart (client_id, cart_number, locked_number) ON UPDATE CASCADE ON DELETE CASCADE
);

-- triggers and events --

/*
Ensures that: 
    If a user refers another user, a discount code appropriate to their position in the referral chain
    should be gifted to all users in that chain.
*/
CREATE OR REPLACE FUNCTION handle_referral() 
RETURNS TRIGGER AS $$
DECLARE
    referrer_id         VARCHAR(20);
    referee_id          VARCHAR(20) := NEW.referee_id;
    current_level       INT := 0;
    discount_percentage DECIMAL(10, 2);
    new_discount_code   INT;
    client_id           INT;
BEGIN
    SELECT r.referrer_id INTO referrer_id
    FROM refers r
    WHERE r.referee_id = referee_id;

    -- Loop through the referral chain
    WHILE referrer_id IS NOT NULL LOOP
        CASE current_level
            WHEN 0 THEN discount_percentage := 50;
            ELSE discount_percentage := 50 / (2 * current_level);
        END CASE;

        SELECT c.client_id INTO client_id
        FROM client c
        WHERE referee_id = referral_code;

        IF discount_percentage < 1 THEN
            --fixed discount (discount_percentage < 1%) 
            INSERT INTO discount_code (code, amount, discount_limit, expiration_time, code_type)
            VALUES (nextval('discount_code_code_seq'), 50000, 50000, NOW() + INTERVAL '1 week', 'private')
            RETURNING code INTO new_discount_code;
            --usage_limit (default) = 1
        ELSE 
            discount_percentage := discount_percentage / 100 ;
            INSERT INTO discount_code (code, amount, discount_limit, expiration_time, code_type)
            VALUES (nextval('discount_code_code_seq'), discount_percentage, 1000000, NOW() + INTERVAL '1 week', 'private')
            RETURNING code INTO new_discount_code;
            --usage_limit (default) = 1
        END IF;

        INSERT INTO private_code (code, client_id, time_stamp)
        VALUES (new_discount_code, client_id, NOW());

        SELECT r.referrer_id , r.referee_id 
        INTO referrer_id, referee_id
        FROM refers r
        WHERE r.referee_id = referrer_id;

        current_level := current_level + 1;
    END LOOP;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER referral_trigger
AFTER INSERT ON refers
FOR EACH ROW
EXECUTE FUNCTION handle_referral();

/*
Ensures that:
    blocked cart should not be accessible or registered.
*/
CREATE OR REPLACE FUNCTION check_blocked_cart()
RETURNS TRIGGER AS $$
DECLARE
    cart_status cart_status_enum;
BEGIN
    SELECT cart_status
    INTO cart_status
    FROM locked_shopping_cart NATURAL JOIN shopping_cart
    WHERE cart_number = NEW.cart_number;

    IF cart_status = 'blocked'THEN
        RAISE EXCEPTION 'Operation not allowed: Cart % is blocked.', NEW.cart_number;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_adding_blocked_cart_trigger
BEFORE INSERT OR UPDATE ON added_to
FOR EACH ROW
EXECUTE FUNCTION check_blocked_cart();

CREATE TRIGGER prevent_issued_for_blocked_cart_trigger
BEFORE INSERT OR UPDATE ON issued_for
FOR EACH ROW
EXECUTE FUNCTION check_blocked_cart();

CREATE TRIGGER prevent_applied_to_blocked_cart_trigger
BEFORE INSERT OR UPDATE ON applied_to
FOR EACH ROW
EXECUTE FUNCTION check_blocked_cart();


/*
Ensures that:
    No user should be able to add a product to their cart that is out of stock.
*/
CREATE OR REPLACE FUNCTION check_product_stock()
RETURNS TRIGGER AS $$
DECLARE
    stock_count SMALLINT;
BEGIN
    SELECT p.stock_count INTO stock_count
    FROM product p
    WHERE p.id = NEW.product_id;

    IF stock_count < NEW.quantity THEN
        RAISE EXCEPTION 
        'Not enough stock available for product_id: %', NEW.product_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_out_of_stock_trigger
BEFORE INSERT OR UPDATE ON added_to
FOR EACH ROW
EXECUTE FUNCTION check_product_stock();

/*
Ensures that: 
    Adding a product to the cart should reduce its stock count in the inventory.
*/
CREATE OR REPLACE FUNCTION reduce_stock_after_add_to_cart()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE product
    SET stock_count = stock_count - NEW.quantity
    WHERE id = NEW.product_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER reduce_stock_trigger
AFTER INSERT ON added_to
FOR EACH ROW
EXECUTE FUNCTION reduce_stock_after_add_to_cart();


/*
Ensures that:
    Registered users should have access to only one shopping cart,
    and VIP users should have access to up to five shopping carts
*/
CREATE OR REPLACE FUNCTION enforce_cart_limit()
RETURNS TRIGGER AS $$
DECLARE
    cart_count INT;
    is_vip BOOLEAN;
BEGIN
    
    SELECT EXISTS (
        SELECT 1
        FROM vip_client
        WHERE client_id = NEW.client_id
    ) INTO is_vip;

    SELECT COUNT(*) INTO cart_count
    FROM shopping_cart
    WHERE client_id = NEW.client_id;

    IF is_vip THEN
        IF cart_count >= 5 THEN
            RAISE EXCEPTION 'vip users cannot have more than five shopping carts.';
        END IF;
    ELSE
        IF cart_count >= 1 THEN
            RAISE EXCEPTION 'Registered users cannot have more than one shopping cart.';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER enforce_cart_limit_trigger
BEFORE INSERT ON shopping_cart
FOR EACH ROW
EXECUTE FUNCTION enforce_cart_limit();


/*
Ensures that: 
    1.No user shall use a discount code more than the number of times it is allowed.
    2.No user shall use an expired discount code. 
*/
CREATE OR REPLACE FUNCTION enforce_apply_discount()
RETURNS TRIGGER AS $$
DECLARE
    code_record RECORD;
BEGIN
    SELECT d.usage_count, d.usage_limit, d.expiration_time
    INTO code_record
    FROM discount_code d
    WHERE d.code = NEW.code;

    IF code_record IS NULL THEN
        RAISE EXCEPTION 'Invalid discount code.';
    END IF;

    IF code_record.expiration_time < NOW() THEN
        RAISE EXCEPTION 'The discount code has expired.';
    END IF;

    IF code_record.usage_count >= code_record.usage_limit THEN
        RAISE EXCEPTION 'The usage limit for this discount code has been reached.';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_expiration_limit_discount_trigger
BEFORE INSERT ON applied_to
FOR EACH ROW
EXECUTE FUNCTION enforce_apply_discount();


/*
Ensures that: 
    With a deposit into the wallet, the wallet balance of the client increases. 
*/
CREATE OR REPLACE FUNCTION deposit()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE client
    SET wallet_balance = wallet_balance + amount
    WHERE client_id = NEW.client_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER deposit_into_wallet_trigger
AFTER INSERT ON deposit_wallet
FOR EACH ROW
EXECUTE FUNCTION deposit();


/*
Ensures that: 
    By purchasing product or subscriptions using a digital wallet,
    the wallet balance of the client  decreases.
*/
CREATE OR REPLACE FUNCTION reduce_wallet()
RETURNS TRIGGER AS $$
DECLARE
    transaction_type transaction_type_enum;
BEGIN

    SELECT transaction_type
    INTO transaction_type
    FROM transaction
    WHERE tracking_code = NEW.tracking_code;

    IF transaction_type = 'wallet' THEN
        UPDATE client
        SET wallet_balance = wallet_balance - amount
        WHERE client_id = NEW.client_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER reduce_user_wallet_trigger
AFTER INSERT ON issued_for
FOR EACH ROW
EXECUTE FUNCTION reduce_wallet();

CREATE TRIGGER reduce_user_wallet_trigger
AFTER INSERT ON subscribes
FOR EACH ROW
EXECUTE FUNCTION reduce_wallet();


/*
Ensures that:
    when an order is finalized and paid, the associated cart is unlocked.
*/
CREATE OR REPLACE FUNCTION unlock_cart_after_payment()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE shopping_cart
    SET cart_status = 'active'
    WHERE cart_number = NEW.cart_number;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER unlock_cart_trigger
AFTER INSERT ON issued_for
FOR EACH ROW
EXECUTE FUNCTION unlock_cart_after_payment();
