CREATE TABLE IF NOT EXISTS image (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_data BYTEA NOT NULL,
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_image_product ON image(product_id);

CREATE TRIGGER update_image_updated_at 
    BEFORE UPDATE ON image 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();