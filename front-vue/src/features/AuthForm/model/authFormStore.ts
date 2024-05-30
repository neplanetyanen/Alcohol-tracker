import { type _GettersTree, defineStore } from 'pinia';

import { useLoginApi, useSignupApi } from '../api/requests';

import { useUserStore, type User, type UserInformation } from '@/entities/User';
import {
	validationAuth,
	validationAuthErrorClear,
} from '../lib/validators/validationHandler';
import { api } from '@/shared/api';
import { router } from '@/app/providers/router';

export const authFormNamespace = 'authForm';

export interface AuthFormSchema {
	id: number;
	username: string;
	email: string;
	password: string;
	nationality: string;
	height: string;
	weight: string;
	body_type: string;
	allergies: Set<string>;
	goal: string;
	error?: string;
}

export interface _AuthFormGetterSchema extends _GettersTree<AuthFormSchema> {}

export interface AuthFormActionsSchema {
	signup: () => Promise<void>;
	login: () => Promise<void>;
	update: () => Promise<void>;
	resetForm: () => void;
}

export const useAuthFormStore = defineStore<
	string,
	AuthFormSchema,
	_AuthFormGetterSchema,
	AuthFormActionsSchema
>(authFormNamespace, {
	state: (): AuthFormSchema => ({
		id: 0,
		username: '',
		email: '',
		password: '',
		nationality: '',
		height: '',
		weight: '',
		body_type: '',
		allergies: new Set<string>(),
		goal: '',
		error: undefined,
	}),
	actions: {
		async signup() {
			validationAuthErrorClear('all');

			this.error = undefined;
			const signupApi = useSignupApi();
			const response = await signupApi.initiate({
				username: this.username,
				email: this.email,
				password: this.password,
				nationality: this.nationality,
				height: Number(this.height),
				weight: Number(this.weight),
				body_type: this.body_type,
				allergies: Array.from(this.allergies),
				goal: this.goal,
			});
			this.resetForm();
			if (response) {
				const user: User = {
					user: {
						id: response.user.id,
						username: response.user.username,
						email: response.user.email,
						nationality: response.user.nationality,
						height: String(response.user.height),
						weight: String(response.user.weight),
						body_type: response.user.body_type,
						allergies: response.user.allergies,
						goal: response.user.goal,
					},
					token: response.token,
				};
				const userStore = useUserStore();
				userStore.setUser(user);
				router.push('/track');
				return;
			}
			this.error = signupApi.error;
		},
		async login() {
			validationAuthErrorClear('all');

			this.error = undefined;
			const loginApi = useLoginApi();
			const response = await loginApi.initiate({
				username: this.username,
				password: this.password,
			});
			this.resetForm();
			if (response) {
				const user: User = {
					user: {
						id: response.user.id,
						username: response.user.username,
						email: response.user.email,
						nationality: response.user.nationality,
						height: String(response.user.height),
						weight: String(response.user.weight),
						body_type: response.user.body_type,
						allergies: response.user.allergies,
						goal: response.user.goal,
					},
					token: response.token,
				};
				const userStore = useUserStore();
				userStore.setUser(user);

				router.push('/track');
				return;
			}
			validationAuth(loginApi.error);
			this.error = loginApi.error;
		},
		async update() {
			this.error = undefined;
			const response = await api.put(
				`/api/users/${this.id}`,
				{
					username: this.username,
					email: this.email,
					nationality: this.nationality,
					height: Number(this.height),
					weight: Number(this.weight),
					body_type: this.body_type,
					allergies: Array.from(this.allergies),
					goal: this.goal,
					id: this.id,
				},
				{ withCredentials: true }
			);
			const user: UserInformation = response.data;
			if (user) {
				const userStore = useUserStore();
				userStore.updateUser(user);
				return;
			}
		},
		resetForm() {
			this.username = '';
			this.email = '';
			this.password = '';
			this.nationality = '';
			this.height = '';
			this.weight = '';
			this.body_type = '';
			this.allergies = new Set<string>();
			this.goal = '';
			this.error = undefined;
		},
	},
});
