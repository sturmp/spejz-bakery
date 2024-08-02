CREATE TABLE pastrytemp (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER,
    unitofmeasure TEXT,
    quantityperpiece TEXT);

CREATE TABLE pastryordertemp (
    id INTEGER NOT NULL PRIMARY KEY,
    pastryid INTEGER,
    customer TEXT,
    quantity REAL,
    preferedDate TEXT NOT NULL,
    scheduledDate TEXT NOT NULL,
    FOREIGN KEY(pastryid) REFERENCES pastrytemp(id));

CREATE TABLE bakingscheduletemp (
    pastryid INTEGER,
    quantity REAL,
    reserved REAL,
    readyDate TEXT,
    PRIMARY KEY(pastryid, readyDate),
    FOREIGN KEY(pastryid) REFERENCES pastrytemp(id));

INSERT INTO pastrytemp (name, description, price, unitofmeasure, quantityperpiece)
    SELECT name, description, price, unitofmeasure, quantityperpiece FROM pastry;

INSERT INTO pastryordertemp (pastryid, customer, quantity, preferedDate, scheduledDate)
    SELECT pastrytemp.id,
        pastryorder.customer,
        pastryorder.quantity,
        pastryorder.preferedDate,
        pastryorder.scheduledDate FROM pastryorder
            JOIN pastrytemp ON pastryorder.pastry = pastrytemp.name;

INSERT INTO bakingscheduletemp (pastryid, quantity, reserved, readyDate)
    SELECT pastrytemp.id,
        bakingschedule.quantity,
        bakingschedule.reserved,
        bakingschedule.readyDate FROM bakingschedule
            JOIN pastrytemp ON bakingschedule.pastry = pastrytemp.name;

DROP TABLE bakingschedule;
DROP TABLE pastryorder;
DROP TABLE pastry;

ALTER TABLE pastrytemp RENAME TO pastry;
ALTER TABLE pastryordertemp RENAME TO pastryorder;
ALTER TABLE bakingscheduletemp RENAME TO bakingschedule;