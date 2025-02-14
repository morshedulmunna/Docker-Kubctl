CREATE TABLE IF NOT EXISTS user_roles (
    id bigserial PRIMARY KEY,
    role_name varchar(255) NOT NULL,
    create_at timestamp with time zone NOT NULL DEFAULT NOW(),
    update_at timestamp with time zone NOT NULL DEFAULT NOW()
);
