<script setup lang="ts">
import { Column } from '@/shared/ui';
import Typography from '@/shared/ui/Typography/Typography.vue';
import Line from '@/shared/assets/figures/Line.vue';
import { type TypographySize } from '@/shared/ui/Typography/Typography.vue';
import type { PropType } from 'vue';

const props = defineProps({
	title: {
		type: String,
		required: true,
	},
	count: {
		type: String,
		required: true,
	},
	unit: {
		type: String,
		default: '',
	},
	sizeTitle: {
		type: String as PropType<TypographySize>,
		default: 'm',
	},
	sizeCount: {
		type: String as PropType<TypographySize>,
		default: 'xl',
	},
	underline: {
		type: Boolean,
		default: true,
	},
});
</script>
<template>
	<Column :gap="'16'" :align="'center'" class="card">
		<Typography :size="sizeTitle">{{ props.title }}</Typography>
		<Column :gap="'4'" :align="'center'">
			<Typography
				:class="{ 'count-card': props.underline }"
				:size="sizeCount"
				>{{ props.count }}</Typography
			>
			<Typography :size="sizeTitle">{{ props.unit }}</Typography>
		</Column>
	</Column>
</template>
<style scoped lang="css">
@import url('./Card.css');
</style>
