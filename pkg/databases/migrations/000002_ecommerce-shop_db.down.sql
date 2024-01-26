BEGIN;

TRUNCATE TABLE "users" CASCADE;
TRUNCATE TABLE "oauth" CASCADE;
TRUNCATE TABLE "roles" CASCADE;
TRUNCATE TABLE "products" CASCADE;
TRUNCATE TABLE "categories" CASCADE;
TRUNCATE TABLE "products_categories" CASCADE;
TRUNCATE TABLE "images" CASCADE;
TRUNCATE TABLE "orders" CASCADE;
TRUNCATE TABLE "products_orders" CASCADE;

    --reset seq
SELECT SETVAL ((SELECT PG_GET_SERIAL_SEQUENCE('"roles"', 'id')), 1, FALSE);
SELECT SETVAL ((SELECT PG_GET_SERIAL_SEQUENCE('"categories"', 'id')), 1, FALSE);

COMMIT;