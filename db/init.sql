CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    publication_date DATE NOT NULL,
    shill_name TEXT NOT NULL,
    ticker TEXT NOT NULL,
    transaction_date DATE NOT NULL,
    transaction_type INT NOT NULL,
    shares INT NOT NULL,
    price_per_share FLOAT NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_hash TEXT NOT NULL,
    user_email TEXT NOT NULL
);

CREATE TABLE user_notifications (
    id SERIAL PRIMARY KEY,
    user_fk INT NOT NULL,
    shill_name TEXT NOT NULL
);

ALTER TABLE user_notifications
    ADD CONSTRAINT fk_user_notifs_user FOREIGN KEY (user_fk) REFERENCES users(id) ON DELETE CASCADE;

CREATE INDEX index_trades_on_shill ON trades USING btree(shill_name);
