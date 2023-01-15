CREATE TYPE enumm AS ENUM ('Основное общее', 'Среднее общее', 'Среднее профессиональное', 'Высшее');
CREATE TABLE users
(
    id              serial       not null unique,
    phone_number    int          not null unique,
    password        varchar(255) not null,
    position        varchar(255) not null,
    experience      int          not null,
    education_level enumm        not null,
    description     varchar(255) not null
);

