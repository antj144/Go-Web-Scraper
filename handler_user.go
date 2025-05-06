package main


import "net/http" 


func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string 'json:"name"'

	}
	decoder := json.NewDecoder(r.Bodyparameters)

	params := parparameters{}
	err := decoder.Decode(&params)
	of err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,

	})
		if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v", err))
		return

		}
	respondWithJSON(w, 201, databaseUserToUser)(user))
}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
		respondWithJSON(w, 200, databaseUserToUser)(user))
}