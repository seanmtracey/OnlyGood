export namespace main {
	
	export class Article {
	    Title: string;
	    Content: string;
	    Sentiment: number;
	    URL: string;
	
	    static createFrom(source: any = {}) {
	        return new Article(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Content = source["Content"];
	        this.Sentiment = source["Sentiment"];
	        this.URL = source["URL"];
	    }
	}

}

