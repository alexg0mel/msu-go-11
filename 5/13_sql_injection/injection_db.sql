CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `login` varchar(300) NOT NULL,
  UNIQUE KEY `login` (`login`)
);

INSERT INTO `users` ( `login`) VALUES ('user'), ('admin');
