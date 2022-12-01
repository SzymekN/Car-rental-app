CREATE TABLE vehicle (
  ID                  int(10) NOT NULL AUTO_INCREMENT, 
  registration_number varchar(7) NOT NULL UNIQUE, 
  brand               varchar(50) NOT NULL, 
  model               varchar(50) NOT NULL, 
  type                varchar(30) NOT NULL, 
  color               varchar(20) NOT NULL, 
  fuel_consumption    float NOT NULL, 
  daily_cost          int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE employee (
  ID      int(10) NOT NULL AUTO_INCREMENT, 
  name    varchar(30) NOT NULL, 
  surname int(30) NOT NULL, 
  salary  int(10) NOT NULL, 
  PESEL   varchar(255) NOT NULL, 
  user_id int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE client (
  ID           int(10) NOT NULL AUTO_INCREMENT, 
  name         varchar(30) NOT NULL, 
  surname      varchar(30) NOT NULL, 
  PESEL        varchar(11) NOT NULL UNIQUE, 
  phone_number varchar(9) NOT NULL UNIQUE, 
  user_id      int(10) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE rental (
  ID             int(10) NOT NULL AUTO_INCREMENT, 
  start_date     date NOT NULL, 
  end_date       date NOT NULL, 
  pickup_address varchar(255), 
  driver_id      int(10), 
  client_id      int(10) NOT NULL, 
  vehicle_id     int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE repairs (
  ID              int(10) NOT NULL AUTO_INCREMENT, 
  cost            int(10) NOT NULL, 
  approved        bit(1), 
  vehicle_id      int(10) NOT NULL, 
  notification_id int(10) UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE salary (
  `date`      date NOT NULL, 
  amount      int(10) NOT NULL, 
  employee_id int(10) NOT NULL) CHARACTER SET UTF8;
CREATE TABLE notification (
  ID          int(10) NOT NULL AUTO_INCREMENT, 
  description varchar(255) NOT NULL UNIQUE, 
  employee_id int(10) NOT NULL, 
  client_id   int(10), 
  vechicle_id int(10), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
CREATE TABLE `user` (
  ID       int(10) NOT NULL AUTO_INCREMENT, 
  email    varchar(50) NOT NULL UNIQUE, 
  password varchar(255) NOT NULL, 
  role     varchar(20) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
ALTER TABLE rental ADD CONSTRAINT FKrental154746 FOREIGN KEY (vehicle_id) REFERENCES vehicle (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE rental ADD CONSTRAINT FKrental590843 FOREIGN KEY (client_id) REFERENCES client (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE notification ADD CONSTRAINT FKnotificati466077 FOREIGN KEY (vechicle_id) REFERENCES vehicle (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE notification ADD CONSTRAINT FKnotificati174904 FOREIGN KEY (employee_id) REFERENCES employee (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE notification ADD CONSTRAINT FKnotificati402236 FOREIGN KEY (client_id) REFERENCES client (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE repairs ADD CONSTRAINT FKrepairs601556 FOREIGN KEY (vehicle_id) REFERENCES vehicle (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE employee ADD CONSTRAINT FKemployee939388 FOREIGN KEY (user_id) REFERENCES `user` (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE client ADD CONSTRAINT FKclient245213 FOREIGN KEY (user_id) REFERENCES `user` (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE rental ADD CONSTRAINT FKrental398356 FOREIGN KEY (driver_id) REFERENCES employee (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE salary ADD CONSTRAINT FKsalary843084 FOREIGN KEY (employee_id) REFERENCES employee (ID) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE repairs ADD CONSTRAINT FKrepairs934942 FOREIGN KEY (notification_id) REFERENCES notification (ID) ON DELETE CASCADE ON UPDATE CASCADE;
