create schema if not exists pixel_battle;
use pixel_battle;

create table if not exists pixels
(
    pixel_row       int,
    pixel_col  int,
    color     varchar(255),
    player_id varchar(255) null,
    constraint pixels_pk
        primary key (pixel_row, pixel_col)
);
