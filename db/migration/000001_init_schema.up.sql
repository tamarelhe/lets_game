-- public.lg_user table
CREATE TABLE public.lg_users (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	avatar varchar NULL,
	is_active bool NOT NULL,
	created_at timetz NOT NULL,
	"groups" text[] NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);
CREATE INDEX user_id_idx ON public.lg_users USING btree (id);
CREATE UNIQUE INDEX lg_user_email_idx ON public.lg_users (email);


-- public.lg_user lg_group
CREATE TABLE public.lg_groups (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	avatar varchar NULL,
	members jsonb NOT NULL,
	created_at timestamptz NOT NULL,
	CONSTRAINT group_pk PRIMARY KEY (id)
);
CREATE INDEX group_id_idx ON public.lg_groups USING btree (id);


-- public.lg_user lg_game
CREATE TABLE public.lg_games (
	id uuid NOT NULL,
	group_id uuid NOT NULL,
	type_id uuid NULL,
	datetime timestamptz NOT NULL,
	members jsonb NULL,
	"location" jsonb NULL,
	constrains jsonb NULL,
	message text NULL,
	created_at timetz NOT NULL,
	CONSTRAINT game_pk PRIMARY KEY (id)
);
CREATE INDEX game_datetime_idx ON public.lg_games USING btree (datetime);
CREATE INDEX game_id_idx ON public.lg_games USING btree (id);


ALTER TABLE public.lg_games ADD CONSTRAINT game_fk FOREIGN KEY (id) REFERENCES public.lg_groups(id);
