# Database
Rename the  `docker-compose-dev.yml` to `docker-compose.yml` and update the values for the passwords

If you don't have `Docker` installed in your local machine follow the steps [here](https://docs.docker.com/engine/install/) to install

Execute `docker-compose up` this will spin up the MySQL Database.

# Tables
This statement should only be run once, this will create the users table

```
CREATE TABLE `users_db`.`users`(
	`id` BIGINT(24) NOT NULL auto_increment,
    `first_name` VARCHAR(45) NULL,
    `last_name` VARCHAR(45) NULL,
    `email` VARCHAR(45) NOT NULL,
    `date_created` VARCHAR(45) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC)
)
```