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
