-- Drop Tables in reverse order of creation cos 
-- there is a foreign key dependency on accounts table.
-- if we drop accounts table first, then we will not be able to 
-- drop the transfers table or entries as references will be gone.
DROP TABLE if exists "transfers";
DROP TABLE if exists "entries";
DROP TABLE if exists "accounts";