export type AuthContentType = {
	authType: 'signup' | 'login';
	content: {
		title: undefined | string;
		inputs: undefined | { name: string; model: string }[];
		btns: undefined | { name: string; value: string; model: string }[];
		link: undefined | string;
		btnText: string;
		factTitle: string;
		factSubtitle: string;
	}[];
};
