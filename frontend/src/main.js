import { WindowToggleMaximise } from '../wailsjs/runtime/runtime'
import { GetArticles, Echo } from '../wailsjs/go/app/App'
import { ListFeeds, AddFeed } from '../wailsjs/go/feeds/Feeds'

(async function(){

    'use strict';

    const sources = document.querySelector("section#sources");
    const feed = document.querySelector("section#feed");
    const viewer = document.querySelector("section#viewer");
    const stretchWindow = document.querySelector("#stretchWindow");
    
    const addFeedBtn = sources.querySelector("button#addFeedBtn");
    const dialogBtns = Array.from(document.querySelectorAll(".overlay button.close"));

    const addFeedOverlay = document.querySelector(".overlay#addFeed");

    viewer.addEventListener("click", function(e){
        e.preventDefault();
        e.stopImmediatePropagation();
    }, true);

    function processArticles(articles){
        
        // console.log("articles:", articles);
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

    function processFeeds(feeds){

        console.log(feeds);

        const sources = document.querySelector("section#sources");
        const docFrag = document.createDocumentFragment();

        feeds.forEach(thisFeed => {

            if(!thisFeed.icon){
                thisFeed.icon = "/src/assets/images/rss_icon.png";
            }

            const feedItemTemplate = document.createElement("template");
            const feedItemTemplateContent = `
                <li data-hash="${thisFeed.hash}" data-src="${thisFeed.url}" data-name="${thisFeed.name}">
                    <img src="${thisFeed.icon}" />
                    <p>${thisFeed.name}</p>
                    <span data-sentiment="positive"></span>
                </li>
            `;

            feedItemTemplate.innerHTML = feedItemTemplateContent;

            const feedItemNode = feedItemTemplate.content.cloneNode(true);
            const feedItemEl = feedItemNode.querySelector("li");
            console.log("feedItemEl:", feedItemEl);

            feedItemEl.addEventListener("click", function(e){
                e.preventDefault();
                e.stopImmediatePropagation();

                console.log(this);

                feed.querySelector("#feedTitle h1").textContent = this.dataset.name;

            });

            docFrag.appendChild(feedItemEl);

        });

        const sourcesListEl = sources.querySelector("ol");

        sourcesListEl.innerHTML = "";
        sourcesListEl.appendChild(docFrag);

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

    addFeedOverlay.querySelector("form").addEventListener("submit", function(e){
        
        e.preventDefault();
        e.stopImmediatePropagation();

        const rssAddr = this.querySelector(`input[type="url"]`).value;

        console.log(rssAddr);

        AddFeed({
                name : "Test Feed",
                url : rssAddr
            })
            .then(result => {
                console.log(result);
                return ListFeeds();
            })
            .then(feeds => {
                processFeeds(feeds);
                addFeedOverlay.dataset.active = "false";
                this.reset();
            })
            .catch(err => {
                console.log("AddFeed err:", err);
            })
        ;

    })

    stretchWindow.addEventListener("click", function(e){

        e.preventDefault();
        e.stopImmediatePropagation();

        console.log("Stretch!");

        WindowToggleMaximise();

    }, false);

    ListFeeds()
        .then(feeds => processFeeds(feeds))
        .catch(err => {
            console.log("ListFeeds err:", err);
        })
    ;

    /*GetArticles()
        .then(articles => processArticles(articles))
        .catch(err => {
            console.log("GetArticles err:", err);
        })
    ;*/

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
