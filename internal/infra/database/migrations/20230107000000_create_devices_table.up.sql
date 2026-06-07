CREATE TABLE IF NOT EXISTS public.devices
(
    id                serial PRIMARY KEY,
    organization_id   integer NOT NULL REFERENCES public.organizations(id),
    room_id           integer REFERENCES public.rooms(id),
    guid              uuid NOT NULL,
    inventory_number  varchar(100),
    serial_number     varchar(100),
    characteristics   text,
    category          varchar(20) NOT NULL CHECK (category IN ('SENSOR', 'ACTUATOR')),
    units             varchar(50),
    power_consumption double precision,
    created_date      timestamptz NOT NULL,
    updated_date      timestamptz NOT NULL,
    deleted_date      timestamptz
);
