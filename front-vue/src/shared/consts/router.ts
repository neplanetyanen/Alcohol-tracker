export const Routes = {
	HOME: 'home',
	LOGIN: 'login',
	SIGNUP: 'signup',
	TRACK: 'track',
	TRACKINPUT: 'trackInput',
	COUNT: 'count',
	USER: 'user',
} as const;

export type Route = ValueOf<typeof Routes>;

export const getRouteHome = () => '/';
export const getRouteLogIn = () => '/login';
export const getRouteSignUp = () => '/signup';
export const getRouteTrackPage = () => '/track';
export const getRouteTrackInputPage = () => '/track/input';
export const getRouteCountPage = () => '/count';
export const getRouteUserPage = () => '/user';
