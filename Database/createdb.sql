--POSTGRESQL--
--PISSAZE SYSTEM DATABASE--

-- Create the database
CREATE DATABASE pissaze_system;

-- Connect to the database
\c pissaze_system

-- Create extension for job scheduled (pg_cron)
/*
   This extension allows for scheduling cron jobs directly within PostgreSQL (for Unix-based systems).
   Ensure pg_cron is installed in the default database (postgres) before using it in other databases.
   After installing, change `shared_preload_libraries = 'pg_cron'` in `postgresql.conf`.
   Also, add `cron.database_name = 'pissaze_system'` to `postgresql.conf`.
*/
CREATE EXTENSION IF NOT EXISTS pg_cron;


-- Create ENUM types
CREATE TYPE cooling_method_enum AS ENUM ('liquid', 'air');
CREATE TYPE discount_enum AS ENUM ('public', 'private');
CREATE TYPE transaction_status_enum AS ENUM ('Successful', 'semi-successful', 'unsuccessful');
CREATE TYPE transaction_type_enum AS ENUM ('bank', 'wallet');
CREATE TYPE cart_status_enum AS ENUM ('locked', 'blocked', 'active');

-- Table Definitions --

CREATE TABLE product (
    id              SERIAL PRIMARY KEY, 
    brand           VARCHAR(50) NOT NULL,
    model           VARCHAR(50) NOT NULL,
    current_price   INT,
    stock_count     SMALLINT,
    category        VARCHAR(50),
    product_image   BYTEA
);

CREATE TABLE product_hdd (
    product_id          INT PRIMARY KEY, 
    capacity            DECIMAL(5, 2),          
    rotational_speed    INT,  
    wattage             INT,           
    depth               DECIMAL(5, 2), 
    height              DECIMAL(5, 2),    
    width               DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_cooler (
    product_id              INT PRIMARY KEY, 
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
    product_id          INT PRIMARY KEY, 
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
    product_id  INT PRIMARY KEY, 
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
    product_id      INT PRIMARY KEY, 
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
    product_id          INT PRIMARY KEY, 
    supported_wattage   INT,
    depth               DECIMAL(5, 2), 
    height              DECIMAL(5, 2),    
    width               DECIMAL(5, 2),
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_gpu (
    product_id  INT PRIMARY KEY, 
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
    product_id  INT PRIMARY KEY, 
    capacity    DECIMAL(5, 2), 
    wattage     INT,
    FOREIGN KEY (product_id) REFERENCES product (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_motherboard (
    product_id          INT PRIMARY KEY, 
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
    PRIMARY KEY (cpu_id, cooler_id), 
    FOREIGN KEY (cooler_id) REFERENCES product_cooler (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (cpu_id) REFERENCES product_cpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_mc_socket (
    cpu_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    PRIMARY KEY (cpu_id, motherboard_id), 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (cpu_id) REFERENCES product_cpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_rm_slot (
    ram_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL, 
    PRIMARY KEY (ram_id, motherboard_id),
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ram_id) REFERENCES product_ram_stick (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_gp_connector (
    gpu_id          INT NOT NULL, 
    power_supply_id INT NOT NULL, 
    PRIMARY KEY (gpu_id, power_supply_id),
    FOREIGN KEY (gpu_id) REFERENCES product_gpu (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (power_supply_id) REFERENCES product_power_supply (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_sm_slot (
    ssd_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL,
    PRIMARY KEY (ssd_id, motherboard_id), 
    FOREIGN KEY (motherboard_id) REFERENCES product_motherboard (product_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (ssd_id) REFERENCES product_ssd (product_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE compatible_gm_slot (
    gpu_id          INT NOT NULL, 
    motherboard_id  INT NOT NULL,
    PRIMARY KEY (gpu_id, motherboard_id), 
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
    code            SERIAL PRIMARY KEY, 
    amount          DECIMAL(12, 2) CHECK (amount > 0),
    discount_limit  DECIMAL(12, 2) CHECK (discount_limit > 0),
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
    time_stamp      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
    quantity        INT CHECK (quantity > 0),
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

-- triggers  --

/*
Ensures that: 
    If a user refers another user, a discount code appropriate to their position in the referral chain
    should be gifted to all users in that chain.
*/
CREATE OR REPLACE FUNCTION handle_referral() 
RETURNS TRIGGER AS $$
DECLARE
    referrer            VARCHAR(20);
    referee             VARCHAR(20) := NEW.referee_id;
    current_level       INT := 1;
    discount_percentage DECIMAL(12, 2);
    new_discount_code   INT;
    client_id_val       INT;
BEGIN
    -- Apply a 50% discount to the **new referee** (first-time user)
    SELECT client_id INTO client_id_val FROM client WHERE referral_code = referee;
    
    INSERT INTO discount_code (code, amount, discount_limit, expiration_time, code_type)
    VALUES (
        nextval('discount_code_code_seq'), 
        0.5,  
        1000000, 
        NOW() + INTERVAL '1 week', 
        'private'
    ) RETURNING code INTO new_discount_code;
    
    INSERT INTO private_code (code, client_id, time_stamp)
    VALUES (new_discount_code, client_id_val, NOW());

    SELECT r.referrer_id INTO referrer FROM refers r WHERE r.referee_id = referee;

    WHILE referrer IS NOT NULL LOOP
        discount_percentage := 50 / (2 * current_level);

        SELECT client_id INTO client_id_val FROM client WHERE referral_code = referrer;

        IF discount_percentage < 1 THEN
            -- Fixed discount (50,000 Tomans)
            INSERT INTO discount_code (code, amount, discount_limit, expiration_time, code_type)
            VALUES (
                nextval('discount_code_code_seq'), 
                50000,  
                50000,  
                NOW() + INTERVAL '1 week', 
                'private'
            ) RETURNING code INTO new_discount_code;
        ELSE 
            -- Percentage-based discount
            INSERT INTO discount_code (code, amount, discount_limit, expiration_time, code_type)
            VALUES (
                nextval('discount_code_code_seq'), 
                discount_percentage / 100,  
                1000000,  
                NOW() + INTERVAL '1 week', 
                'private'
            ) RETURNING code INTO new_discount_code;
        END IF;

        INSERT INTO private_code (code, client_id, time_stamp)
        VALUES (new_discount_code, client_id_val, NOW());

        SELECT r.referrer_id INTO referrer FROM refers r WHERE r.referee_id = referrer;

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
    cart_status_val cart_status_enum;
BEGIN
    SELECT cart_status
    INTO cart_status_val
    FROM shopping_cart
    WHERE cart_number = NEW.cart_number 
      AND client_id = NEW.client_id;

    IF cart_status_val = 'blocked'THEN
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
    stock_count_val SMALLINT;
BEGIN
    SELECT stock_count INTO stock_count_val
    FROM product 
    WHERE id = NEW.product_id;

    IF stock_count_val < NEW.quantity THEN
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
    cart_count_active   INT;
    cart_count_total    INT;
    is_vip              BOOLEAN;
BEGIN

    SELECT EXISTS (
        SELECT 1
        FROM vip_client
        WHERE client_id = NEW.client_id
    ) INTO is_vip;

    SELECT COUNT(*) 
    INTO cart_count_active
    FROM shopping_cart
    WHERE client_id = NEW.client_id
      AND cart_status = 'active';

    SELECT COUNT(*) 
    INTO cart_count_total
    FROM shopping_cart
    WHERE client_id = NEW.client_id;

    -- Restrict total carts for both regular and VIP users
    IF cart_count_total >= 5 THEN 
        RAISE EXCEPTION 'users cannot have more than five shopping carts.';
    END IF;

    IF is_vip THEN
        -- VIP users can have up to 5 total carts
        IF cart_count_active >= 5 THEN
            RAISE EXCEPTION 'VIP users cannot have more than five active shopping carts.';
        END IF;
    ELSE
        -- Non-VIP users can have only 1 active cart
        IF cart_count_active >= 1 THEN
            RAISE EXCEPTION 'Registered users cannot have more than one active shopping cart.';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER enforce_cart_limit_trigger
BEFORE INSERT OR UPDATE ON shopping_cart
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
    code_record RECORD := NULL;
    usage INT;
BEGIN
    SELECT d.usage_count, d.usage_limit, d.expiration_time
    INTO code_record
    FROM discount_code d
    WHERE d.code = NEW.code;

    IF code_record IS NULL THEN
        RAISE EXCEPTION 'Invalid discount code.';
    END IF;

    SELECT COUNT(code)   
    INTO usage
    FROM applied_to
    WHERE code = NEW.code
    GROUP BY code;

    IF code_record.expiration_time < NOW() THEN
        RAISE EXCEPTION 'The discount code has expired.';
    END IF;

    IF usage >= code_record.usage_limit THEN
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

    IF NEW.amount <= 0 THEN
        RAISE EXCEPTION 'Deposit amount must be positive.';
    END IF;

    UPDATE client
    SET wallet_balance = wallet_balance + NEW.amount
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
    - when an order is finalized and paid, the associated cart is unlocked.
    - If the user's subscription has expired and the shopping cart is locked,
        the cart will be blocked after finalizing.
*/
CREATE OR REPLACE FUNCTION unlock_cart_after_payment()
RETURNS TRIGGER AS $$
DECLARE
    is_vip_expired BOOLEAN;
BEGIN
    SELECT (v.expiration_time < NOW()) INTO is_vip_expired
    FROM vip_client v
    WHERE v.client_id = NEW.client_id;

    IF is_vip_expired THEN
        UPDATE shopping_cart
        SET cart_status = 'blocked'
        WHERE client_id = NEW.client_id
          AND cart_number = NEW.cart_number;
    ELSE
        UPDATE shopping_cart
        SET cart_status = 'active'  
        WHERE client_id = NEW.client_id
          AND cart_number = NEW.cart_number;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


/*
Ensures that:
    After a subscription :
    - If a client already exists in vip_client, update the expiration_time.
    - If the client is not in vip_client, insert them with an expiration time of 1 month.
*/
CREATE OR REPLACE FUNCTION convert_to_vip()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO vip_client(client_id, expiration_time)
    VALUES (NEW.client_id, NOW() + INTERVAL '1 month')
    ON CONFLICT (client_id) 
    DO UPDATE SET expiration_time = NOW() + INTERVAL '1 month'; 

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER convert_vip_after_sub_trigger
AFTER INSERT ON subscribes 
FOR EACH ROW
EXECUTE FUNCTION convert_to_vip();


-- Job scheduler --

/*
Job:
    15% of the cost paid on orders must be returned to the digital wallet of special users on a monthly basis.
*/
CREATE OR REPLACE FUNCTION add_monthly_cashback()
RETURNS VOID AS $$
DECLARE
    vip_client_record RECORD;
    cashback_amount DECIMAL(12,2);
BEGIN
    FOR vip_client_record IN 
        SELECT c.client_id, SUM(at.cart_price) * 0.15 AS total_cashback
        FROM issued_for ifo
        JOIN transaction t ON ifo.tracking_code = t.tracking_code
        JOIN added_to adt ON ifo.client_id = adt.client_id AND ifo.cart_number = adt.cart_number AND ifo.locked_number = adt.locked_number
        JOIN vip_client vc ON ifo.client_id = vc.client_id
        WHERE t.transaction_status = 'Successful'
        AND t.time_stamp >= DATE_TRUNC('month', CURRENT_DATE) - INTERVAL '1 month'
        AND t.time_stamp < DATE_TRUNC('month', CURRENT_DATE)
        GROUP BY c.client_id
    LOOP
        cashback_amount := vip_client_record.total_cashback;
        
        UPDATE client
        SET wallet_balance = wallet_balance + cashback_amount
        WHERE client_id = vip_client_record.client_id;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule(
    '0 0 1 * *', -- Runs at midnight on the 1st of each month
    'SELECT add_monthly_cashback();'
);


/*
Job:
    If the user fails to finalize their order and make the payment within the 3-day period:
    - The following actions must be taken:
    - Items in the shopping cart should be returned to inventory, and their stock levels must be updated.
    - The user's shopping cart should be blocked for 7 days.
*/
CREATE OR REPLACE FUNCTION check_order()
RETURNS VOID AS $$
DECLARE 
    locked_cart_expired RECORD;
    product_rec RECORD;
BEGIN
    FOR locked_cart_expired IN 
        SELECT *
        FROM locked_shopping_cart NATURAL JOIN shopping_cart 
        WHERE cart_status = 'locked' 
          AND (NOW() - time_stamp) > INTERVAL '3 day'
    LOOP
        -- Return amount of products stock count
        FOR product_rec IN 
            SELECT product_id, quantity
            FROM locked_cart_expired NATURAL JOIN added_to
        LOOP
            UPDATE product 
            SET stock_count = stock_count + product_rec.quantity
            WHERE id = product_rec.product_id;
        END LOOP;

        -- Blocking the shopping cart for 7 days
        UPDATE shopping_cart
        SET cart_status = 'blocked'
        WHERE cart_number = locked_cart_expired.cart_number
          AND client_id = locked_cart_expired.client_id;
        
        UPDATE locked_shopping_cart
        SET time_stamp = NOW() + INTERVAL '7 days'
        WHERE cart_number = locked_cart_expired.cart_number
          AND client_id = locked_cart_expired.client_id
          AND locked_number = locked_cart_expired.locked_number;

    END LOOP;

END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule(
    '0 0 * * *', -- Runs daily at midnight
    'SELECT check_order();'
);


/*
Job:
    When a VIP subscription expires, the user's additional shopping carts 
    (except the first one that all registered users have) should be removed.
*/
CREATE OR REPLACE FUNCTION handle_subscription_end()
RETURN VOID AS $$
DECLARE
    vip_rec RECORD;
BEGIN
    FOR vip_rec IN
        SELECT *
        FROM vip_client
        WHERE expiration_time < NOW();
    LOOP
        -- Block additional carts (cart_number > 1) except locked cart
        FOR cart_rec IN
            SELECT cart_number
            FROM shopping_cart
            WHERE client_id = vip_rec.client_id
              AND cart_number > 1  
              AND cart_status <> 'locked'
        LOOP
            UPDATE shopping_cart
            SET cart_status = 'blocked'
            WHERE client_id = vip_rec.client_id
              AND cart_number = cart_rec.cart_number;
        END LOOP;

    END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule(
    '0 0 * * *', -- Runs daily at midnight
    'SELECT handle_subscription_end();'
);