CREATE TABLE IF NOT EXISTS posts (
  id UUID PRIMARY KEY uuid_generate_v4(),
  title text NOT NULL,
  user_id UUID NOT NULL,
  content text NOT NULL,
  created_at timestamp(0) with time zone NOT NULL
);

ALTER TABLE posts ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id);