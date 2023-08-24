CREATE TABLE IF NOT EXISTS public.migrations (
  id int NOT NULL,
  name character varying(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  CONSTRAINT migrations_pkey PRIMARY KEY (id)
)