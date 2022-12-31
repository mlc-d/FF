create table if not exists roles
(
    id   tinyint not null auto_increment,
    name varchar(20),
    primary key (id)
);

create table if not exists users
(
    id         integer unsigned not null auto_increment,
    nick       varchar(25)      not null unique,
    password   varchar(100)     not null,
    role_id    tinyint unsigned not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    primary key (id)
);

create table if not exists topics
(
    id              tinyint unsigned not null auto_increment,
    short_name      varchar(5)       not null unique,
    name            varchar(25)      not null unique,
    thumbnail_url   varchar(50)      not null unique,
    is_nsfw         boolean          not null default false,
    maximum_threads tinyint          not null default 64,
    created_by      integer unsigned not null,
    created_at      timestamp                 default current_timestamp,
    primary key (id)
);

create table if not exists media
(
    id             integer unsigned not null auto_increment,
    hash           varchar(50)      not null unique,
    extension      varchar(5)       not null,
    created_at     timestamp                 default current_timestamp,
    is_blacklisted boolean          not null default false,
    primary key (id)
);

create table if not exists threads
(
    id         integer unsigned not null auto_increment,
    topic_id   tinyint unsigned not null,
    user_id    integer unsigned not null,
    hash       varchar(25)      not null unique,
    title      varchar(30)      not null,
    body       varchar(2000)    not null,
    media_id   integer unsigned not null,
    sticky     boolean          not null,
    created_at timestamp default current_timestamp,
    primary key (id)
);

create table if not exists comments
(
    id         integer unsigned not null auto_increment,
    thread_id  integer unsigned not null,
    user_id    integer unsigned not null,
    tag        varchar(7)       not null unique,
    body       varchar(2000),
    unique_id  varchar(3),
    is_op      boolean          not null default false,
    color      tinyint          not null,
    created_at timestamp                 default current_timestamp,
    primary key (id)
);
