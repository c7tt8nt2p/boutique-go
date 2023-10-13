-- Create "checkouts" table
CREATE TABLE "checkouts" ("id" uuid NOT NULL, "total_price" double precision NOT NULL, "created_at" timestamptz NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "checkouts_users_checkout" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "checkout_items" table
CREATE TABLE "checkout_items" ("id" uuid NOT NULL, "cart_item_id" uuid NOT NULL, "checkout_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "checkout_items_cart_items_checkout_item" FOREIGN KEY ("cart_item_id") REFERENCES "cart_items" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "checkout_items_checkouts_checkout_item" FOREIGN KEY ("checkout_id") REFERENCES "checkouts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "checkout_items_cart_item_id_key" to table: "checkout_items"
CREATE UNIQUE INDEX "checkout_items_cart_item_id_key" ON "checkout_items" ("cart_item_id");
-- Add pk ranges for ('checkouts'),('checkout_items') tables
INSERT INTO "ent_types" ("type") VALUES ('checkouts'), ('checkout_items');
