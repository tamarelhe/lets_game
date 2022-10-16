CREATE TABLE public.lg_sessions (
    "id" uuid PRIMARY KEY,
    "email" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT false,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE public.lg_sessions ADD FOREIGN KEY (email) REFERENCES public.lg_users (email);