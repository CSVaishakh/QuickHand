CREATE TABLE message_queue (
   queue_id UUID NOT NULL DEFAULT gen_random_uuid(),
   message_id UUID NOT NULL UNIQUE,

   CONSTRAINT fk_service_requests_client
      FOREIGN KEY (message_id)
      REFERENCES messages(message_id)
      ON DELETE CASCADE,
);