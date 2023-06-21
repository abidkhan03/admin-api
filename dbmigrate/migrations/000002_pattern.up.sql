BEGIN;

CREATE TABLE token
(
    id			BIGSERIAL		NOT	NULL PRIMARY KEY,
    pos_id 		BIGINT	REFERENCES pos(id)		ON DELETE CASCADE,
    word_id 	BIGINT	REFERENCES word(id)		ON DELETE CASCADE,
    class_id	INT		REFERENCES "class"(id)	ON DELETE SET NULL,

    UNIQUE (pos_id,word_id,class_id)
);

CREATE TABLE pattern
(
    id		BIGSERIAL	NOT	NULL 	PRIMARY KEY,
    t1 		BIGINT		NOT	NULL	REFERENCES token(id),
    t2 		BIGINT	REFERENCES token(id),
    t3 		BIGINT	REFERENCES token(id),
    t4 		BIGINT	REFERENCES token(id),
    t5 		BIGINT	REFERENCES token(id),
    t6 		BIGINT	REFERENCES token(id),
    t7 		BIGINT	REFERENCES token(id),
    t8 		BIGINT	REFERENCES token(id),
    t9 		BIGINT	REFERENCES token(id),
    t10 	BIGINT	REFERENCES token(id),
    t11		BIGINT	REFERENCES token(id),
    t12 	BIGINT	REFERENCES token(id),
    t13 	BIGINT	REFERENCES token(id),
    t14 	BIGINT	REFERENCES token(id),
    t15 	BIGINT	REFERENCES token(id),
    t16 	BIGINT	REFERENCES token(id),
    t17 	BIGINT	REFERENCES token(id),
    t18 	BIGINT	REFERENCES token(id),
    t19 	BIGINT	REFERENCES token(id),
    t20 	BIGINT	REFERENCES token(id),

    UNIQUE (t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15,t16,t17,t18,t19,t20)
);

CREATE TABLE phrase
(
    id				BIGSERIAL		NOT	NULL PRIMARY KEY,
    text 			VARCHAR(255)	NOT	NULL UNIQUE,
    rule 			VARCHAR(255),
    tip 			VARCHAR(255),
    category_id 	BIGINT  NOT NULL    REFERENCES category(id) ON DELETE CASCADE,
    pattern_id 		BIGINT	NOT NULL    REFERENCES pattern(id)  ON DELETE CASCADE
);

COMMIT;
