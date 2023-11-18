# GO_CYPHER


First Setup your Database



install golang migrate and create migration file using migrate create ....keywords
Copy the SQL code into the generated .up file and revert in the .down file

CREATE DATABASE on the docker instance running
CREATE a Makefile and add the above command in it

MIGRATE the SQL code to the DB with "MIGRATE PATH---" for up and down

GENERATE CRUD OPERATION GOLANG CODE FROM SQL
Go to sqlc.dev to install with homebrew

run sqlc init to generate sqlc.yml file, go to github and search for settings to get the settings file or in the sqlc.dev

run SQLC GENERATE after setting the sqlc.yaml file

CREATE Query FILE in the DB/QUERY folder for the migration files(github page and search for getting started)

then run SQLC GENERATE again...will create file in the SQLC folder config in the sqlc.yaml file

Create a test file in the generated app for file in the DB/SQLC folder i.e for account.sql.go you will create account_test.go to test the CRUD GOlang code in the account.sql.go file

CREATE main_test.go file to create the test server in the DB/SQLC folder

Create test function to test them(i.e func TestAccountCreation(t* testing.T))
install testify and use the require feature of it to test the TestAccountCreation return object

CREATE UTIL folder and CREATE random.go file to automatically generate username

Then use it in the test file and test the package

CREATE cypher.go file in DB/SQLC folder to run individual query and transaction, each query will do one operation and doesn't support transaction so we embed it in the Cypher Struct in the cypher.go file (this is called composition)


TRANSFER TRANSACTION HERE


Create Transfer API, to implement the transaction function as a transaction
The transaction API should also create entry for each of the transaction, then test the API


CREATE USER 
make owner in the ACCOUNT table a foreign key to the USERNAME field in the USER table
Allow user to have multiple account but of different currency 
To actualize  this add composite unique index to the ACCOUNT table i.e in index of ACCOUNT table 

indexes {
    ....others
    ((owner, currency), [unique])
}

generate a  add_users migration  up and down file and paste the sql code from db diagram 
create a user.sql file in the queries folder and create the appropriate query there

THEN run SQLC GENERATE or MAKE SQLC if you have got it in the MakeFile.

THEN, proceed to create a user_test.go file in the sqlc folder of db folder.

CREATE USER create and login API and test, remember not to return the users hashed password among the the response