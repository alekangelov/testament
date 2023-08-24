CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at_column()   
  RETURNS TRIGGER AS $$
  BEGIN
      NEW.updated_at = now();
      RETURN NEW;   
  END;
$$ language 'plpgsql';


CREATE TABLE IF NOT EXISTS public.users
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    email character varying(255) NOT NULL UNIQUE,
    password character varying(255) NOT NULL,
    "SOCIAL_PROVIDER" character varying(255),
    "SOCIAL_ID" character varying(255),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TRIGGER update_users_modtime
  BEFORE UPDATE ON public.users
  for each row execute procedure
  update_updated_at_column();


CREATE TABLE IF NOT EXISTS public.verifications
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    type character varying(255) NOT NULL,
    code character varying(6) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT primary_key_id PRIMARY KEY (id),
    CONSTRAINT foreign_key_user_verification FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);


CREATE TRIGGER update_verifications_modtime
  BEFORE UPDATE ON public.verifications
  for each row execute procedure
  update_updated_at_column();
