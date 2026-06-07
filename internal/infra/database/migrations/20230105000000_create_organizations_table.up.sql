CREATE TABLE IF NOT EXISTS public.organizations
(
    id           serial PRIMARY KEY,
    user_id      integer NOT NULL REFERENCES public.users(id),
    name         varchar(100) NOT NULL,
    description  text,
    city         varchar(100) NOT NULL,
    address      varchar(200) NOT NULL,
    lat          double precision,
    lon          double precision,
    created_date timestamptz NOT NULL,
    updated_date timestamptz NOT NULL,
    deleted_date timestamptz
);
