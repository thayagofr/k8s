CREATE TABLE IF NOT EXISTS public.votes (
  id uuid DEFAULT gen_random_uuid(),
  voting_date DATE NOT NULL,
  category VARCHAR(3) NOT NULL,
  PRIMARY KEY (id)
);