package sql

const (
	CheckUserRegistered = `select id from users where deleted_at is null and name=$1`
	CheckUserRegisteredByEmail = `select id from users where deleted_at is null and email=$1`
	Login               = `select security_word,name from users where email=$1`
	Register            = `insert into users 
						   (name,email,security_word,full_name) 
						   values (uuid_generate_v4(),$1,$2,$3) returning id,name`

	GetAllTransactionsWithDate = `	select id,sender_wallet_id,recipient_wallet_id,amount,source_of_funds,description,updated_at  from transactions 
	where sender_wallet_id=$1 or recipient_wallet_id =$1 and  updated_at between $2 and $3 limit $4 offset $5;`
	GetAllTransactions = `	select id,sender_wallet_id,recipient_wallet_id,amount,source_of_funds,description,updated_at  from transactions 
	where sender_wallet_id=$1 or recipient_wallet_id =$1 limit $2 offset $3;`
)
