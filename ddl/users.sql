CREATE TABLE IF NOT EXISTS `users` (
     `id` INT(255) NOT NULL AUTO_INCREMENT,
     `first_name` VARCHAR(255) NOT NULL,
     `last_name` VARCHAR(255) NOT NULL,
     `age` INT NOT NULL,
     `address` VARCHAR(255) NOT NULL,
     `city` VARCHAR(255) NOT NULL,
     `country` VARCHAR(255) NOT NULL,
     `phone_number` VARCHAR(255) NOT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB;

INSERT IGNORE INTO users (id, first_name, last_name, age, address, city, country, phone_number) VALUES (1, 'John', 'Smith', 30, '12 Abc Street', 'Some Place', 'Some Where', '+0 123 456 789');
INSERT IGNORE INTO users (id, first_name, last_name, age, address, city, country, phone_number) VALUES (2, 'Jane', 'Smith', 20, '34 Def Road', 'Another Place', 'Another Where', '+0 987 654 321');
