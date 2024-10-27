CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    score DOUBLE PRECISION NOT NULL
);
CREATE INDEX idx_leaderboard_country_score ON leaderboard (country, score DESC);
CREATE INDEX idx_leaderboard_state_score ON leaderboard (state, score DESC);
CREATE INDEX idx_leaderboard_score ON leaderboard (score DESC);