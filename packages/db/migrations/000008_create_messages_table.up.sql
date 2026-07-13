CREATE TYPE status_types AS ENUM(
   'queued'
   'delivered',
)

CREATE TABLE messages (
   message_id  UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
   sender_id   UUID NOT NULL,
   reciver_id  UUID NOT NULL,
   message     TEXT NOT NULL,
   status      status_types NOT NULL DEFAULT 'delivered',
   sent_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()

   CONSTRAINT fk_service_requests_client
      FOREIGN KEY (sender_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE,

   CONSTRAINT fk_service_requests_client
      FOREIGN KEY (reciver_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE,
);