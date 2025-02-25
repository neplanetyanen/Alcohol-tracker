import { type User } from '../model/types/user';

import { buildApi } from '@/shared/api/lib/useApi';

export const checkAuthNamespace = 'check';
export const logoutNamespace = 'logout';

export const useCheckAuthApi = buildApi<User, void>(checkAuthNamespace, {
	url: '/auth/refresh',
	method: 'POST',
	withCredentials: true,
});

export const useLogoutApi = buildApi<void>(logoutNamespace, {
	url: '/auth/logout',
	method: 'POST',
	withCredentials: true,
});
