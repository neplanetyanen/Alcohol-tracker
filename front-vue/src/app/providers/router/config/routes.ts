import type { RouteRecordRaw } from 'vue-router';

import { HomePage } from '@/pages/HomePage';
import { LogInPage } from '@/pages/LogInPage';
import { SignUpPage } from '@/pages/SignUpPage';
import { TrackPage } from '@/pages/TrackPage';
import { TrackInputPage } from '@/pages/TrackInputPage';
import { CountPage } from '@/pages/CountPage';
import { UserPage } from '@/pages/UserPage';
import {
	Routes,
	getRouteHome,
	getRouteLogIn,
	getRouteSignUp,
	getRouteTrackPage,
	getRouteTrackInputPage,
	getRouteCountPage,
	getRouteUserPage,
} from '@/shared/consts/router';

export const routes: RouteRecordRaw[] = [
	{
		path: getRouteHome(),
		name: Routes.HOME,
		component: HomePage,
	},
	{
		path: getRouteLogIn(),
		name: Routes.LOGIN,
		component: LogInPage,
	},
	{
		path: getRouteSignUp(),
		name: Routes.SIGNUP,
		component: SignUpPage,
	},
	{
		path: getRouteTrackPage(),
		name: Routes.TRACK,
		component: TrackPage,
	},
	{
		path: getRouteTrackInputPage(),
		name: Routes.TRACKINPUT,
		component: TrackInputPage,
	},
	{
		path: getRouteCountPage(),
		name: Routes.COUNT,
		component: CountPage,
	},
	{
		path: getRouteUserPage(),
		name: Routes.USER,
		component: UserPage,
	},
];
