CREATE TABLE "events" (
    id serial PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    active BOOLEAN NOT NULL
);

INSERT INTO "events" (name, active) values ('Txalaparta', true), ('Kalejira', false);
