ALTER TABLE pastry DROP COLUMN name;
ALTER TABLE pastry DROP COLUMN description;
ALTER TABLE pastry DROP COLUMN unitofmeasure;

CREATE TABLE pastrytranslation (
    language CHAR(2) NOT NULL,
    pastryid INTEGER,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY(language, pastryid),
    FOREIGN KEY (pastryid) REFERENCES pastry(id));

INSERT INTO pastrytranslation (language, pastryid, name, description)
    VALUES ("hu", 1, "Biscuit", "Nem az az édes, ez sós... Vajas pogácsa jó?! Vajas pogácsa!"),
        ("en", 1, "Biscuit", "No, not the sweet kind but the fluffy savoury one."),
        ("hu", 2, "Focaccia", "Olasz olajos kenyér lángos... Feltét nélkül."),
        ("en", 2, "Focaccia", "Oily, bubbly deliciousness."),
        ("hu", 3, "Kenyér", "Sima kenyér. Semmi extra."),
        ("en", 3, "Bread", "Simple sourdough bread."),
        ("hu", 4, "English muffin", 'Nem, ez nem az édesség. <a href=\"https://www.google.com/search?client=firefox-b-d&q=english+muffin\">Nézz utána!</a>'),
        ("en", 4, "English muffin", 'Nope, not the sweet one. <a href=\"https://www.google.com/search?client=firefox-b-d&q=english+muffin\">Look it up!</a>'),
        ("hu", 5, "Kakaós csiga", "Kakaós és fel van tekerve."),
        ("en", 5, "Cocoa roll", "Made with cocoa and it's rolled up."),
        ("hu", 6, "Tortilla", "Mexikói lapos lángos. Kaja origami."),
        ("en", 6, "Tortilla", "For food origami."),
        ("hu", 7, "Bagel", "Te a kutyára gondolsz, de az beagle. Ez olyan mint a szalagos fánk csak sós."),
        ("en", 7, "Bagel", "You are thinking of the dog but that's beagle. This is like a donut but savoury.");

ALTER TABLE pastry ADD COLUMN unitofmeasure INTEGER;

UPDATE pastry SET unitofmeasure = 2 WHERE id = 1;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 2;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 3;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 4;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 5;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 6;
UPDATE pastry SET unitofmeasure = 1 WHERE id = 7;

CREATE TABLE unitofmeasuretranslation (
    language CHAR(2) NOT NULL,
    unitofmeasureid INTEGER,
    name TEXT NOT NULL,
    PRIMARY KEY(language, unitofmeasureid),
    FOREIGN KEY (unitofmeasureid) REFERENCES pastry(unitofmeasure));

INSERT INTO unitofmeasuretranslation (language, unitofmeasureid, name)
    VALUES ("hu", 1,"db"),
        ("en", 1,"pcs"),
        ("hu", 2,"kg"),
        ("en", 2,"kg");