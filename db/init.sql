CREATE DATABASE banner_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE banner {
    id uuid  PRIMARY KEY DEFAULT uuid_generate_v4()
};

