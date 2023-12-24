CREATE DATABASE IF NOT EXISTS fullcycle;

use fullcycle;


CREATE TABLE IF NOT EXISTS categories (
    id varchar(36) NOT NULL,
    name varchar(20),
    description varchar(20),
    PRIMARY KEY(id)
);

REPLACE INTO categories values ("cf1dcd10-9a41-4be1-b54b-b4d7c630fe33", "name cat 1", "desc cat 1");
