<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Column, Row, Typography } from '@/shared/ui';
import { Card } from '@/widgets/Card';
import { UpdateForm } from '@/features/AuthForm';
import {
	USER_LOCAL_STORAGE_KEY,
	USER_STATISTIC_LOCAL_STORAGE_KEY,
} from '@/shared/consts/localStorage';
import { type UserInformation } from '@/entities/User';
import { computed, onMounted, ref, type Ref } from 'vue';
import type { UserStatistic } from '@/entities/User/model/types/user';
import { useTrackStore } from '@/pages/TrackPage';

const trackStore = useTrackStore();

const currentUser: Ref<UserInformation | null> = ref(null);
const userStatistic: Ref<UserStatistic | null> = ref(null);

onMounted(() => {
	const user = localStorage.getItem(USER_LOCAL_STORAGE_KEY);
	if (user) {
		currentUser.value = JSON.parse(user);
	}
	if (currentUser.value) {
		trackStore.id = currentUser.value?.id;
	}
	trackStore.getTrack();
	const statistic = localStorage.getItem(USER_STATISTIC_LOCAL_STORAGE_KEY);
	if (statistic) {
		userStatistic.value = JSON.parse(statistic);
	}
});

const lastConsumptionDate = computed(() => {
	if (userStatistic.value) {
		const dateObject: Date = new Date(
			userStatistic.value?.last_consumption_date
		);
		const year: number = dateObject.getUTCFullYear();
		const month: number = dateObject.getUTCMonth() + 1;
		const day: number = dateObject.getUTCDate();

		const formattedDate: string = `${day.toString().padStart(2, '0')}-${month
			.toString()
			.padStart(2, '0')}-${year}`;
		return formattedDate;
	} else {
		return Date.now();
	}
});
</script>
<template>
	<Page :header="'auth'" :background="'content-bckgr'">
		<Row :justify="'center'" class="user-container">
			<Column :justify="'start'" :align="'center'" class="left-side">
				<UpdateForm />
			</Column>
			<Column
				:gap="'16'"
				:align="'center'"
				:justify="'start'"
				class="right-side"
			>
				<Column :gap="'16'" class="statistic" :align="'center'">
					<Typography :size="'m'">Моя статистика</Typography>
					<Card
						:title="'Последний раз я пил(а)'"
						:count="String(lastConsumptionDate)"
						:size-count="'m'"
						:size-title="'m'"
					/>
					<Card
						v-if="userStatistic"
						:title="'Чаще всего я пью'"
						:count="String(userStatistic?.more_drinks)"
						:size-count="'m'"
						:size-title="'m'"
					/>
				</Column>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./UserPage.css');
</style>
