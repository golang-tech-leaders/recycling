CREATE TABLE waste_type(
    id text,
    name text,
    description text,
    CONSTRAINT pk_waste_type PRIMARY KEY (id)
);

INSERT INTO waste_type(id, name) VALUES ('Стекло', 'Стекло');
INSERT INTO waste_type(id, name) VALUES ('Пластик', 'Одноразовый пластик');
INSERT INTO waste_type(id, name) VALUES ('Метал', 'Отходы из метала');
INSERT INTO waste_type(id, name) VALUES ('Одежда', 'Одежда');
INSERT INTO waste_type(id, name) VALUES ('Иное', 'Иное');
INSERT INTO waste_type(id, name) VALUES ('Опасное', 'Опасные отходы');
INSERT INTO waste_type(id, name) VALUES ('Батарейки', 'Батарейки, аккумуляторы, АКБ');
INSERT INTO waste_type(id, name) VALUES ('Лампочки', 'Лампы, лампы дневного света');
INSERT INTO waste_type(id, name) VALUES ('Техника', 'Бытовая техника');
INSERT INTO waste_type(id, name) VALUES ('ТетраПак', 'Тара/упаковка Tetra Pack');
INSERT INTO waste_type(id, name) VALUES ('Крышечки', 'Пластиковые крышечки');
INSERT INTO waste_type(id, name) VALUES ('Шины', 'Шины резиновые');

