import { type _GettersTree, defineStore } from 'pinia';

import { api } from '@/shared/api';
export const countNormNamespace = 'countNorm';

export interface CountNormSchema {
	id: number;
	alcoholPercentage: string;
	currentNorm: string;
}

export interface _CountNormGetterSchema extends _GettersTree<CountNormSchema> {}

export interface CountNormActionsSchema {
	countNorm: () => Promise<void>;
}

export const useCountNormStore = defineStore<
	string,
	CountNormSchema,
	_CountNormGetterSchema,
	CountNormActionsSchema
>(countNormNamespace, {
	state: (): CountNormSchema => ({
		id: 0,
		alcoholPercentage: '',
		currentNorm: '',
	}),
	actions: {
		async countNorm() {
			const response = await api.post(
				`/api/analytics/norm/${this.id}`,
				{
					alcohol_percentage: Number(this.alcoholPercentage.replace(/,/g, '.')),
				},
				{ withCredentials: true }
			);
			if (response.data) {
				const result: string = String(response.data.result);
				this.currentNorm = result.slice(0, 7);
				return;
			}
		},
	},
});
