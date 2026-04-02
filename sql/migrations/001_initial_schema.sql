create extension if not exists "pgcrypto";

create table users
(
    id           uuid primary key      default gen_random_uuid(),
    username     varchar(32)  not null unique,
    display_name varchar(100) not null,
    bio          text,
    avatar_url   text,
    created_at   timestamptz  not null default now(),
    updated_at   timestamptz  not null default now()
);

create table user_follows
(
    follower_user_id uuid        not null references users (id) on delete cascade,
    followee_user_id uuid        not null references users (id) on delete cascade,
    created_at       timestamptz not null default now(),
    primary key (follower_user_id, followee_user_id),
    check (follower_user_id <> followee_user_id)
);

create table venues
(
    id         uuid primary key      default gen_random_uuid(),
    name       varchar(200) not null,
    city       varchar(100) not null,
    country    varchar(100) not null,
    created_at timestamptz  not null default now()
);

create table shows
(
    id          uuid primary key      default gen_random_uuid(),
    title       varchar(255) not null,
    description text,
    created_at  timestamptz  not null default now(),
    updated_at  timestamptz  not null default now()
);

create table productions
(
    id              uuid primary key     default gen_random_uuid(),
    show_id         uuid        not null references shows (id) on delete cascade,
    venue_id        uuid        not null references venues (id) on delete restrict,
    production_type varchar(50) not null,
    opening_date    date,
    closing_date    date,
    is_current      boolean     not null default true,
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now()
);

create table posts
(
    id                uuid primary key     default gen_random_uuid(),
    author_user_id    uuid        not null references users (id) on delete cascade,
    production_id     uuid        references productions (id) on delete set null,
    parent_post_id    uuid references posts (id) on delete cascade,
    body              text        not null,
    contains_spoilers boolean     not null default false,
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create table reviews
(
    id                uuid primary key     default gen_random_uuid(),
    author_user_id    uuid        not null references users (id) on delete cascade,
    production_id     uuid        not null references productions (id) on delete cascade,
    attendance_date   date,
    rating            numeric(2, 1),
    headline          varchar(255),
    body              text,
    contains_spoilers boolean     not null default false,
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    check (rating is null or (rating >= 0 and rating <= 5))
);

create table post_likes
(
    user_id    uuid        not null references users (id) on delete cascade,
    post_id    uuid        not null references posts (id) on delete cascade,
    created_at timestamptz not null default now(),
    primary key (user_id, post_id)
);

create table review_likes
(
    user_id    uuid        not null references users (id) on delete cascade,
    review_id  uuid        not null references reviews (id) on delete cascade,
    created_at timestamptz not null default now(),
    primary key (user_id, review_id)
);

create index idx_posts_author_created_at
    on posts (author_user_id, created_at desc);

create index idx_posts_production_created_at
    on posts (production_id, created_at desc);

create index idx_reviews_author_created_at
    on reviews (author_user_id, created_at desc);

create index idx_reviews_production_created_at
    on reviews (production_id, created_at desc);

create index idx_user_follows_followee
    on user_follows (followee_user_id);

create index idx_productions_is_current
    on productions (is_current);

create index idx_productions_venue_id
    on productions (venue_id);

create index idx_shows_title
    on shows (title);