export class User {
	id: string;
	name: string;
	password: string;
	created_at: Date;
	constructor(id: string, name: string, password: string, created_at: string) {
		this.id = id;
		this.name = name;
		this.password = password;
		this.created_at = new Date(created_at);
	}
}
