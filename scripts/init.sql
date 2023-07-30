create schema if not exists pixel_battle;
use pixel_battle;

create table if not exists pixel_battle.pixels
(
    pixel_row       int,
    pixel_col  int,
    color     varchar(255),
    player_id varchar(255) null,
    constraint pixels_pk
        primary key (pixel_row, pixel_col)
);
create table pixel_battle.history
(
    id      varchar(255)                        not null,
    pixel_col     int                           not null,
    pixel_row     int                           not null,
    color   varchar(255)                        not null,
    user_id varchar(255)                        not null,
    time    timestamp default current_timestamp not null,
    constraint history_pk
        primary key (id)
);

