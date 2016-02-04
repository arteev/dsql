-- create sqlite db
--
-- Файл сгенерирован с помощью SQLiteStudio v3.0.7 в пн янв. 25 13:14:39 2016
--
-- Использованная кодировка текста: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Таблица: client
CREATE TABLE client (
    ID_CLIENT   INTEGER PRIMARY KEY AUTOINCREMENT,
    NAME_CLIENT TEXT
);

INSERT INTO client (ID_CLIENT, NAME_CLIENT) VALUES (1, 'Aleksey LTD');
INSERT INTO client (ID_CLIENT, NAME_CLIENT) VALUES (2, 'Ivan Co & Son');

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
