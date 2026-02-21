class AnimatedLoading extends HTMLElement {
    constructor() {
        super(); 
    }
    // only runs when first loaded 
    connectedCallback() {
        // access custom attributes
        // via dataset
        const elements = this.dataset.elements; 
        const width = this.dataset.width
        const height = this.dataset.height
        // this.getAttribute()
        for (let i=0; i<elements; i++) {
            const wrapper = document.createElement("div");
            wrapper.classList.add("loading-wave"); 
            wrapper.style.width = width; 
            wrapper.style.height = height;
            wrapper.style.margin = "10px"; 
            wrapper.style.display = "inline-block"; 
            this.appendChild(wrapper)
        }
    }
    // re-render when attributes change? 
}
customElements.define("animated-loading", AnimatedLoading)