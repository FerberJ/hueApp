export namespace models {
	
	export class Room {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updateAt: any;
	    // Go type: gorm
	    deletedAt: any;
	    liked: boolean;
	    name: string;
	    lights: Light[];
	    on: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Room(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updateAt = this.convertValues(source["updateAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.liked = source["liked"];
	        this.name = source["name"];
	        this.lights = this.convertValues(source["lights"], Light);
	        this.on = source["on"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Light {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updateAt: any;
	    // Go type: gorm
	    deletedAt: any;
	    liked: boolean;
	    customName: string;
	    name: string;
	    on: boolean;
	    groupRef: string;
	    group: Group;
	    roomRef: string;
	    room: Room;
	    brightness: number;
	
	    static createFrom(source: any = {}) {
	        return new Light(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updateAt = this.convertValues(source["updateAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.liked = source["liked"];
	        this.customName = source["customName"];
	        this.name = source["name"];
	        this.on = source["on"];
	        this.groupRef = source["groupRef"];
	        this.group = this.convertValues(source["group"], Group);
	        this.roomRef = source["roomRef"];
	        this.room = this.convertValues(source["room"], Room);
	        this.brightness = source["brightness"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Group {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updateAt: any;
	    // Go type: gorm
	    deletedAt: any;
	    liked: boolean;
	    name: string;
	    lights: Light[];
	    on: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Group(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updateAt = this.convertValues(source["updateAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.liked = source["liked"];
	        this.name = source["name"];
	        this.lights = this.convertValues(source["lights"], Light);
	        this.on = source["on"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GroupOptions {
	    label: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new GroupOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.value = source["value"];
	    }
	}
	
	
	export class RoomOptions {
	    label: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new RoomOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.label = source["label"];
	        this.value = source["value"];
	    }
	}

}

