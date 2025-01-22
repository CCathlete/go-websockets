-- Wrapping the whole migration in a transaction to allow rollback.
    begin
;

-- Creating a function that sets an updated at field when creating a new row.
   create or replace function trigger_set_timestamp () returns trigger as $$
    begin new.updated_at = now()
;
return new
;
end
;
$$ language plpgsql
;

-- Creating the users table.
   create table users (
          id serial primary key,
          first_name varchar(50) not null unique,
          last_name varchar(50) not null unique,
          user_active boolean not null default false,
          access_level int not null default 3,
          email varchar(100) not null unique,
          password_hash varchar(255) not null,
          deleted_at timestamptz default null,
          created_at timestamptz default now(),
          updated_at timestamptz default now()
          )
;

   create trigger set_timestamp before
   update on users for each row
  execute procedure trigger_set_timestamp ()
;

-- Creating the preferences table.
   create table preferences (
          id serial primary key,
          name varchar(50) not null unique,
          preference varchar(255) not null unique,
          deleted_at timestamptz default null,
          created_at timestamptz default now(),
          updated_at timestamptz default now()
          )
;

   create trigger set_timestamp before
   update on preferences for each row
  execute procedure trigger_set_timestamp ()
;

-- Creating the remember_tokens table.
   create table remember_tokens (
          id serial primary key,
          user_id int not null references users (id) on delete cascade,
          remember_token varchar(100) not null unique,
          created_at timestamptz default now(),
          updated_at timestamptz default now()
          )
;

   create trigger set_timestamp before
   update on remember_tokens for each row
  execute procedure trigger_set_timestamp ()
;

-- Creating the sessions table.
   create table sessions (
          token varchar(255) primary key,
          data bytea not null,
          expiry timestamptz not null
          )
;

   create index sessions_expiry_idx on sessions (expiry)
;

   commit
;
