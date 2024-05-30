export interface UserItem {
	[key: string]: string;
}

export interface UserContent {
	[key: string]: UserItem;
}

export const userContent: UserContent = {
	nationality: {
		north: 'Север',
		northeast: 'Северо-восток',
		east: 'Восток',
		southeast: 'Юго-восток',
		south: 'Юг',
		southwest: 'Юго-запад',
		west: 'Запад',
		northwest: 'Северо-запад',
	},
	body_type: {
		ectomorphic: 'Эктоморфный',
		mesomorphic: 'Мезоморфный',
		endomorphic: 'Эндоморфный',
	},
	allergies: {
		yeast: 'Дрожжи',
		hops: 'Хмель',
		wheatAlcohol: 'Пшеничный спирт',
		nuts: 'Орехи',
		fruitsBberries: 'Фрукты и ягоды',
		foodThickeners: 'Пищевые загустители',
		polysaccharides: 'Полисахариды',
		none: 'Ничего из перечисленного',
	},
	goal: {
		monitor: 'Следить за количеством употребляемого алкоголя',
		fun: 'Ради интереса',
		getRid: 'Избавится от вредной привычки',
		wantKnow: 'Хочу знать больше при помощи приложения',
	},
};
