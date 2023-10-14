-- Create the table to store cron job runs
CREATE TABLE cron_runs (
        id SERIAL PRIMARY KEY,
        run_at TIMESTAMP,
        expression TEXT
);