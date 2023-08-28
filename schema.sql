create table if not exists segments(
    id serial primary key,
    name varchar(255),
    unique(name)
);

create table if not exists user_segments(
    id serial primary key,
    user_id int,
    segment_id int,
    constraint fk_segment
    foreign key(segment_id)
    references segments(id)
    on delete cascade
);