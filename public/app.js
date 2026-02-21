// create global app object that we can call
// global functions we can call from client side app

// forms: sends to server, reloads page and send back
// but for client side app; 
// event.preventDefault() - to allow custom handling of form event, 
// instead of reloading page - causes lost of all existing states

// require .js behind API, 
// to get file from server
import { HomePage } from "./components/HomePage.js";
import { API } from "./services/API.js"; 
import "./components/AnimatedLoading.js"
import { MovieDetailsPage } from "./components/MovieDetailsPage.js"

// execute when DOM loaded; 
window.addEventListener("DOMContentLoaded", event => {
    document.querySelector("main").appendChild(new HomePage())
    document.querySelector("main").appendChild(new MovieDetailsPage())
});
window.app = {
    search: (event) => {
        event.preventDefault();
        // select an input element of type search
        const q = document.querySelector("input[type='search']")
    }, 
    api: API
}