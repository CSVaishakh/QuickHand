CREATE TYPE job_types AS ENUM (
   'plumbing',
   'electrical',
   'carpentry',
   'masonry',
   'mechanical',
   'hvac',
   'landscaping',
   'deep_cleaning'
);

CREATE TYPE hire_types AS ENUM (
   'direct_hire',
   'bid_to_get'
);

CREATE TYPE urgency_levels AS ENUM (
    'instant',
    'urgent',
    'earliest_available',
    'flexible'
);

CREATE TABLE jobs (
   job_id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   client_id          UUID NOT NULL,
   handyman_id       UUID NULL,
   hire_type         hire_types NOT NULL,
   job_type          job_types NOT NULL,
   description       VARCHAR(150) NOT NULL,
   budget            NUMERIC(10,2),
   created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   deadline_at       TIMESTAMP NOT NULL,
   urgency           urgency_levels NOT NULL,

   CONSTRAINT fk_addresses_user
      FOREIGN KEY (user_id)
      REFERENCES users(user_id)
      ON DELETE CASCADE,
      
   CONSTRAINT fk_jobs_handyman
      FOREIGN KEY (handyman_id)
      REFERENCES handymen(user_id)
      ON DELETE SET NULL
);