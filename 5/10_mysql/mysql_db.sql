CREATE TABLE `students` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `fio` varchar(300) NOT NULL,
  `info` text NULL,
  `score` int NOT NULL
);

INSERT INTO `students` (`fio`, `info`, `score`) 
VALUES ('Vasily Romanov', 'company: Mail.ru Group', '10');
