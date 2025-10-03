import { Echo } from '../wailsjs/go/main/App';
import { GetArticles } from '../wailsjs/go/main/App';

(async function(){

    'use strict';

    const sources = document.querySelector("section#sources");
    const feed = document.querySelector("section#feed");
    const viewer = document.querySelector("section#viewer");
    
    const addFeedBtn = sources.querySelector("button#addFeedBtn");
    const dialogBtns = Array.from(document.querySelectorAll(".overlay button.close"));

    viewer.addEventListener("click", function(e){
        e.preventDefault();
        e.stopImmediatePropagation();
    }, true);

    function processArticles(articles){
        
        console.log("articles:", articles);
        const feedScroller = feed.querySelector("#feedScroller");

        const docFrag = document.createDocumentFragment();

        articles.forEach(article => {
            
            const articleTemplate = document.createElement("template");
            const articleTemplateContent = `

                <article class="piece" data-src="${article.url}" data-sentiment="${article.sentimentGroup}" data-selected="false">
    
                    <div class="articleContent">
    
                        <div class="sentimentIndicator"></div>
                        
                        <div class="content">
                            
                            <div class="topSection">
                                <h1>${article.title}</h1>
                                <div class="newIndicator" data-read="${article.alreadyRead}"></div>
                            </div>
                            <div class="bottomSection">
                                <span class="score">${article.sentimentScore * 100 | 0}% Positive</span>
                                <span class="lastUpdate">5 minutes ago</span>
                            </div>
    
                        </div>
    
                    </div>
    
                </article>
            
            `;

            articleTemplate.innerHTML = articleTemplateContent;

            const articleNode = articleTemplate.content.cloneNode(true);
            const articleEl = articleNode.querySelector("article");
            console.log("articleEl:", articleEl);

            articleEl.addEventListener("click", function(e){
                
                console.log(this);
    
                e.preventDefault();
                e.stopImmediatePropagation();
    
                Array.from(feed.querySelectorAll("article")).forEach(article => {
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
                    // console.error('Iframe error:', e);
                });
                
                viewer.querySelector("h1").dataset.active = "false";
                viewer.querySelector("iframe").dataset.active = "true";
    
            }, false);

            docFrag.appendChild(articleEl);
    
        });

        feedScroller.innerHTML = "";
        feedScroller.appendChild(docFrag);

    }

    addFeedBtn.addEventListener("click", function(){

        const addFeedOverlay = document.querySelector("#addFeed");
        addFeedOverlay.dataset.active = "true";

    }, false);

    dialogBtns.forEach(closeBtn => {

        closeBtn.addEventListener("click", function(e){
            e.preventDefault();
            e.stopImmediatePropagation();

            const overlayToClose = this.parentNode.parentNode;
            const formToReset = overlayToClose.querySelector("form");

            overlayToClose.dataset.active = "false";
            formToReset.reset();

        }, false);

    });

    GetArticles()
        .then(articles => processArticles(articles))
        .catch(err => {
            console.log("GetArticles err:", err);
        })
    ;

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
