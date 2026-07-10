CREATE TYPE status_types AS ENUM (
   'requested',
   'hired',
   'rejected'
);

CREATE TABLE service_requests(
   req_id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   client_id         UUID  NOT   NULL,
   handyman_id       UUID  NOT   NULL,
   job_id            UUID  NOT   NULL,
   status            status_types   NOT NULL,

   CONSTRAINT fk_service_requests_client
      FOREIGN KEY (client_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE,

   CONSTRAINT fk_service_requests_handyman
      FOREIGN KEY (handyman_id)
      REFERENCES users(user_id)
      ON DELETE SET NULL,

   CONSTRAINT fk_service_requests_job
      FOREIGN KEY (job_id)
      REFERENCES jobs(job_id)
      ON DELETE CASCADE
);

