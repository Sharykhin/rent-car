CREATE TABLE public.consumers (
	id uuid not null DEFAULT uuid_generate_v4 (),
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	email varchar(80) not null unique,
	created_at timestamptz(0) not null,
	CONSTRAINT consumers_pk PRIMARY KEY (id)
);
