<script setup lang="ts">
import { Button, Column, Typography } from '@/shared/ui';
import { Page } from '@/widgets/Page';
</script>
<template>
	<Page :background="'home-bckgr'">
		<Column :gap="'42'" class="home-container">
			<Typography :size="'l'">Сделайте свою жизнь <u>лучше</u></Typography>
			<Typography :size="'m'"
				>нормализуйте потребление алкоголя и избавьтесь от вредной
				привычки</Typography
			>
			<Button
				:variant="'round-btn'"
				isLeftArrow
				@click="$router.push('/signup')"
				><Typography :size="'s'">Выбрать цель</Typography></Button
			>
		</Column>
	</Page>
</template>
<style scoped lang="css">
@import url('./HomePage.css');
</style>
