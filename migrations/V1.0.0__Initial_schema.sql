CREATE TABLE IF NOT EXISTS public.expense
(
    id          BIGSERIAL PRIMARY KEY,   -- Unique ID for each expense (auto-incremented)
    "timestamp" TIMESTAMP      NOT NULL DEFAULT now(), -- Timestamp of the expense
    amount      NUMERIC(10, 2) NOT NULL  -- Amount of the expense (up to 10 digits, 2 after decimal)
);

CREATE TABLE IF NOT EXISTS public.tag
(
    id   SERIAL PRIMARY KEY,  -- Unique ID for each tag (auto-incremented)
    "name" VARCHAR(50) NOT NULL -- Name of the tag (max 50 characters)
);

CREATE TABLE IF NOT EXISTS public.expense_tag
(
    expense_id BIGINT REFERENCES expense (id) ON DELETE CASCADE, -- References to expense table
    tag_id     INT REFERENCES tag (id) ON DELETE CASCADE,        -- References to tag table
    PRIMARY KEY (expense_id, tag_id)                             -- Composite primary key (unique pair)
);
