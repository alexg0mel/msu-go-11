CREATE TABLE "students" (
  "id" serial NOT NULL,
  "fio" character varying(300) NOT NULL,
  "info" text NULL,
  "score" integer NOT NULL
);

INSERT INTO `students` (`fio`, `info`, `score`)
VALUES ('Vasily Romanov', 'company: Mail.ru Group', '10');
