USE car_rental;
CREATE TABLE vechicles (
  ID                    int(10) NOT NULL AUTO_INCREMENT, 
  `registration number` varchar(7) NOT NULL UNIQUE, 
  brand                 varchar(50) NOT NULL, 
  model                 varchar(50) NOT NULL, 
  type                  varchar(30) NOT NULL, 
  color                 varchar(20) NOT NULL, 
  `fuel consumption`    float NOT NULL, 
  `daily cost`          int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE employees (
  ID      int(10) NOT NULL AUTO_INCREMENT, 
  name    varchar(30) NOT NULL, 
  surname int(30) NOT NULL, 
  salary  int(10) NOT NULL, 
  PESEL   varchar(255) NOT NULL, 
  userID  int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE clients (
  ID             int(10) NOT NULL AUTO_INCREMENT, 
  name           varchar(30) NOT NULL, 
  surname        varchar(30) NOT NULL, 
  PESEL          varchar(11) NOT NULL UNIQUE, 
  `phone number` int(9) NOT NULL UNIQUE, 
  userID         int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE rental (
  ID               int(10) NOT NULL AUTO_INCREMENT, 
  driverID         int(10), 
  clientiID        int(10) NOT NULL, 
  vechicleID       int(10) NOT NULL, 
  `start date`     date NOT NULL, 
  `end date`       date NOT NULL, 
  `pickup address` varchar(255), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE repairs (
  ID             int(10) NOT NULL AUTO_INCREMENT, 
  vechicleID     int(10) NOT NULL, 
  cost           int(10) NOT NULL, 
  approved       bit(1), 
  notificationID int(10) UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE salaries (
  employeeID int(10) NOT NULL, 
  `date`     date NOT NULL, 
  amount     int(10) NOT NULL) CHARACTER SET UTF8;
CREATE TABLE notification (
  ID          int(10) NOT NULL AUTO_INCREMENT, 
  description varchar(255) NOT NULL UNIQUE, 
  employeeID  int(10) NOT NULL, 
  clientID    int(10), 
  vechicleID  int(10), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE users (
  ID       int(10) NOT NULL AUTO_INCREMENT, 
  email    varchar(50) NOT NULL UNIQUE, 
  password varchar(255) NOT NULL, 
  role     varchar(20) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
ALTER TABLE rental ADD CONSTRAINT FKrental235275 FOREIGN KEY (vechicleID) REFERENCES vechicles (ID);
ALTER TABLE rental ADD CONSTRAINT FKrental69978 FOREIGN KEY (clientiID) REFERENCES clients (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati576117 FOREIGN KEY (vechicleID) REFERENCES vechicles (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati879721 FOREIGN KEY (employeeID) REFERENCES employees (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati738006 FOREIGN KEY (clientID) REFERENCES clients (ID);
ALTER TABLE repairs ADD CONSTRAINT FKrepairs521027 FOREIGN KEY (vechicleID) REFERENCES vechicles (ID);
ALTER TABLE employees ADD CONSTRAINT FKemployees371719 FOREIGN KEY (userID) REFERENCES users (ID);
ALTER TABLE clients ADD CONSTRAINT FKclients137689 FOREIGN KEY (userID) REFERENCES users (ID);
ALTER TABLE rental ADD CONSTRAINT FKrental982359 FOREIGN KEY (driverID) REFERENCES employees (ID);
ALTER TABLE salaries ADD CONSTRAINT FKsalaries155533 FOREIGN KEY (employeeID) REFERENCES employees (ID);
ALTER TABLE repairs ADD CONSTRAINT FKrepairs288519 FOREIGN KEY (notificationID) REFERENCES notification (ID);
