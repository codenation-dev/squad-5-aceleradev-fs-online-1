CREATE TABLE public.customer (
	id varchar(26) NOT NULL,
	name varchar(100) NOT NULL,
    inserted_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT customer_pk PRIMARY KEY (id)
);
