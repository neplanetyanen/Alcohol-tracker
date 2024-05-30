<script setup lang="ts">
import { Row, Button, Typography } from '@/shared/ui';
import Icon from '@/shared/assets/icons/Icon.vue';
import { type PropType } from 'vue';
import { useAuthFormStore } from '@/features/AuthForm/model/authFormStore';
import { useUserStore } from '@/entities/User';

const userStore = useUserStore();

const authFormStore = useAuthFormStore();

export type HeaderVariants = 'auth' | 'unauth' | 'none';

const props = defineProps({
	header: {
		type: String as PropType<HeaderVariants>,
		default: 'unauth',
	},
});
</script>
<template>
	<Row class="header" v-if="props.header !== 'none'">
		<Row class="common-header" :justify="'between'">
			<Row :gap="'16'">
				<Icon
					@click="
						props.header === 'auth' ? $router.push('/track') : $router.push('/')
					"
					class="icon"
				/>
				<Typography :size="'s'" :weight="500">Drunk-O-Meter</Typography>
			</Row>
			<Row v-if="props.header === 'unauth'" :gap="'16'">
				<Button
					:variant="'round-btn'"
					:variantColor="'green-btn'"
					@click="$router.push('/login')"
					><Typography :size="'s'">Вход</Typography></Button
				>
			</Row>
			<Row v-if="props.header === 'auth'" :gap="'42'" :justify="'end'">
				<Typography
					:size="'s'"
					class="user-btn"
					:class="{
						selected: $router.currentRoute.value.fullPath.includes('/track'),
					}"
					@click="$router.push('/track')"
					>Трекер</Typography
				>
				<Typography
					:size="'s'"
					class="user-btn"
					:class="{
						selected: $router.currentRoute.value.fullPath === '/count',
					}"
					@click="$router.push('/count')"
					>Счётчик нормы</Typography
				>
				<Typography
					:size="'s'"
					class="user-btn"
					:class="{
						selected: $router.currentRoute.value.fullPath === '/user',
					}"
					@click="$router.push('/user')"
					>Мой профиль</Typography
				>
				<Button
					@click="
						userStore.logout();
						authFormStore.resetForm();
					"
					:variant="'round-btn'"
					:variantColor="'green-btn'"
					><Typography :size="'s'">Выход</Typography></Button
				>
			</Row>
		</Row>
	</Row>
</template>
<style lang="css" scoped>
@import url('./Header.css');
</style>
