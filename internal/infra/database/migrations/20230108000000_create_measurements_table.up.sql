CREATE TABLE IF NOT EXISTS public.measurements
(
    id           serial PRIMARY KEY,
    device_id    integer NOT NULL REFERENCES public.devices(id),
    room_id      integer REFERENCES public.rooms(id),
    value        double precision NOT NULL,
    created_date timestamptz NOT NULL,
    updated_date timestamptz NOT NULL,
    deleted_date timestamptz
);
