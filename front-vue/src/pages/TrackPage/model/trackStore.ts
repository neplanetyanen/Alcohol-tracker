import { type _GettersTree, defineStore } from 'pinia';

import { useUserStore } from '@/entities/User';

import { api } from '@/shared/api';
import type { UserStatistic } from '@/entities/User/model/types/user';

export const trackNamespace = 'track';

export interface TrackSchema {
	id: number;
}

export interface _TrackGetterSchema extends _GettersTree<TrackSchema> {}

export interface TrackActionsSchema {
	getTrack: () => Promise<void>;
}

export const useTrackStore = defineStore<
	string,
	TrackSchema,
	_TrackGetterSchema,
	TrackActionsSchema
>(trackNamespace, {
	state: (): TrackSchema => ({
		id: 0,
	}),
	actions: {
		async getTrack() {
			const responseTrack = await api.get(`/api/analytics/tracker/${this.id}`);
			const responseStatistic = await api.get(`/api/userStatistics/${this.id}`);
			if (responseTrack.data && responseStatistic.data) {
				const userStore = useUserStore();
				const userStatistic: UserStatistic = {
					tracker_score: responseTrack.data.tracker_score,
					last_consumption_date: responseStatistic.data.last_consumption_date,
					more_drinks: responseStatistic.data.more_drinks,
				};
				userStore.userStatistic(userStatistic);
			}
		},
	},
});
