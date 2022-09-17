ALTER TABLE IF EXISTS history_turn DROP CONSTRAINT IF EXISTS fk_winning_card;
DROP TABLE IF EXISTS history_played_card;
DROP TABLE IF EXISTS history_turn;
DROP TABLE IF EXISTS history_room;
DROP TABLE IF EXISTS history_player;
DROP TABLE IF EXISTS card;
DROP TABLE IF EXISTS deck;

CREATE TABLE deck (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255),
    author VARCHAR(255),
    selected_by_default BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE card (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    deck_id INTEGER NOT NULL REFERENCES deck(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    is_black_card BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE history_player (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255)
);

CREATE TABLE history_room (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    code VARCHAR(6),
    password VARCHAR(255),
    started_at TIMESTAMP NOT NULL DEFAULT NOW(),

    zen_mode BOOLEAN NOT NULL DEFAULT FALSE,
    max_turns INTEGER NULL DEFAULT NULL
);

CREATE TABLE history_turn (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    room_id INTEGER NOT NULL REFERENCES history_room(id),
    turn INTEGER NOT NULL,

    judge INTEGER NOT NULL REFERENCES history_player,
    winning_card INTEGER NULL DEFAULT NULL,

    UNIQUE (room_id, turn)
);

CREATE TABLE history_played_card (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    turn_id INTEGER NOT NULL REFERENCES history_turn(id),
    player_id INTEGER NOT NULL REFERENCES history_player(id),

    card_id INTEGER NOT NULL REFERENCES card(id),

    UNIQUE (turn_id, player_id),
    UNIQUE (turn_id, card_id)
);

ALTER TABLE history_turn ADD CONSTRAINT fk_winning_card FOREIGN KEY (winning_card) REFERENCES history_played_card(id) ON DELETE SET NULL;