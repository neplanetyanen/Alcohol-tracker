import { type _GettersTree, defineStore } from 'pinia';

import { api } from '@/shared/api';
import { router } from '@/app/providers/router';

export const trackInputNamespace = 'trackInput';

export interface TrackInputSchema {
	id: number;
	date?: Date;
	time?: Date;
	volume: string;
	drinkName: string;
	alcoholPercentage: string;
}

export interface _TrackInputGetterSchema
	extends _GettersTree<TrackInputSchema> {}

export interface TrackInputActionsSchema {
	createNewTrack: () => Promise<void>;
}

export const useTrackInputStore = defineStore<
	string,
	TrackInputSchema,
	_TrackInputGetterSchema,
	TrackInputActionsSchema
>(trackInputNamespace, {
	state: (): TrackInputSchema => ({
		id: 0,
		date: undefined,
		time: undefined,
		volume: '',
		drinkName: '',
		alcoholPercentage: '',
	}),
	actions: {
		async createNewTrack() {
			const response = await api.post(`/api/drinkRecords/${this.id}`, {
				date: this.date,
				time: this.time,
				quantity: Number(this.volume),
				drink_name: this.drinkName,
				alcohol_percentage: Number(this.alcoholPercentage),
			});

			router.push('/track');
		},
	},
});
