import { type _GettersTree, defineStore } from 'pinia';

import { useCheckAuthApi, useLogoutApi } from '../api/requests';

import {
	type User,
	type UserInformation,
	type UserStatistic,
} from './types/user';

import {
	USER_LOCAL_STORAGE_KEY,
	TOKEN_LOCAL_STORAGE_KEY,
	USER_STATISTIC_LOCAL_STORAGE_KEY,
} from '@/shared/consts/localStorage';
import { router } from '@/app/providers/router';

export const userNamespace = 'user';

export interface UserSchema {
	authData?: User;
}

export interface _UserGetterSchema extends _GettersTree<UserSchema> {}

export interface UserActionsSchema {
	setUser: (params: User) => void;
	updateUser: (params: UserInformation) => void;
	logout: () => Promise<void>;
	checkAuth: () => Promise<void>;
	userStatistic: (params: UserStatistic) => void;
}

export const useUserStore = defineStore<
	string,
	UserSchema,
	_UserGetterSchema,
	UserActionsSchema
>(userNamespace, {
	state: (): UserSchema => ({
		authData: undefined,
	}),
	actions: {
		setUser(user: User) {
			this.authData = user;
			localStorage.setItem(USER_LOCAL_STORAGE_KEY, JSON.stringify(user.user));
			localStorage.setItem(TOKEN_LOCAL_STORAGE_KEY, user.token);
		},
		updateUser(user: UserInformation) {
			localStorage.setItem(USER_LOCAL_STORAGE_KEY, JSON.stringify(user));
		},
		async logout() {
			this.authData = undefined;
			localStorage.removeItem(USER_LOCAL_STORAGE_KEY);
			localStorage.removeItem(TOKEN_LOCAL_STORAGE_KEY);
			localStorage.removeItem(USER_STATISTIC_LOCAL_STORAGE_KEY);
			router.push('/');
		},
		async checkAuth() {
			const checkAuthApi = useCheckAuthApi();
			const response = await checkAuthApi.initiate();
			if (response) {
				this.authData = response;
				return;
			}
			this.authData = undefined;
			localStorage.removeItem(USER_LOCAL_STORAGE_KEY);
		},
		userStatistic(userStatistic: UserStatistic) {
			localStorage.setItem(
				USER_STATISTIC_LOCAL_STORAGE_KEY,
				JSON.stringify(userStatistic)
			);
		},
	},
});
