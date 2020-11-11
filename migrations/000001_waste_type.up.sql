CREATE TABLE waste_type(
    id text,
    name text,
    description text,
    CONSTRAINT pk_waste_type PRIMARY KEY (id)
);

INSERT INTO waste_type(id, name) VALUES ('type1', 'waste type 1');
INSERT INTO waste_type(id, name) VALUES ('type2', 'waste type 2');
INSERT INTO waste_type(id, name) VALUES ('type3', 'waste type 3');
INSERT INTO waste_type(id, name) VALUES ('type4', 'waste type 4');
INSERT INTO waste_type(id, name) VALUES ('type5', 'waste type 5');
INSERT INTO waste_type(id, name, description) VALUES ('type6', 'waste type 6', 'description 6');
INSERT INTO waste_type(id, name) VALUES ('type7', 'waste type 7');
INSERT INTO waste_type(id, name) VALUES ('type8', 'waste type 8');
INSERT INTO waste_type(id, name) VALUES ('type9', 'waste type 9');

