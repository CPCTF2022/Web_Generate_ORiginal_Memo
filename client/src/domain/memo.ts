export class Memo {
	id: string;
	userName: string;
	content: string;
	createdAt: Date;
	constructor(id: string, userName: string, content: string);
	constructor(id: string, userName: string, content: string, createdAt: string);
	constructor(id: string, userName: string, content: string, createdAt?: string) {
		this.id = id;
		this.userName = userName;
		this.content = content;
		this.createdAt = new Date(createdAt);
	}
}
