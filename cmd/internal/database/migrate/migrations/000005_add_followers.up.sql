CREATE TABLE IF NOT EXISTS followers (
  user_id UUID NOT NULL,
  follower_id UUID NOT NULL,
  created_at timestamp(0) with time zone NOT NULL,

  PRIMARY KEY (user_id, follower_id),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);