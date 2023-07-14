CREATE TABLE "Titles" (
    "title_id" bigserial PRIMARY KEY,
    "title" VARCHAR(10) NOT NULL
);

CREATE TABLE "Users" (
    "userId" bigserial PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(60) NOT NULL,
    "phone" CHAR(13),
    "region" VARCHAR(100),
    "gender" VARCHAR(1) DEFAULT 'N',
    "title_id" INTEGER NOT NULL,
    "firstName" VARCHAR(255) NOT NULL,
    "lastName" VARCHAR(255) NOT NULL,
    "age" SMALLINT,
    CONSTRAINT fk_title FOREIGN KEY ("title_id") REFERENCES "Titles"("title_id")
);
