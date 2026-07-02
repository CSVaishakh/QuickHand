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
    handymen_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id       UUID NOT NULL UNIQUE,
    type          handyman_type NOT NULL,

    CONSTRAINT fk_handymen_user
      FOREIGN KEY (user_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE
);