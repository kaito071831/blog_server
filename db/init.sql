USE blogDB;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` text NOT NULL,
  `hash_password` text NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE posts (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` text NOT NULL,
  `body` text,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`)
    REFERENCES users(`id`)
)
