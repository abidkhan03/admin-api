BEGIN;

CREATE TABLE "user"
(
    id       BIGSERIAL   NOT NULL PRIMARY KEY,
    email    VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE word
(
    id   BIGSERIAL   NOT NULL PRIMARY KEY,
    word VARCHAR(64) NOT NULL UNIQUE
);

CREATE TABLE pos
(
    id            BIGSERIAL NOT NULL PRIMARY KEY,

    category      INT       NOT NULL,
    type          INT,
    degree        INT,
    gen           INT,
    num           INT,
    possessorpers INT,
    possessornum  INT,
    person        INT,
    neclass       INT,
    nesubclass    INT,
    "case"        INT,
    polite        INT,
    mood          INT,
    tense         INT,
    punctenclose  INT,


    UNIQUE (category, type, degree, gen, num, possessorpers, possessornum, person, neclass,
            nesubclass, "case", polite, mood, tense, punctenclose)
);

CREATE TABLE word_pos
(
    word_id  BIGINT NOT NULL REFERENCES word (id),
    pos_id   BIGINT NOT NULL REFERENCES pos (id),
    lemma_id BIGINT NOT NULL REFERENCES word (id),

    UNIQUE (word_id, pos_id, lemma_id)
);

CREATE TABLE "class"
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
    word_id     BIGINT      REFERENCES word (id),
    description VARCHAR(500)
);

CREATE TABLE word_class
(
    word_id  BIGINT NOT NULL REFERENCES word (id),
    class_id BIGINT NOT NULL REFERENCES "class" (id) ON DELETE CASCADE,

    UNIQUE (word_id, class_id)
);

CREATE TABLE token
(
    id       BIGSERIAL NOT NULL PRIMARY KEY,
    pos_id   BIGINT REFERENCES pos (id),
    word_id  BIGINT REFERENCES word (id),
    class_id BIGINT REFERENCES "class" (id) ON DELETE CASCADE,

    UNIQUE (pos_id, word_id, class_id)
);

INSERT INTO token(id, pos_id, word_id, class_id)
VALUES (1, NULL, NULL, NULL);

SELECT setval(pg_get_serial_sequence('token', 'id'), max(id))
FROM token;


CREATE TABLE pattern
(
    id        BIGSERIAL       NOT NULL PRIMARY KEY,
    token_ids BIGINT[] UNIQUE NOT NULL -- REFERENCES token (id)
);

CREATE TABLE category
(
    id        BIGSERIAL   NOT NULL PRIMARY KEY,
    name      VARCHAR(64) NOT NULL UNIQUE,
    parent_id BIGINT REFERENCES category (id) ON DELETE CASCADE
);

CREATE TABLE category_example
(
    category_id     BIGINT UNIQUE NOT NULL REFERENCES category (id) ON DELETE CASCADE,
    rule            VARCHAR(255),
    tip             VARCHAR(255),
    phrase          VARCHAR(255)  NOT NULL,
    pattern_id      BIGINT UNIQUE NOT NULL REFERENCES pattern (id),
    full_pattern_id BIGINT        NOT NULL REFERENCES pattern (id)
);

COMMIT;
