CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.cars (
	id uuid not null DEFAULT uuid_generate_v4 (),
	model varchar(50) not null,
	created_at timestamptz(0) not null,
	CONSTRAINT cars_pk PRIMARY KEY (id)
);
