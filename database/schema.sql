CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "uuid" VARCHAR(64) NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("email"),
    UNIQUE ("uuid")
);

CREATE TABLE "sessions" (
    "id" SERIAL NOT NULL,
    "uuid" VARCHAR(64) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "user_id" INTEGER,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") 
        REFERENCES "users" ("id") 
        ON UPDATE NO ACTION 
        ON DELETE NO ACTION
);

CREATE TABLE "threads" (
    "id" SERIAL NOT NULL,
    "uuid" VARCHAR(64) NOT NULL,
    "topic" TEXT,
    "user_id" INTEGER,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    CONSTRAINT "threads_user_id_fkey" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE "posts" (
    "id" SERIAL NOT NULL,
    "uuid" VARCHAR(64) NOT NULL,
    "body" TEXT NOT NULL,
    "user_id" INTEGER,
    "thread_id" INTEGER,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    CONSTRAINT "posts_thread_id_fkey" FOREIGN KEY ("thread_id")
        REFERENCES "threads" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT "posts_user_id_fkey" FOREIGN KEY ("user_id")
        REFERENCES "users" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);