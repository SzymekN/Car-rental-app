DROP TABLE IF EXISTS car_rental.Pojazdy;
DROP TABLE IF EXISTS car_rental.Pracownicy;
DROP TABLE IF EXISTS car_rental.Klienci;
DROP TABLE IF EXISTS car_rental.Wypozyczenie;
DROP TABLE IF EXISTS car_rental.Naprawy;
DROP TABLE IF EXISTS car_rental.Wyplaty;
DROP TABLE IF EXISTS car_rental.Zgloszenia;

CREATE TABLE car_rental.Pojazdy (
  ID                    int(10) NOT NULL AUTO_INCREMENT, 
  `Numer rejestracyjny` varchar(7) NOT NULL UNIQUE, 
  Marka                 varchar(50) NOT NULL, 
  Model                 varchar(50) NOT NULL, 
  Typ                   varchar(30) NOT NULL, 
  Kolor                 varchar(20) NOT NULL, 
  Spalanie              float NOT NULL, 
  `Koszt dobowy`        int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;
  
CREATE TABLE car_rental.Pracownicy (
  ID       int(10) NOT NULL AUTO_INCREMENT, 
  Imie     varchar(30) NOT NULL, 
  Nazwisko int(30) NOT NULL, 
  Rola     varchar(20) NOT NULL, 
  Pensja   int(10) NOT NULL, 
  PESEL    varchar(255) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;

CREATE TABLE car_rental.Klienci (
  ID       int(10) NOT NULL AUTO_INCREMENT, 
  Imie     varchar(30) NOT NULL, 
  Nazwisko varchar(30) NOT NULL, 
  PESEL    varchar(11) NOT NULL UNIQUE, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;

CREATE TABLE car_rental.Wypozyczenie (
  ID                int(10) NOT NULL AUTO_INCREMENT, 
  KierowcaID        int(10), 
  KlienciID         int(10) NOT NULL, 
  PojazdyID         int(10) NOT NULL, 
  `Data początkowa` date NOT NULL, 
  `Data końcowa`    date NOT NULL, 
  `Miejsce odbioru` varchar(255), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;

CREATE TABLE car_rental.Naprawy (
  ID           int(10) NOT NULL AUTO_INCREMENT, 
  PojazdyID    int(10) NOT NULL, 
  Koszt        int(10) NOT NULL, 
  Zatwierdzone bit(1), 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;

CREATE TABLE car_rental.Wyplaty (
  PracownicyID int(10) NOT NULL, 
  Data         date NOT NULL, 
  Wartość      int(10) NOT NULL) CHARACTER SET UTF8;

CREATE TABLE car_rental.Zgloszenia (
  ID           int(10) NOT NULL AUTO_INCREMENT, 
  Opis         varchar(255) NOT NULL UNIQUE, 
  PracownicyID int(10), 
  KlienciID    int(10), 
  PojazdyID    int(10), 
  NaprawyID    int(10) NOT NULL, 
  PRIMARY KEY (ID)) CHARACTER SET UTF8;

ALTER TABLE car_rental.Wypozyczenie ADD CONSTRAINT FKWypozyczen167024 FOREIGN KEY (PojazdyID) REFERENCES car_rental.Pojazdy (ID);
ALTER TABLE car_rental.Wypozyczenie ADD CONSTRAINT FKWypozyczen329455 FOREIGN KEY (KlienciID) REFERENCES car_rental.Klienci (ID);
ALTER TABLE car_rental.Wyplaty ADD CONSTRAINT FKWyplaty514456 FOREIGN KEY (PracownicyID) REFERENCES car_rental.Pracownicy (ID);
ALTER TABLE car_rental.Zgloszenia ADD CONSTRAINT FKZgloszenia930423 FOREIGN KEY (PojazdyID) REFERENCES car_rental.Pojazdy (ID);
ALTER TABLE car_rental.Zgloszenia ADD CONSTRAINT FKZgloszenia768176 FOREIGN KEY (NaprawyID) REFERENCES car_rental.Naprawy (ID);
ALTER TABLE car_rental.Zgloszenia ADD CONSTRAINT FKZgloszenia756303 FOREIGN KEY (PracownicyID) REFERENCES car_rental.Pracownicy (ID);
ALTER TABLE car_rental.Zgloszenia ADD CONSTRAINT FKZgloszenia566055 FOREIGN KEY (KlienciID) REFERENCES car_rental.Klienci (ID);
ALTER TABLE car_rental.Naprawy ADD CONSTRAINT FKNaprawy277584 FOREIGN KEY (PojazdyID) REFERENCES car_rental.Pojazdy (ID);
ALTER TABLE car_rental.Wypozyczenie ADD CONSTRAINT FKWypozyczen924206 FOREIGN KEY (KierowcaID) REFERENCES car_rental.Pracownicy (ID);
