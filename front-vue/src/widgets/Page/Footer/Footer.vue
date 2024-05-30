<script lang="ts" setup>
import { Column, Typography } from '@/shared/ui';
</script>
<template>
	<Column class="footer">
		<Typography :size="'xs'" :weight="500">
			Выполнено по программе предмета "Интернет технологии"
		</Typography>
		<Typography :size="'xs'" :weight="500"> © 2024, "Пряники" </Typography>
	</Column>
</template>
<style lang="css" scoped>
@import url('./Footer.css');
</style>
