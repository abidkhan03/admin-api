BEGIN;

CREATE TABLE word
(
	id		BIGSERIAL		NOT	NULL PRIMARY KEY,
	word	VARCHAR(64)		NOT	NULL UNIQUE
);

CREATE TABLE pos
(
	id				BIGSERIAL	NOT	NULL PRIMARY KEY,

	category		INT	NOT	NULL,
	type			INT,
	degree			INT,
	gen				INT,
	num				INT,
	possessorpers	INT,
	possessornum	INT,
	person			INT,
	neclass			INT,
	nesubclass		INT,
	"case"			INT,
	polite			INT,
	mood			INT,
	tense			INT,
	punctenclose	INT,


	UNIQUE (category,type,degree,gen,num,possessorpers,possessornum,person,neclass,
		nesubclass,"case",polite,mood,tense,punctenclose)
);

CREATE TABLE word_pos
(
	word_id		BIGINT		NOT	NULL	REFERENCES word(id)	ON DELETE CASCADE,
	lemma_id	BIGINT		NOT	NULL	REFERENCES word(id)	ON DELETE CASCADE,
	pos_id		BIGINT		NOT	NULL	REFERENCES pos(id)	ON DELETE CASCADE,

	UNIQUE (word_id,lemma_id,pos_id)
);

CREATE TABLE "user" 
(
	id			SERIAL 			NOT NULL PRIMARY KEY,
    email       VARCHAR(64)     NOT NULL UNIQUE,
    password    VARCHAR(64)     NOT NULL
);

CREATE TABLE category
(
	id			SERIAL 			NOT NULL PRIMARY KEY,
	name        VARCHAR(64)     NOT NULL UNIQUE,
	parent_id 	INT 			REFERENCES category (id) ON DELETE CASCADE
);

CREATE TABLE "class" 
(
	id				SERIAL			NOT	NULL PRIMARY KEY,
	name 			VARCHAR(50)		NOT	NULL UNIQUE,
	word_id 		BIGINT			REFERENCES word(id) ON DELETE SET NULL,
	description		VARCHAR(500)
);

CREATE TABLE word_class
(
	word_id 	BIGINT		NOT	NULL	REFERENCES word(id)		ON DELETE CASCADE,
	class_id 	INT			NOT	NULL	REFERENCES "class"(id)	ON DELETE CASCADE,

	UNIQUE (word_id, class_id)
);

COMMIT;
