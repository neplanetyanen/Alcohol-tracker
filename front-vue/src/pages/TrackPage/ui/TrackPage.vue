<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Button, Column, Row, Typography } from '@/shared/ui';
import { Card } from '@/widgets/Card';
import { useTrackStore } from '../model/trackStore';
import { computed, onMounted, ref, type Ref } from 'vue';
import type { UserInformation } from '@/entities/User';
import {
	USER_LOCAL_STORAGE_KEY,
	USER_STATISTIC_LOCAL_STORAGE_KEY,
} from '@/shared/consts/localStorage';
import type { UserStatistic } from '@/entities/User/model/types/user';

const trackStore = useTrackStore();

const currentUser: Ref<UserInformation | null> = ref(null);
const userStatistic: Ref<UserStatistic | null> = ref(null);

const record = computed(() => {
	if (userStatistic.value) {
		const lastDrinkDate: Date = new Date(userStatistic.value.tracker_score);
		const currentDate: Date = new Date();
		const differenceInMilliseconds: number =
			currentDate.getTime() - lastDrinkDate.getTime();
		const differenceInDays: number =
			differenceInMilliseconds / (1000 * 60 * 60 * 24);
		const roundedDifferenceInDays: number = Math.ceil(differenceInDays);

		if (roundedDifferenceInDays < 0) {
			return 0;
		} else {
			return Math.round(roundedDifferenceInDays);
		}
	} else {
		return 0;
	}
});
const canDrive = computed(() => {
	if (userStatistic.value) {
		const trackerScoreDate: Date = new Date(userStatistic.value.tracker_score);
		const currentDate: Date = new Date();
		const differenceInMilliseconds: number =
			trackerScoreDate.getTime() - currentDate.getTime();
		const differenceInHours: number =
			differenceInMilliseconds / (1000 * 60 * 60);

		if (differenceInHours < 0) {
			return 0;
		} else {
			return Math.round(differenceInHours);
		}
	} else {
		return 0;
	}
});

onMounted(async () => {
	const user = localStorage.getItem(USER_LOCAL_STORAGE_KEY);
	if (user) {
		currentUser.value = JSON.parse(user);
		if (currentUser.value) {
			trackStore.id = currentUser.value.id;
		}
	}
	await trackStore.getTrack();
	const statistic = localStorage.getItem(USER_STATISTIC_LOCAL_STORAGE_KEY);
	if (statistic) {
		userStatistic.value = JSON.parse(statistic);
	}
});
</script>
<template>
	<Page :header="'auth'" :background="'content-bckgr'">
		<Row :justify="'center'" class="track-container">
			<Column :justify="'center'" :align="'center'" class="left-side">
				<Typography :align="'center'" :size="'l'">
					Продолжайте в том же духе!<br />Не будьте эгоистами, в первую очередь
					думайте о себе! - Джейсон Стетхэм (с)
				</Typography>
			</Column>
			<Column
				:gap="'16'"
				:align="'center'"
				:justify="'center'"
				class="right-side"
			>
				<Card
					:title="'Вы не пьёте уже'"
					:count="String(record)"
					:unit="'дней'"
				></Card>
				<Card
					:title="'За руль можно сесть через'"
					:count="String(canDrive)"
					:unit="'часа(ов)'"
					:size-title="'s'"
				></Card>
				<Button class="btn-track" @click="$router.push('/track/input')"
					><Typography :size="'s'">Внести новые данные</Typography></Button
				>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./TrackPage.css');
</style>
