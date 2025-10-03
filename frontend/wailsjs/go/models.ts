export namespace main {
	
	export class Article {
	    title: string;
	    content: string;
	    sentimentGroup: string;
	    sentimentScore: number;
	    url: string;
	    alreadyRead: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Article(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.content = source["content"];
	        this.sentimentGroup = source["sentimentGroup"];
	        this.sentimentScore = source["sentimentScore"];
	        this.url = source["url"];
	        this.alreadyRead = source["alreadyRead"];
	    }
	}

}

