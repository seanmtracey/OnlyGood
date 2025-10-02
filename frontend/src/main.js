import { Echo } from '../wailsjs/go/main/App';

(function(){

    'use strict';

    const viewer = document.querySelector("section#viewer");
    const articles = document.querySelectorAll("article");

    console.log(articles);

    viewer.addEventListener("click", function(e){
        e.preventDefault();
        e.stopImmediatePropagation();
    }, true);

    Array.from(articles).forEach(article => {

        article.addEventListener("click", function(e){
            
            console.log(this);

            e.preventDefault();
            e.stopImmediatePropagation();

            articles.forEach(article => {
                article.dataset.selected = "false";
            });

            this.dataset.selected = "true";

            // document.querySelector("iframe").src = this.dataset.src;

            const iframe = document.querySelector("iframe");

            const proxyURL = `/api/proxy?url=${encodeURIComponent(this.dataset.src)}`;
            iframe.src = proxyURL;
            
            // Optional: Add load handlers
            iframe.addEventListener('load', () => {
                console.log('Iframe loaded:', targetURL);
            });
            
            iframe.addEventListener('error', (e) => {
                console.error('Iframe error:', e);
            });

            document.querySelector("iframe").dataset.active = "true";

        }, false);

    });

    console.log("Ready.");

}());

// Setup the greet function
window.echo = function (value) {

    Echo(value)
        .then((result) => {
            // Update result with data back from App.Greet()
            console.log(result);
        })
        .catch((err) => {
            console.error(err);
        })
    ;

};
