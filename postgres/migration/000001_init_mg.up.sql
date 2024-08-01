--  categories
CREATE TABLE IF NOT EXISTS categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO categories (name)
VALUES ('cat'),
       ('dog'),
       ('rodent')
;

-- tags
CREATE TABLE IF NOT EXISTS tags
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO tags (name)
VALUES ('fluffy'),
       ('funny'),
       ('kind'),
       ('playful'),
       ('calm'),
       ('happy'),
       ('energetic')
;

-- pets
CREATE TABLE IF NOT EXISTS pets
(
    id          SERIAL PRIMARY KEY,
    category_id INTEGER,
    name        VARCHAR(255),
    status      VARCHAR(255),
    is_deleted  BOOLEAN
);

ALTER TABLE pets
    ADD CONSTRAINT check_status
        CHECK ( status IN ('available', 'pending', 'sold') );

INSERT INTO pets (category_id, name, status, is_deleted)
VALUES (1, 'Poppy', 'available', FALSE),
       (1, 'Bella', 'pending', FALSE),
       (1, 'Tilly', 'sold', FALSE),
       (2, 'Abby', 'available', FALSE),
       (2, 'Bailey', 'pending', FALSE),
       (2, 'Rex', 'sold', FALSE),
       (3, 'Basil', 'available', FALSE),
       (3, 'Danger Mouse', 'pending', FALSE),
       (3, 'Jerry', 'sold', FALSE)
;

-- photo urls
CREATE TABLE IF NOT EXISTS photo_urls
(
    id     SERIAL PRIMARY KEY,
    pet_id INTEGER,
    url    VARCHAR(511)
);

INSERT INTO photo_urls (pet_id, url)
VALUES (1, 'https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg'),
       (1, 'https://wallbox.ru/resize/1024x1024/wallpapers/main/201634/8b7e73ae5927008.jpg'),
       (2, 'https://pixy.org/src/471/4710119.jpg'),
       (3, 'https://jooinn.com/images/happy-cat-resting-6.jpg'),
       (4, 'https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg'),
       (5, 'https://c.pxhere.com/photos/f6/1d/dog_pet_small_dog-912658.jpg!d'),
       (6, 'https://jooinn.com/images/pet-dog-142.jpg'),
       (7, 'https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg'),
       (8, 'https://i.pinimg.com/originals/fe/e8/8a/fee88a1d551c31b2217d999146bfdeb1.jpg'),
       (9, 'https://i.pinimg.com/originals/3a/69/ae/3a69aee66a3f324915ee3085baf9c6c4.jpg')
;

-- pet tags
CREATE TABLE IF NOT EXISTS pet_tags
(
    id     SERIAL PRIMARY KEY,
    pet_id INTEGER,
    tag_id INTEGER
);

INSERT INTO pet_tags (pet_id, tag_id)
VALUES (1, 1),
       (1, 7),
       (2, 4),
       (2, 6),
       (3, 1),
       (4, 5),
       (5, 1),
       (5, 2),
       (5, 7),
       (6, 2),
       (6, 6),
       (7, 1),
       (8, 3),
       (9, 4)
;

-- users
CREATE TABLE IF NOT EXISTS users
(
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(255) UNIQUE,
    first_name  VARCHAR(255),
    last_name   VARCHAR(255),
    email       VARCHAR(255),
    password    VARCHAR(255),
    phone       VARCHAR(255),
    user_status INTEGER,
    is_deleted  BOOLEAN
);

INSERT INTO users (username, first_name, last_name, email, password, phone, user_status, is_deleted)
VALUES ('wanomir', 'Ivan', 'Romadin', 'wanomir@yandex.ru',
        '$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu', '7-999-999-99-99', 0, FALSE),
       ('johndoe001', 'John', 'Doe', 'john.doe@gmail.com',
        '$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu', '7-999-999-99-99', 0, FALSE),
       ('jenstar', 'Jennifer', 'Lawrence', 'jen.lawrence@gmail.com',
        '$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu', '7-999-999-99-99', 0, FALSE),
       ('dragonrider', 'Rhaenyra', 'Targaryen', 'r.targaryen@dragonstone.com',
        '$2a$10$TzogjOIjVZ9fY8/J.1EgOOlV9E1IOSGTC5WWYoP.tDewfMkYUAUXu', '7-999-999-99-99', 0, FALSE)
;

-- store
CREATE TABLE IF NOT EXISTS store
(
    id          SERIAL PRIMARY KEY,
    pet_id      INTEGER,
    quantity    INTEGER,
    ship_date   TIMESTAMP,
    status      VARCHAR(255),
    is_complete BOOLEAN,
    is_deleted  BOOLEAN
);

ALTER TABLE store
    ADD CONSTRAINT check_status
        CHECK ( status IN ('placed', 'approved', 'delivered') );

INSERT INTO store (pet_id, quantity, ship_date, status, is_complete, is_deleted)
VALUES (2, 1, now(), 'placed', FALSE, FALSE),
       (3, 1, now(), 'delivered', TRUE, FALSE),
       (5, 1, now(), 'approved', FALSE, FALSE),
       (6, 1, now(), 'delivered', TRUE, FALSE),
       (8, 1, now(), 'placed', FALSE, FALSE),
       (9, 1, now(), 'delivered', TRUE, FALSE)
;
