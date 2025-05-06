package main


import "net/http" 


func (apiCfg *apiConfig)handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Name string 'json:"name"'
		URL string 'json:"url"'

	}
	decoder := json.NewDecoder(r.Bodyparameters)

	params := parparameters{}
	err := decoder.Decode(&params)
	of err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		URL: params.URL,
		UserID: user.ID,

	})
		if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed: %v", err))
		return

		}
		respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig)handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds, err := apiCfg.DB.GetFeeds(r.Context()){
		if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldnt get feed: %v", err))
		return

		}
		respondWithJSON(w, 201, databaseFeedToFeeds(feeds))
}
