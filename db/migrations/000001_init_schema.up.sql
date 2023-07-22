CREATE TABLE "answers"
(
    "id"          BIGSERIAL PRIMARY KEY NOT NULL,
    "question_id" bigint                NOT NULL,
    "text"        varchar               NOT NULL,
    "is_correct"  bool                  NOT NULL DEFAULT False,
    "created_at"  TIMESTAMP             NOT NULL DEFAULT (now())
);

CREATE TABLE "questions"
(
    "id"         BIGSERIAL PRIMARY KEY NOT NULL,
    "text"       varchar UNIQUE        NOT NULL,
    "hint"       varchar               NOT NULL,
    "category"   varchar               NOT NULL,
    "created_at" TIMESTAMP             NOT NULL DEFAULT (now())
);

CREATE TABLE "categories"
(
    "id"         BIGSERIAL PRIMARY KEY NOT NULL,
    "key"        varchar UNIQUE        NOT NULL,
    "name"       varchar UNIQUE        NOT NULL,
    "created_at" TIMESTAMP             NOT NULL DEFAULT (now())
);

CREATE TABLE "answered_questions"
(
    "id"          BIGSERIAL PRIMARY KEY NOT NULL,
    "user_id"     bigint                NOT NULL,
    "question_id" bigint                NOT NULL,
    "answered_at" TIMESTAMP             NOT NULL DEFAULT (now())
);

CREATE TABLE "users"
(
    "id"         BIGSERIAL PRIMARY KEY NOT NULL,
    "email"      varchar UNIQUE        NOT NULL,
    "password"   varchar               NOT NULL,
    "created_at" TIMESTAMP             NOT NULL DEFAULT (now())
);

CREATE INDEX ON "answers" ("question_id");

CREATE INDEX ON "questions" ("id");

CREATE INDEX ON "questions" ("category");

CREATE INDEX ON "questions" ("category", "id");

CREATE INDEX ON "categories" ("key");

CREATE INDEX ON "answered_questions" ("user_id");

CREATE INDEX ON "answered_questions" ("question_id");

CREATE UNIQUE INDEX "idx_unique_user_question" ON "answered_questions" ("user_id", "question_id");

CREATE INDEX ON "users" ("id");

COMMENT
ON COLUMN "categories"."name" IS 'name for user';

COMMENT
ON COLUMN "categories"."key" IS 'name for dev';

ALTER TABLE "answers"
    ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE;

ALTER TABLE "questions"
    ADD FOREIGN KEY ("category") REFERENCES "categories" ("key") ON DELETE CASCADE;

ALTER TABLE "answered_questions"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "answered_questions"
    ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE;