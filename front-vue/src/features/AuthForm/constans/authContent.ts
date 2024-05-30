import { type AuthContentType } from './types/authContentType';

export const loginContent: AuthContentType = {
	authType: 'login',
	content: [
		{
			title: undefined,
			inputs: [
				{
					name: 'Имя пользователя',
					model: 'username',
				},
				{
					name: 'Пароль',
					model: 'password',
				},
			],
			btns: undefined,
			link: 'Нет аккаунта? <u>Регистрация</u>',
			btnText: 'Вход',
			factTitle: 'Интересный факт...',
			factSubtitle:
				'100 граммов водки убивают 7,5 тысяч клеток головного мозга',
		},
	],
};

export const signupContent: AuthContentType = {
	authType: 'signup',
	content: [
		{
			title: undefined,
			inputs: [
				{
					name: 'Имя',
					model: 'username',
				},
				{
					name: 'Электронный адрес',
					model: 'email',
				},
				{
					name: 'Пароль',
					model: 'password',
				},
			],
			btns: undefined,
			link: undefined,
			btnText: 'Далее',
			factTitle: 'Еще один интересный факт...',
			factSubtitle:
				'Алкоголь оказывает влияние на мозг уже через 6 минут после приема',
		},
		{
			title: 'К какой части света принадлежит Ваша национальность?',
			inputs: undefined,
			btns: [
				{
					name: 'Север',
					value: 'north',
					model: 'nationality',
				},
				{
					name: 'Северо-восток',
					value: 'northeast',
					model: 'nationality',
				},
				{
					name: 'Восток',
					value: 'east',
					model: 'nationality',
				},
				{
					name: 'Юго-восток',
					value: 'southeast',
					model: 'nationality',
				},
				{
					name: 'Юг',
					value: 'south',
					model: 'nationality',
				},
				{
					name: 'Юго-запад',
					value: 'southwest',
					model: 'nationality',
				},
				{
					name: 'Запад',
					value: 'west',
					model: 'nationality',
				},
				{
					name: 'Северо-запад',
					value: 'northwest',
					model: 'nationality',
				},
			],
			link: undefined,
			btnText: 'Далее',
			factTitle: 'Почему мы это спрашиваем?',
			factSubtitle:
				'Организм каждого человека по-разному реагирует на алкоголь, и эта особенность наследственная, она записана в генетическом коде.<br/>Учёные считают, что есть целые народы и группы населения, которые подвержены пристрастию к алкоголю или, наоборот, переносят его легче других',
		},
		{
			title: 'Тип телосложения',
			inputs: [
				{
					name: 'Рост(см)',
					model: 'height',
				},
				{
					name: 'Вес(кг)',
					model: 'weight',
				},
			],
			btns: [
				{
					name: 'Эктоморфный',
					value: 'ectomorphic',
					model: 'body_type',
				},
				{
					name: 'Мезоморфный',
					value: 'mesomorphic',
					model: 'body_type',
				},
				{
					name: 'Эндоморфный',
					value: 'endomorphic',
					model: 'body_type',
				},
			],
			link: undefined,
			btnText: 'Далее',
			factTitle: 'Как определить свой тип телосложения',
			factSubtitle:
				'Измерьте обхват запястья при помощи сантиметровой ленты и соотнесите результаты со значениями из таблицы',
		},
		{
			title: 'Есть ли у Вас аллергии?',
			inputs: undefined,
			btns: [
				{
					name: 'Дрожжи',
					value: 'yeast',
					model: 'allergies',
				},
				{
					name: 'Хмель',
					value: 'hops',
					model: 'allergies',
				},
				{
					name: 'Пшеничный спирт',
					value: 'wheatAlcohol',
					model: 'allergies',
				},
				{
					name: 'Орехи',
					value: 'nuts',
					model: 'allergies',
				},
				{
					name: 'Фрукты и ягоды',
					value: 'fruitsBberries',
					model: 'allergies',
				},
				{
					name: 'Пищевые загустители',
					value: 'foodThickeners',
					model: 'allergies',
				},
				{
					name: 'Полисахариды',
					value: 'polysaccharides',
					model: 'allergies',
				},
				{
					name: 'Ничего из перечисленного',
					value: 'none',
					model: 'allergies',
				},
			],
			link: undefined,
			btnText: 'Далее',
			factTitle: 'Важно!',
			factSubtitle:
				'Мы не несём ответственность за последствия от употребления алкоголя. Если у Вас есть аллергии или хронические заболевания, ограничивающие/запрещающие употребления алкоголя, то воздержитесь от распития данных напитков<br/><i>Берегите здоровье!</i>',
		},
		{
			title: 'Какая цель Вам подходит больше всего?',
			inputs: undefined,
			btns: [
				{
					name: 'Следить за количеством употребляемого алкоголя',
					value: 'monitor',
					model: 'goal',
				},
				{
					name: 'Ради интереса',
					value: 'fun',
					model: 'goal',
				},
				{
					name: 'Избавится от вредной привычки',
					value: 'getRid',
					model: 'goal',
				},
				{
					name: 'Хочу знать больше при помощи приложения',
					value: 'wantKnow',
					model: 'goal',
				},
			],
			link: undefined,
			btnText: 'Далее',
			factTitle: 'И это тоже важно!',
			factSubtitle:
				'Наше приложения не предназначено для лечения алкоголизма. Если у Вас серьезные проблемы с употреблением алкоголя, обратитесь к специалисту.<br/>Заранее спасибо!<br/><i>Ваша команда Drunk-O-Meter</i>',
		},
	],
};
