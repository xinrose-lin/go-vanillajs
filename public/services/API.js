// this module, exports this constant
// needs to be imported to app.js in order to be global 
export const API = {
    baseURL: "/api/", 
    getTopMovies: async () => {
        // API.fetch directly returns the Promise via fetch
        // to return the data directly; 
        // use keyword await -- why?
        return API.fetch("movies/top");
        // return await API.fetch("movies/top/");
    }, 
    getRandomMovies: async () => {
        return API.fetch("movies/random");
    }, 
    getMovieById: async (id) => {
        return API.fetch(`movies/${id}`);
    }, 
    // searchMovies: async (q, order, genre) => {
    //     return API.fetch(`movies/search/', {q, order, genre}); 
    // }, 
    fetch: async (serviceName, args) => {
        try {
            // const response = await fetch("/api/movies/top");
            console.log(API.baseURL + serviceName) 
            const response = await fetch(API.baseURL + serviceName); 
            console.log(response)
            const result = await response.json(); 
            return result; 
        } catch(e) {
            console.error(e); 
        }
    }


}
