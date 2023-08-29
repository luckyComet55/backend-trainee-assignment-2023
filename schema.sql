create table if not exists segments(
    name varchar primary key,
    audience_cvg int,
);

create table if not exists users(
    id int generated always as identity,
    primary key(id)
);

create table if not exists user_segments(
    id int generated always as identity,
    user_id int,
    segment_name varchar,
    primary key(id),
    constraint fk_segment
        foreign key(segment_name)
        references segments(name)
        on delete cascade,
    constraint fk_user
        foreign key(user_id)
        references users(id)
        on delete cascade
);