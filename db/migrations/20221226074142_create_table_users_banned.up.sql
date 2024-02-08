create table users_banned
(
    id         int auto_increment,
    user_id int not null,
    status tinyint,
    reason text,
    start_date date null,
    end_date date null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null,
    deleted_at timestamp null,
    constraint users_banned_pk
        primary key (id)
);
