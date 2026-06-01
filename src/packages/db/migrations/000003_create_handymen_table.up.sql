CREATE TYPE handyman_type AS ENUM (
    'plumber',
    'electrician',
    'carpenter',
    'mason',
    'mechanic',
    'hvac_technician',
    'landscaper',
    'deep_cleaner'
);

CREATE TABLE handymen (
    id BIGSERIAL PRIMARY KEY,

    user_id UUID NOT NULL UNIQUE,

    type handyman_type NOT NULL,

    CONSTRAINT fk_handymen_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);