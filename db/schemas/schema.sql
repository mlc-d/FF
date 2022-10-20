create table if not exists users (
    id integer unsigned not null auto_increment,
    nick varchar(25) not null unique,
    password varchar(100) not null,
    role_id tinyint unsigned not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    primary key (id)
);

create table if not exists topics (
    id tinyint unsigned not null auto_increment,
    short_name varchar(5) not null unique,
    name varchar(25) not null unique,
    created_by integer unsigned not null,
    created_at timestamp default current_timestamp,
    primary key (id)
);
