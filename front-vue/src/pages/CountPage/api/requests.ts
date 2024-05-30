import { buildApi } from '@/shared/api/lib/useApi';

export const countNormNamespace = 'count';

export interface Norm {
	norm: string;
}

export interface CountNormArgs {
	id: number;
	option: string;
}

export const useCountNormApi = buildApi<Norm, CountNormArgs>(
	countNormNamespace,
	{
		url: '/count',
		method: 'POST',
		withCredentials: true,
	}
);
