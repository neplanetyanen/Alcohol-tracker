import { type User } from '@/entities/User';
import { buildApi } from '@/shared/api/lib/useApi';

export const loginNamespace = 'login';
export const signupNamespace = 'signup';

export interface LoginRequestArgs {
	username: string;
	password: string;
}

export const useLoginApi = buildApi<User, LoginRequestArgs>(loginNamespace, {
	url: '/auth/sign-in',
	method: 'POST',
	withCredentials: true,
});

export interface SignupRequestArgs {
	username: string;
	email: string;
	password: string;
	nationality: string;
	height: number;
	weight: number;
	body_type: string;
	allergies: string[];
	goal: string;
}

export const useSignupApi = buildApi<User, SignupRequestArgs>(signupNamespace, {
	url: '/auth/sign-up',
	method: 'POST',
	withCredentials: true,
});
