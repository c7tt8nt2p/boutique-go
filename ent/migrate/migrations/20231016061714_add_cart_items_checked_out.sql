-- Modify "cart_items" table
ALTER TABLE "cart_items" ADD COLUMN "checked_out" boolean NOT NULL DEFAULT false;
