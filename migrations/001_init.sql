CREATE TABLE IF NOT EXISTS users(
    id UUID default gen_random_uuid(),
    email varchar(20),
    password_hash TEXT,
    created_at TIMESTAMP,
);

CREATED TABLE if NOT EXISTS cinemas(
    id serial,
    address TEXT,
);

CREATE TABLE IF NOT EXISTS halls(
    id UUID,
    cinema_id UUID,
    capacity INTEGER
);

CREATED TABLE IF NOT EXISTS movies(
    id UUID,
    title TEXT,
    description TEXT,
    duration interval,
    age_ration INTEGER
);

CREATE TABLE IF NOT EXISTS sessions(
    id UUID,
    movie_id UUID,
    hall_id UUID,
    start_time TIMESTAMP,
    price money
);

