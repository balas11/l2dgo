
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL
);

insert into categories(name) values('Electronics');
insert into categories(name) values('Home Appliances');
insert into categories(name) values('Clothing');



CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    price numeric(10,2) NOT NULL,
    category_id integer REFERENCES categories(id)
);

insert into products(name, price, category_id) values('Sony Television', 45569.00, 1);
insert into products(name, price, category_id) values('Samsung Television', 42999.00, 1);
insert into products(name, price, category_id) values('Sony Home Theatre', 25000.00, 1);
insert into products(name, price, category_id) values('Philips Mixer', 3699.00, 2);
insert into products(name, price, category_id) values('Butterfly Grinder', 2799.00, 2);
insert into products(name, price, category_id) values('Arrow Full Sleeve', 3500.00, 3);
insert into products(name, price, category_id) values('Levis 511', 2599.00, 3);
