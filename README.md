# Database
Rename the  `docker-compose-dev.yml` to `docker-compose.yml` and update the values for the passwords

If you don't have `Docker` installed in your local machine follow the steps [here](https://docs.docker.com/engine/install/) to install

Execute `docker-compose up` this will spin up the MySQL Database.

# Tables
This statement should only be run once, this will create the users table

```
CREATE TABLE `users` (
  `id` bigint(24) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(45) NOT NULL,
  `last_name` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `date_created` datetime NOT NULL,
  `status` varchar(45) NOT NULL,
  `password` varchar(65) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1
```