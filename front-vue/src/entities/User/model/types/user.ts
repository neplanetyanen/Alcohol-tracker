export interface User {
	token: string;
	user: UserInformation;
}

export interface UserInformation {
	id: number;
	username: string;
	email: string;
	nationality: string;
	height: string;
	weight: string;
	body_type: string;
	allergies: string[];
	goal: string;
}

export interface UserStatistic {
	tracker_score: Date;
	last_consumption_date: Date;
	more_drinks: string;
}
