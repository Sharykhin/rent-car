CREATE TABLE public.requisitions (
	id uuid NOT NULL DEFAULT uuid_generate_v4 (),
	car_id uuid NOT NULL,
	consumer_id uuid NOT NULL,
	date_from timestamptz(0) NOT NULL,
	date_to timestamptz(0) NOT NULL,
	created_at timestamptz(0) NOT NULL,
	CONSTRAINT requisitions_pk PRIMARY KEY (id),
	CONSTRAINT fk_requisitions_car_id_cars_id FOREIGN KEY (car_id) REFERENCES cars(id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT fk_consumers_consumer_id_cars_id FOREIGN KEY (consumer_id) REFERENCES consumers(id) ON UPDATE CASCADE ON DELETE CASCADE
);