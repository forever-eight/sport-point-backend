--CREATE TYPE curr AS ENUM ('EUR', 'TRY', 'RUB');

CREATE TABLE IF NOT EXISTS studio
(
    id    SERIAL NOT NULL PRIMARY KEY,
    title TEXT   NOT NULL
);

INSERT INTO studio (title)
VALUES ('Sunflower dance family'),
       ('First GYM');

CREATE TABLE IF NOT EXISTS classes_type
(
    id    SERIAL NOT NULL PRIMARY KEY,
    title TEXT   NOT NULL
);

INSERT INTO classes_type (title)
VALUES ('Relax'),
       ('Fitness'),
       ('Beauty');


CREATE TABLE IF NOT EXISTS classes
(
    id          SERIAL NOT NULL PRIMARY KEY,
    studio_id   INT REFERENCES studio (id) ON DELETE CASCADE,
    title       TEXT   NOT NULL,
    type_id     INT    NOT NULL REFERENCES classes_type (id),
    description TEXT   NOT NULL,
    weekday     INT    NOT NULL,
    started_at  INT    NOT NULL,
    ended_at    INT    NOT NULL,
    amount      INT    NOT NULL,
    currency    curr   NOT NULL DEFAULT 'EUR'
);

INSERT INTO classes (studio_id, title, type_id, description, weekday, started_at, ended_at, amount, currency)
VALUES (1, 'hip hop', 2, 'dance classes', 1, 1200, 1260, 30, 'EUR'),
       (1, 'hip hop', 2, 'dance classes', 5, 1200, 1260, 30, 'EUR'),
       (2, 'cycle', 2, 'cardio', 7, 600, 690, 60, 'TRY');

CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL NOT NULL PRIMARY KEY,
    name     TEXT   NOT NULL,
    email    TEXT   NOT NULL,
    password TEXT   NOT NULL
);

INSERT INTO users (name, email, password)
VALUES ('Maria T', '123@t.ru', '1234'),
       ('Ivan I', '12345@gmail.com', '0987');

CREATE TABLE IF NOT EXISTS transaction
(
    id         SERIAL NOT NULL PRIMARY KEY,
    points     INT    NOT NULL,
    ending_at  DATE   NOT NULL,
    user_id    INT    NOT NULL,
    is_deleted BOOL   NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS booking
(
    id             SERIAL NOT NULL PRIMARY KEY,
    class_id       INT    NOT NULL REFERENCES classes (id),
    user_id        INT    NOT NULL REFERENCES users (id),
    transaction_id INT    NOT NULL REFERENCES transaction (id)
);
