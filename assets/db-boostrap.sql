-- DROP TABLE public.lg_user;

CREATE TABLE public.lg_user (
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
CREATE INDEX user_id_idx ON public.lg_user USING btree (id);
CREATE UNIQUE INDEX lg_user_email_idx ON public.lg_user (email);



-- DROP TABLE public.lg_group;

CREATE TABLE public.lg_group (
	id uuid NOT NULL,
	"name" varchar NOT NULL,
	avatar varchar NULL,
	members jsonb NOT NULL,
	created_at timestamptz NOT NULL,
	CONSTRAINT group_pk PRIMARY KEY (id)
);
CREATE INDEX group_id_idx ON public.lg_group USING btree (id);


-- DROP TABLE public.lg_game;

CREATE TABLE public.lg_game (
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
CREATE INDEX game_datetime_idx ON public.lg_game USING btree (datetime);
CREATE INDEX game_id_idx ON public.lg_game USING btree (id);


-- public.lg_game foreign keys

ALTER TABLE public.lg_game ADD CONSTRAINT game_fk FOREIGN KEY (id) REFERENCES public.lg_group(id);


-- DROP delete_group;

CREATE OR REPLACE FUNCTION delete_group(group_id text)
   RETURNS integer AS $$
DECLARE  
  rec_member RECORD;
 
  c_members CURSOR (gid text) FOR 
     SELECT m->>'id' AS member_id
	   FROM lg_group g,
	        jsonb_array_elements(g.members) m
	  WHERE g.id = gid::uuid;	
	 
BEGIN
  OPEN c_members(group_id);
	
  LOOP
     FETCH c_members INTO rec_member;
     EXIT WHEN NOT FOUND;
     
     UPDATE lg_user SET groups = array_remove(groups, group_id)
      WHERE id = rec_member.member_id::uuid;       
  END LOOP;
 
  DELETE FROM lg_group g WHERE g.id = group_id::uuid;	
 
  RETURN 1;
 
  EXCEPTION 
	   WHEN others THEN 
		 RAISE EXCEPTION 'error deleting group %',sqlerrm;
		 RETURN 0;
 
end; 
$$

language plpgsql;