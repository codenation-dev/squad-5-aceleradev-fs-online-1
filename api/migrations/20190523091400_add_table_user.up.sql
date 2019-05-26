CREATE TABLE public."user" (
	id varchar(26) NOT NULL,
	name varchar(100) NOT NULL,
	email varchar(160) NOT NULL,
	username varchar(20) NOT NULL,
	password varchar(300) NOT NULL,
	inserted_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX user_username_idx ON public."user" (username) ;
