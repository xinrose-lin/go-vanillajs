import { API } from "../services/API.js"; 
import { MovieItemComponent } from "./MovieItem.js";
// creating web components: 
// 1. create html template 
// 2. create 'wrapper' javascript class; import html template 
// 3. customElements.define
// by extending from HTMLElements
// 4. import this to app.js (so that browser can recognise this component)
export class HomePage extends HTMLElement {
    // constructor() {
    //     fetch("/template/home-page.html")
    // }

    // render() - not from super class
    // allows us to dynamically render content 
    async render () {
        // get movies from backend
        const topMovies = await API.getTopMovies()
        // console.log(topMovies)
        //  top 10 returns the section
        // renderMoviesInList(topMovies, document.getElementById("top-10"))
        renderMoviesInList(topMovies, document.querySelector("#top-10 ul"))

        function renderMoviesInList(movies, ul){
            // treats what element that it is given, as the "ul"
            // clear list first
            ul.innerHTML = ""; 
            // movie refers to every element in list "movies"
            movies.forEach(movie => {
                console.log(movie)
                const li = document.createElement("li"); 
                // li.textContent = movie.title; 
                // ul.appendChild(li); 

                li.appendChild(new MovieItemComponent(movie));
                ul.appendChild(li); 
            }); 
        }
    }

    // a method to override from super class
    connectedCallback() {
        // if tempalte is not part of html page, can fetch it into this element
        
        const template = document.getElementById("template-home"); 
        // const content = template.content.clone
        // console.log(template)
        const content = template.content.cloneNode(true); // deep clone
        this.appendChild(content);

        this.render(); 
    }
}
// add definition, so that html knows how to call this
customElements.define("home-page", HomePage); 
// export class H
