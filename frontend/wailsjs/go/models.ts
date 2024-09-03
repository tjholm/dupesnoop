export namespace analyzer {
	
	export class DiskInfo {
	    device: string;
	    mount: string;
	    fs: string;
	    used: number;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new DiskInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.device = source["device"];
	        this.mount = source["mount"];
	        this.fs = source["fs"];
	        this.used = source["used"];
	        this.total = source["total"];
	    }
	}

}

export namespace main {
	
	export class FileDetailsResult {
	    name: string;
	    size: number;
	    path: string;
	    hash: string;
	
	    static createFrom(source: any = {}) {
	        return new FileDetailsResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.size = source["size"];
	        this.path = source["path"];
	        this.hash = source["hash"];
	    }
	}

}

