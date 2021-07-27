**Running the backend**
 Ensure Docker is installed, (Run `docker -v`) then run `docker compose up`.

 The API should be exposed at localhost:8080

 # User Auth

 *User Schema*
 ```
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"Email,omitempty" bson:"Email,omitempty"`
	Name     string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Type     string             `json:"Type,omitempty" bson:"Type,omitempty"`
	Password []byte             `json:"Password,omitempty" bson:"Password,omitempty"`
}
 ```
 
 *POST this to localhost:8080/register*
 ```
 {
    "Email":"a@a.com",
    "Name":"Dustin",
    "Type":"Manager",
    "Password":"passwd"
}
 ```
 
 *POST this to localhost:8080/login*
 ```
 {
    "Email":"a@a.com",
    "Password":"passwd"
}
 ```
 
# Data Processing

 *POST this to localhost:8080/calc*
 ```
 {
     "Step": "1",
     "Product": "Wine", 
     "Amount": "50", 
     "At": "25832", 
     "To": "38736"
 }
 ```

  *POST this to localhost:8080/evaluate*
 ```
 {
     "Farm": "33175", 
     "Process": "25832", 
     "Retail": "38736", 
     "Product": "Wine", 
     "Amount": "100"
 }
 ```