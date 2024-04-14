CREATE TABLE banner (
    id SERIAL PRIMARY KEY NOT NULL ,
    content jsonb NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    is_active boolean NOT NULL
);

CREATE TABLE tag (
    id SERIAL PRIMARY KEY NOT NULL
);

CREATE TABLE feature (
    id SERIAL PRIMARY KEY NOT NULL
);

CREATE TABLE banner_tag (
    id SERIAL PRIMARY KEY NOT NULL ,
    banner_id INTEGER REFERENCES banner(id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tag(id) ON DELETE CASCADE
);

CREATE TABLE banner_feature (
    id SERIAL PRIMARY KEY NOT NULL ,
    banner_id INTEGER REFERENCES banner(id) ON DELETE CASCADE,
    feature_id INTEGER REFERENCES feature(id) ON DELETE CASCADE
);