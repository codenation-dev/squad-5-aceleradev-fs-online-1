CREATE TABLE public.public_agent (
	id varchar(26) NOT NULL,
	name varchar(100) NOT NULL,
	department varchar(100) NULL,
	occupation varchar(100) NULL,
	salary numeric(18,2) NOT NULL,
	inserted_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	verified_at timestamp NOT NULL,
	CONSTRAINT public_agent_pk PRIMARY KEY (id)
);
