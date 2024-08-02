CREATE TABLE pastry(
    name TEXT NOT NULL PRIMARY KEY,
    description TEXT,
    price INTEGER,
    unitofmeasure TEXT,
    quantityperpiece TEXT);

CREATE TABLE pastryorder(
    id INTEGER NOT NULL PRIMARY KEY,
    pastry TEXT,
    customer TEXT,
    quantity REAL,
    preferedDate TEXT,
    scheduledDate TEXT);

CREATE TABLE bakingschedule(
    pastry TEXT NOT NULL,
    quantity REAL,
    reserved REAL,
    readyDate TEXT,
    PRIMARY KEY(pastry, readyDate),
    FOREIGN KEY(pastry) REFERENCES pastry(name));

CREATE TABLE dayoff(
    id INTEGER NOT NULL PRIMARY KEY,
    day TEXT NOT NULL );

CREATE TABLE migration (
    name TEXT NOT NULL PRIMARY KEY,
    date TEXT);