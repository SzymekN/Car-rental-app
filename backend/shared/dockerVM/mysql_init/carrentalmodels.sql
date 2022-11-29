USE car_rental;
CREATE TABLE Vehicle (
  ID                 int(10) NOT NULL AUTO_INCREMENT, 
  registrationNumber varchar(7) NOT NULL UNIQUE, 
  brand              varchar(50) NOT NULL, 
  model              varchar(50) NOT NULL, 
  type               varchar(30) NOT NULL, 
  color              varchar(20) NOT NULL, 
  fuelConsumption    float NOT NULL, 
  dailyCost          int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE Employee (
  ID      int(10) NOT NULL AUTO_INCREMENT, 
  name    varchar(30) NOT NULL, 
  surname int(30) NOT NULL, 
  Salary  int(10) NOT NULL, 
  PESEL   varchar(255) NOT NULL, 
  userID  int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE Client (
  ID          int(10) NOT NULL AUTO_INCREMENT, 
  name        varchar(30) NOT NULL, 
  surname     varchar(30) NOT NULL, 
  PESEL       varchar(11) NOT NULL UNIQUE, 
  phoneNumber varchar(9) NOT NULL UNIQUE, 
  userID      int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE Rental (
  ID            int(10) NOT NULL AUTO_INCREMENT, 
  driverID      int(10), 
  clientiID     int(10) NOT NULL, 
  vehicleID     int(10) NOT NULL, 
  startDate     date NOT NULL, 
  endDate       date NOT NULL, 
  pickupAddress varchar(255), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE Repairs (
  ID             int(10) NOT NULL AUTO_INCREMENT, 
  vehicleID      int(10) NOT NULL, 
  cost           int(10) NOT NULL, 
  approved       bit(1), 
  notificationID int(10) UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE Salary (
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
CREATE TABLE `User` (
  ID       int(10) NOT NULL AUTO_INCREMENT, 
  email    varchar(50) NOT NULL UNIQUE, 
  password varchar(255) NOT NULL, 
  role     varchar(20) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
ALTER TABLE Rental ADD CONSTRAINT FKrental327656 FOREIGN KEY (vehicleID) REFERENCES Vehicle (ID);
ALTER TABLE Rental ADD CONSTRAINT FKrental599429 FOREIGN KEY (clientiID) REFERENCES Client (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati288469 FOREIGN KEY (vechicleID) REFERENCES Vehicle (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati740257 FOREIGN KEY (employeeID) REFERENCES Employee (ID);
ALTER TABLE notification ADD CONSTRAINT FKnotificati931400 FOREIGN KEY (clientID) REFERENCES Client (ID);
ALTER TABLE Repairs ADD CONSTRAINT FKrepairs887631 FOREIGN KEY (vehicleID) REFERENCES Vehicle (ID);
ALTER TABLE Employee ADD CONSTRAINT FKemployee40675 FOREIGN KEY (userID) REFERENCES `User` (ID);
ALTER TABLE Client ADD CONSTRAINT FKclient827663 FOREIGN KEY (userID) REFERENCES `User` (ID);
ALTER TABLE Rental ADD CONSTRAINT FKrental369251 FOREIGN KEY (driverID) REFERENCES Employee (ID);
ALTER TABLE Salary ADD CONSTRAINT FKsalary786655 FOREIGN KEY (employeeID) REFERENCES Employee (ID);
ALTER TABLE Repairs ADD CONSTRAINT FKrepairs288519 FOREIGN KEY (notificationID) REFERENCES notification (ID);
