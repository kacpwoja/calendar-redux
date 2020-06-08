CREATE TABLE events (
    ev_id UUID PRIMARY KEY,
    ev_date DATE NOT NULL,
    ev_time TIME NOT NULL,
    ev_name TEXT NOT NULL
);