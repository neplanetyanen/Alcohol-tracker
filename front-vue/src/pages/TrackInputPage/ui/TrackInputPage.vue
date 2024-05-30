<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Row, Column, Button, Typography, Input } from '@/shared/ui';
import { useTrackInputStore } from '../model/trackInputStore';
import { computed } from 'vue';
import { USER_LOCAL_STORAGE_KEY } from '@/shared/consts/localStorage';
import { onMounted, ref, type Ref } from 'vue';
import type { UserInformation } from '@/entities/User';

const trackInputStore = useTrackInputStore();
const currentUser: Ref<UserInformation | null> = ref(null);

onMounted(async () => {
	const user = localStorage.getItem(USER_LOCAL_STORAGE_KEY);
	if (user) {
		currentUser.value = JSON.parse(user);
		if (currentUser.value) {
			trackInputStore.id = currentUser.value.id;
		}
	}
	// await trackStore.getTrack();
});

const isTrackInputValid = computed(() => {
	if (
		trackInputStore.drinkName !== '' &&
		trackInputStore.volume !== '' &&
		trackInputStore.alcoholPercentage !== '' &&
		trackInputStore.date !== undefined &&
		trackInputStore.time !== undefined
	) {
		return true;
	} else {
		return false;
	}
});
</script>
<template>
	<Page :header="'auth'" :background="'content-bckgr'">
		<Row :justify="'center'" class="track-input-container">
			<Column :justify="'center'" :align="'center'" class="left-side">
				<Typography :align="'center'" :size="'l'">
					Важно!<br />После внесения данных Ваш трекер обнулится!
				</Typography>
			</Column>
			<Column :align="'center'" :justify="'center'" class="right-side">
				<Column :gap="'8'" class="track-input" :align="'start'">
					<Column full-width :align="'start'" :gap="'0'">
						<Typography :size="'s'">Введите название напитка </Typography>
						<Input full-width v-model="trackInputStore.drinkName" />
					</Column>
					<Column full-width :align="'start'" :gap="'0'">
						<Typography :size="'s'">Введите объём (мл) </Typography>
						<Row full-width :justify="'between'">
							<Input full-width v-model="trackInputStore.volume" />
						</Row>
					</Column>
					<Column full-width :align="'start'" :gap="'0'">
						<Typography :size="'s'"
							>Введите процент алкоголя в напитке
						</Typography>
						<Row full-width :justify="'between'">
							<Input full-width v-model="trackInputStore.alcoholPercentage" />
						</Row>
					</Column>
					<Column full-width :align="'start'" :gap="'0'">
						<Typography full-width :size="'s'"
							>Введите дату употребления
						</Typography>
						<Input full-width :type="'date'" v-model="trackInputStore.date" />
					</Column>
					<Column full-width :align="'start'" :gap="'0'">
						<Typography full-width :size="'s'"
							>Введите время употребления
						</Typography>
						<Input full-width :type="'time'" v-model="trackInputStore.time" />
					</Column>
					<Button
						full-width
						class="btn-track"
						@click="trackInputStore.createNewTrack()"
						:disabled="!isTrackInputValid"
						><Typography :size="'s'">Сохранить</Typography></Button
					>
				</Column>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./TrackInputPage.css');
</style>
