CREATE TABLE alerts (
   user_id     UUID NOT NULL,
   created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   title       VARCHAR(255) NOT NULL,
   message     TEXT NOT NULL,

   PRIMARY KEY (user_id, created_at),

   CONSTRAINT fk_alerts_user
      FOREIGN KEY (user_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE
);