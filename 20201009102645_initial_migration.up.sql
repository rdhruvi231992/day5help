CREATE Table users (
     id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid (),
     email TEXT NOT NULL UNIQUE
     
     --password_hash TEXT NOT NULL 
);