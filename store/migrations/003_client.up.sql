CREATE TABLE IF NOT EXISTS client (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_name TEXT NOT NULL,
    client_surname TEXT NOT NULL,
    birthday DATE,
    gender CHAR(1) CHECK (gender IN ('M', 'F')),
    registration_date TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
    address_id UUID REFERENCES address(id) ON DELETE SET NULL,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_client_name ON client(client_name, client_surname);
CREATE INDEX IF NOT EXISTS idx_client_address ON client(address_id);

CREATE TRIGGER update_client_updated_at 
    BEFORE UPDATE ON client 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();