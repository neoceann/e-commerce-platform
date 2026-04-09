CREATE TABLE IF NOT EXISTS supplier (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    address_id UUID REFERENCES address(id) ON DELETE SET NULL,
    phone_number VARCHAR(12) NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_supplier_address ON supplier(address_id);

CREATE TRIGGER update_supplier_updated_at 
    BEFORE UPDATE ON supplier 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();