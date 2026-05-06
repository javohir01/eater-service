-- Drop indexes
DROP INDEX IF EXISTS eater.idx_address_by_eater;
DROP INDEX IF EXISTS eater.idx_r_rating_by_eater;
DROP INDEX IF EXISTS eater.idx_order_by_eater;
DROP INDEX IF EXISTS eater.idx_card_by_eater;
DROP INDEX IF EXISTS eater.idx_sms_code_by_eater;
DROP INDEX IF EXISTS eater.idx_eater_phone_number;

-- Drop tables
DROP TABLE IF EXISTS eater.restaurant_ratings;
DROP TABLE IF EXISTS eater.orders;
DROP TABLE IF EXISTS eater.payment_cards;
DROP TABLE IF EXISTS eater.addresses;
DROP TABLE IF EXISTS eater.eater_profiles;
DROP TABLE IF EXISTS eater.eater_sms_codes;
DROP TABLE IF EXISTS eater.eaters;

-- Drop schema
DROP SCHEMA IF EXISTS eater;