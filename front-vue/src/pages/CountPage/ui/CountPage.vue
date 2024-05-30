<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Card } from '@/widgets/Card';
import { Column, Row, Typography, Button, Select, Input } from '@/shared/ui';
import { useCountNormStore } from '../model/countNormStore';
import { USER_LOCAL_STORAGE_KEY } from '@/shared/consts/localStorage';
import { computed, onMounted, ref, type Ref } from 'vue';
import type { UserInformation } from '@/entities/User';

const countNormStore = useCountNormStore();
const currentUser: Ref<UserInformation | null> = ref(null);
onMounted(() => {
	const user = localStorage.getItem(USER_LOCAL_STORAGE_KEY);
	if (user) {
		currentUser.value = JSON.parse(user);
		if (currentUser.value) {
			countNormStore.id = currentUser.value.id;
		}
	}
});

const isAlcoholPercentageValid = computed(() => {
	const regexp = /^\d{0,2}[.,]?\d{1,2}$/;
	if (
		(!regexp.test(countNormStore.alcoholPercentage) ||
			Number(countNormStore.alcoholPercentage) > 100) &&
		countNormStore.alcoholPercentage != ''
	) {
		return false;
	} else {
		return true;
	}
});
</script>
<template>
	<Page :header="'auth'" :background="'content-bckgr'">
		<Row :justify="'center'" class="count-container">
			<Column :justify="'center'" :align="'center'" class="left-side">
				<Typography :align="'center'" :size="'m'">
					<strong>Внимание!</strong><br />Данные могут быть неточными, так как
					мы не знаем, пили ли Вы на голодный желудок или нет, какое у Вас
					настроение, фаза луны и так далее Принимайте алкоголь с умом!
				</Typography>
			</Column>
			<Column
				:gap="'16'"
				:align="'center'"
				:justify="'center'"
				class="right-side"
			>
				<Column :gap="'16'" class="count" :align="'start'">
					<Column full-width :align="'start'" :gap="'8'">
						<Typography :size="'s'"
							>Введите процент алкоголя в напитке
						</Typography>
						<Input full-width v-model="countNormStore.alcoholPercentage" />
						<Column class="error-alcohol-percentage">
							<template v-if="!isAlcoholPercentageValid">
								<Typography :color="'error-msg'" :size="'xs'"
									>Неверный формат. Введите десятичное число с точностью до
									сотых, либо натуральное число. Максимальный процент - 100%
								</Typography>
							</template>
						</Column>
					</Column>
					<Card
						:title="'Ваша норма'"
						:count="countNormStore.currentNorm || '?'"
						:unit="'мл'"
					></Card>
					<Button
						v-if="currentUser"
						full-width
						@click="countNormStore.countNorm()"
						:disabled="
							!isAlcoholPercentageValid ||
							countNormStore.alcoholPercentage === ''
						"
						><Typography :size="'s'">Рассчитать норму</Typography></Button
					>
				</Column>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./CountPage.css');
</style>
