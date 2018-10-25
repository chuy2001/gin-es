DROP TABLE "user";

CREATE TABLE "user" (id integer NOT NULL,email character varying,password character varying,name character varying,updated_at integer,created_at integer);
Insert user (id ,email ,password,name ,updated_at,created_at) VALUES (1,'michael.chu@test.com','123456');


SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER(" + form.Email +") LIMIT 1"

