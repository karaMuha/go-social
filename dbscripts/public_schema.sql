CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email CITEXT UNIQUE NOT NULL,
  username TEXT UNIQUE NOT NULL,
  user_password TEXT NOT NULL,
  created_at timestamp(0) with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title text NOT NULL,
  user_id UUID NOT NULL,
  content text NOT NULL,
  updated_at timestamp(0) with time zone NOT NULL,
  created_at timestamp(0) with time zone NOT NULL
);

ALTER TABLE posts ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id);