CREATE TABLE IF NOT EXISTS public.votes (
  id uuid DEFAULT gen_random_uuid(),
  voting_date DATE NOT NULL,
  category VARCHAR(3) NOT NULL,
  PRIMARY KEY (id)
);

CREATE VIEW public.votes_report AS
SELECT v.category, count(1) as total_per_category
FROM public.votes v
GROUP BY v.category;