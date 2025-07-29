# tranfers-system


# Internal Transfers API

A production-ready Golang REST API for managing internal bank account transfers. The system allows creation of accounts, querying account balances, and transferring funds between accounts. Data is persisted in PostgreSQL.


# Features
1.Create account with unique ID and balance
2.View account details (ID, balance)
3.Transfer funds from one account to another
4.PostgreSQL-backed persistence
5.Proper error handling and clean code structure

# Prerequisites
Go 1.20+
PostgreSQL 13+

# Setup PostgreSQL Database

CREATE DATABASE transfers_db;
\c transfers_db
CREATE TABLE accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(20,5) NOT NULL DEFAULT 0.0
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    source_account_id BIGINT NOT NULL,
    destination_account_id BIGINT NOT NULL,
    amount NUMERIC(15, 5) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (source_account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (destination_account_id) REFERENCES accounts(account_id)
);


# Environment Variables
Set up DB connection in db/postgres.go or use .env file if ENV support is added later.
Example DSN: 
POSTGRES_CONN_STRING="host=localhost port=5432 user=postgres password=password dbname= transfers_db sslmode=disable"


# Create Account (POST) 
URL : http://localhost:8080/api/accounts
Body :
    {
    "account_id": 2004,
    "balance": 1200.0
    }

# Transfer Funds (POST)
URL : http://localhost:8080/api/transfer
Body :
    {
  "from_account_id": 2002,
  "to_account_id": 2001,
  "amount": 50.00
    }

# Get Account (GET)
URL : http://localhost:8080/api/accounts/2002


# Assumptions

Account ID is provided by the client and must be unique.
Transfers between same accounts are not allowed.
Sufficient balance is checked before transferring funds.
Numeric fields are formatted with 5 decimal precision.