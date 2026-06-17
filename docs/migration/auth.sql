CREATE TABLE IF NOT EXISTS auth (
    id int 
    public_id varchar
    email varchar
    password varchar
    role varchar
    is_active bool
    created_at timestamp
    updated_at timestamp
)