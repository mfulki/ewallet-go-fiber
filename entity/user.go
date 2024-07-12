package entity

type User struct {
	Id       int    
	Name     string 
	FullName string 
	Email    string 
	Password string 
	Wallet   *Wallet    
}
