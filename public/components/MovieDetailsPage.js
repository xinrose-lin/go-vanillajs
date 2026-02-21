// define custom html element (custom.Element)
// create custom html element (extend from HTMLElement)
// -- to get movie details - via render()
// -- 

// export class: can refer this class via js
import { API } from "../services/API.js"

// customElements.define: 
export class MovieDetailsPage extends HTMLElement {
    id = null
    movie = null

    async render() {
        try {
            this.movie = await API.getMovieById(this.id)
        } catch {
            // error
            alert("Movie does not exist"); // TODO replace alert
            return;
        }
        const template = document.getElementById("template-movie-details")
        const content = template.content.cloneNode(true); 
        this.appendChild(content)

        this.querySelector("h2").textContent = this.movie.title;
        this.querySelector("h3").textContent = this.movie.tagline;
        // innerHTML - opens door for cross site scripting; textContent is faster

    }

    // connected call back is not async function
    connectedCallback() {
        this.id = 14;                                                                                                                                                                                                                                
        // async () => await API.getMovieById(this.movie)
        this.render(); 
    }
}           
//                                                                                                                       
customElements.define("movie-details-page", MovieDetailsPage)