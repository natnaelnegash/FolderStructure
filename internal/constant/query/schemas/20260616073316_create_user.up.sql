 

DO $$
BEGIN
IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname='status')THEN
CREATE TYPE status AS ENUM('Active','Locked','Blocked','Reset');
END IF;
END$$;
CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(255) NOT NULL ,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    password BYTEA NOT NULL,
    pinattempt INT NOT NULL DEFAULT 0,
    pinstatus status NOT NULL DEFAULT 'Active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS id_users_id
ON users(id);

