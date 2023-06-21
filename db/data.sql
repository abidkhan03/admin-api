BEGIN;

INSERT INTO "user"(email, password)
VALUES ('admin@spongeling.com', 'admin123');


INSERT INTO category(id, name, parent_id)
VALUES (1, 'Asking Questions', NULL),
       (2, 'Alternative Questions', 1);

SELECT setval(pg_get_serial_sequence('category', 'id'), max(id))
FROM category;

INSERT INTO pos(id, category, type)
VALUES (1, ASCII('C'), ASCII('C'));

SELECT setval(pg_get_serial_sequence('pos', 'id'), max(id))
FROM pos;

INSERT INTO word(id, word)
VALUES (1, '¿'),
       (2, '?');

SELECT setval(pg_get_serial_sequence('word', 'id'), max(id))
FROM word;

INSERT INTO token(id, pos_id, word_id, class_id)
VALUES (2, NULL, 1, NULL),
       (3, NULL, 2, NULL),
       (4, 1, NULL, NULL);

SELECT setval(pg_get_serial_sequence('token', 'id'), max(id))
FROM token;


INSERT INTO pattern(id, token_ids)
VALUES (1, '{2, 1, 3}'),
       (2, '{2, 1, 1, 1, 3}'),
       (3, '{2, 1, 4, 1, 3}'),
       (4, '{2, 1, 1, 1, 1, 4, 1, 3}');

SELECT setval(pg_get_serial_sequence('pattern', 'id'), max(id))
FROM pattern;


INSERT INTO category_example(category_id, rule, tip, phrase, pattern_id, full_pattern_id)
VALUES (1, 'Starts with ¿ and ends with ?', '', '¿de dónde eres?', 1, 2),
       (2, 'Starts with ¿ and ends with ?, and has a coordinating conjunction.', '',
        '¿quieres alguna bebida fría o caliente?', 3, 4);


COMMIT;
