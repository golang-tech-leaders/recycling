CREATE TABLE waste_type(
    id text,
    name text,
    description text,
    CONSTRAINT pk_waste_type PRIMARY KEY (id)
);

INSERT INTO waste_type(id, name) VALUES ('стекло', 'стекло');
INSERT INTO waste_type(id, name) VALUES ('пластик', 'одноразовый пластик');
INSERT INTO waste_type(id, name) VALUES ('метал', 'отходы из метала');
INSERT INTO waste_type(id, name) VALUES ('одежда', 'одежда');
INSERT INTO waste_type(id, name) VALUES ('иное', 'иное');
INSERT INTO waste_type(id, name) VALUES ('опасное', 'опасные отходы');
INSERT INTO waste_type(id, name) VALUES ('батарейки', 'батарейки, аккумуляторы, АКБ');
INSERT INTO waste_type(id, name) VALUES ('лампочки', 'лампы, лампы дневного света');
INSERT INTO waste_type(id, name) VALUES ('техника', 'бытовая техника');
INSERT INTO waste_type(id, name) VALUES ('тетрапак', 'тара/упаковка tetra pack');
INSERT INTO waste_type(id, name) VALUES ('крышечки', 'пластиковые крышечки');
INSERT INTO waste_type(id, name) VALUES ('шины', 'шины резиновые');
