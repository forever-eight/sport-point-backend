CREATE TABLE IF NOT EXISTS studio
(
    id    SERIAL NOT NULL PRIMARY KEY,
    title TEXT   NOT NULL UNIQUE,
    address TEXT NOT NULL
);

INSERT INTO studio (title, address)
VALUES ('Sunflower dance family', '1st street, 7'),
       ('First GYM', '2nd street, 8');

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
    -- todo сделать отдельную табличку? и массивом подавать туда дни недели
    weekday     INT    NOT NULL,
    started_at  INT    NOT NULL,
    duration    INT    NOT NULL,
    amount      INT    NOT NULL,
);

INSERT INTO classes (studio_id, title, type_id, description, weekday, started_at, duration, amount)
VALUES (1, 'hip hop', 2, 'dance classes', 1, 1200, 60, 30),
       (1, 'hip hop', 2, 'dance classes', 5, 1200, 120, 30),
       (2, 'cycle', 2, 'cardio', 7, 600, 50, 60);

CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL NOT NULL PRIMARY KEY,
    name     TEXT   NOT NULL,
    email    TEXT   NOT NULL UNIQUE,
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
