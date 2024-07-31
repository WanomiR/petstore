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