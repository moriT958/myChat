-- Create "users" table
CREATE TABLE
    "users" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "name" character varying(255) NULL,
        "email" character varying(255) NOT NULL,
        "password" character varying(255) NOT NULL,
        "created_at" timestamp NOT NULL,
        PRIMARY KEY ("id"),
        CONSTRAINT "users_email_key" UNIQUE ("email"),
        CONSTRAINT "users_uuid_key" UNIQUE ("uuid")
    );

-- Create "threads" table
CREATE TABLE
    "threads" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "topic" text NULL,
        "user_id" integer NULL,
        "created_at" timestamp NOT NULL,
        PRIMARY KEY ("id"),
        CONSTRAINT "threads_uuid_key" UNIQUE ("uuid"),
        CONSTRAINT "threads_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
    );

-- Create "posts" table
CREATE TABLE
    "posts" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "body" text NULL,
        "user_id" integer NULL,
        "thread_id" integer NULL,
        "created_at" timestamp NOT NULL,
        PRIMARY KEY ("id"),
        CONSTRAINT "posts_uuid_key" UNIQUE ("uuid"),
        CONSTRAINT "posts_thread_id_fkey" FOREIGN KEY ("thread_id") REFERENCES "threads" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
        CONSTRAINT "posts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
    );

-- Create "sessions" table
CREATE TABLE
    "sessions" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "email" character varying(255) NULL,
        "user_id" integer NULL,
        "created_at" timestamp NOT NULL,
        PRIMARY KEY ("id"),
        CONSTRAINT "sessions_uuid_key" UNIQUE ("uuid"),
        CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
    );