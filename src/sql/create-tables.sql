CREATE DATABASE vacancies WITH ENCODING='UTF8' 
LC_CTYPE='en_US.UTF-8' LC_COLLATE='en_US.UTF-8' 
OWNER=edwica_root TEMPLATE=template0 CONNECTION 
LIMIT=-1

create table city (
    id_hh integer not null unique,
    id_edwica integer not null unique,
    name text not null unique,

    primary key(name)
);

create table position(
    id integer not null,
    name text not null,
    other_names text default "",
    level smallint default 0,
    parent_id integer default 0,
    parsed boolean default false,

    primary key (id)
);

create table currency (
    code varchar(3) not null,
    abbr varchar(10) not null,
    name varchar(50) not null,
    rate double precision not null,

    primary key (code)

);

create table vacancy(
    id SERIAL,
    hh_id INTEGER NOT NULL,
    hh_url VARCHAR(100) NOT NULL,
    name VARCHAR(150) NOT NULL,
    city_id INTEGER NOT NULL,
    position_id INTEGER NOT NULL,
    hh_prof_areas TEXT,
    hh_specs TEXT,
    experience TEXT NOT NULL,
    salary_from FLOAT,
    salary_to FLOAT,
    key_skills TEXT,
    vacancy_date TIMESTAMP NOT NULL,
    parsing_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (hh_id),
    FOREIGN KEY (city_id) REFERENCES city (id_edwica),
    FOREIGN KEY (position_id) REFERENCES position (id)
);

