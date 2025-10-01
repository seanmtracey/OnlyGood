import { Echo } from '../wailsjs/go/main/App';

(function(){

    'use strict';

    const articles = document.querySelectorAll("article");

    console.log(articles);

    Array.from(articles).forEach(article => {

        article.addEventListener("click", function(e){
            
            console.log(this);

            e.preventDefault();
            e.stopImmediatePropagation();

            articles.forEach(article => {
                article.dataset.selected = "false";
            });

            this.dataset.selected = "true";

            document.querySelector("iframe").src = this.dataset.src;
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
