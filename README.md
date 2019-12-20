# gojsonstruct

Converts json documents into struct 

## Example usage 


    {
        "person" : { 
            "name" : "John",
            "age" : 30, 
            "cars" : 
               "car1" : "BMW", 
               "car2" : "Tesla" 
            }, 
            "favouriteLanguages": ["Go","Java","JS"]
        },
        "computer" : {
            "name" : "Apple",
            "year" : 2019,
            "new" : true,
            "version": 10.1
        }
     }

## Result 

         type T struct {
	        Person struct {
		        Name string `json:"name"`
		        Age int `json:"age"`
		        Cars struct {
			       Car1 string `json:"car1"`
			       Car2 string `json:"car2"`
		        } `json:"cars"`
		        FavouriteLanguages []string `json:"favouriteLanguages"`
	        } `json:"person"`
	        Computer struct {
		        Name string `json:"name"`
		        Year int `json:"year"`
		        New bool `json:"new"`
		        Version float64 `json:"version"`
	        } `json:"computer"`
         }
  
