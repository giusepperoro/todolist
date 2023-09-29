create table  if not exists todolist (
    id   SERIAL PRIMARY KEY,
    title varchar(255),
    description varchar(2000),
    date_time timestamp,
    done boolean
);