CREATE TABLE "answer" (
                          "id" BIGSERIAL PRIMARY KEY NOT NULL,
                          "question_id" bigint NOT NULL,
                          "text" varchar NOT NULL,
                          "is_correct" bool NOT NULL DEFAULT False
);

CREATE TABLE "question" (
                            "id" BIGSERIAL PRIMARY KEY NOT NULL,
                            "text" varchar NOT NULL,
                            "hint" varchar NOT NULL,
                            "category" varchar NOT NULL
);

CREATE TABLE "category" (
                            "name" varchar UNIQUE NOT NULL,
                            "key" varchar UNIQUE NOT NULL
);

CREATE TABLE "answered_question" (
                                     "user_id" bigint NOT NULL,
                                     "question_id" bigint NOT NULL,
                                     "answered_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "user" (
                        "id" BIGSERIAL PRIMARY KEY NOT NULL,
                        "mail" varchar UNIQUE NOT NULL,
                        "password" varchar NOT NULL,
                        "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE INDEX ON "answer" ("question_id");

CREATE INDEX ON "question" ("id");

CREATE INDEX ON "question" ("category");

CREATE INDEX ON "question" ("category", "id");

CREATE INDEX ON "category" ("key");

CREATE INDEX ON "answered_question" ("user_id");

CREATE INDEX ON "answered_question" ("question_id");

CREATE INDEX ON "answered_question" ("user_id", "question_id");

CREATE INDEX ON "user" ("id");

COMMENT ON COLUMN "category"."name" IS 'name for user';

COMMENT ON COLUMN "category"."key" IS 'name for dev';

ALTER TABLE "answer" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");

ALTER TABLE "question" ADD FOREIGN KEY ("category") REFERENCES "category" ("key");

ALTER TABLE "answered_question" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "answered_question" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");