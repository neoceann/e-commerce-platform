CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    category_id UUID NOT NULL REFERENCES category(id) ON DELETE RESTRICT,
    price DECIMAL(8, 2) NOT NULL CHECK (price >= 0),
    available_stock SMALLINT NOT NULL DEFAULT 0 CHECK (available_stock >= 0),
    last_available_update_date DATE DEFAULT NOW(),
    supplier_id UUID NOT NULL REFERENCES supplier(id) ON DELETE RESTRICT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_product_available ON product(available_stock) WHERE available_stock > 0;
CREATE INDEX idx_product_category ON product(category_id);
CREATE INDEX idx_product_supplier ON product(supplier_id);

CREATE OR REPLACE FUNCTION update_last_available_update_date()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.available_stock != OLD.available_stock THEN
        NEW.last_available_update_date = CURRENT_DATE;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_product_updated_at 
    BEFORE UPDATE ON product 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_last_available_update_date
    BEFORE UPDATE ON product
    FOR EACH ROW
    EXECUTE FUNCTION update_last_available_update_date();