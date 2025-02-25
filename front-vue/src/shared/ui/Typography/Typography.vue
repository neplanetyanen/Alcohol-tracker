<script setup lang="ts">
import { type PropType, computed } from 'vue';

export type TypographySize = 'xs' | 's' | 'm' | 'l' | 'xl' | 'xxl';
export type TypographyAlign = 'auto' | 'start' | 'center' | 'end' | 'justify';
export type TypographyWeight = 100 | 300 | 500;
export type TypographyColor = 'dark' | 'pale' | 'light' | 'error-msg';
export type TypographyStyleText = 'italic' | 'default';

const props = defineProps({
	size: {
		type: String as PropType<TypographySize>,
		default: 's',
	},
	align: {
		type: String as PropType<TypographyAlign>,
		default: 'auto',
	},
	weight: {
		type: Number as PropType<TypographyWeight>,
		default: 500,
	},
	as: {
		type: String,
		default: 'p',
	},
	fullWidth: {
		type: Boolean,
	},
	uppercase: {
		type: Boolean,
	},
	color: {
		type: String as PropType<TypographyColor>,
		default: 'dark',
	},
	styleText: {
		type: String as PropType<TypographyStyleText>,
		default: 'default',
	},
});

const textSize: Record<TypographySize, string> = {
	xs: 'size-xs',
	s: 'size-s',
	m: 'size-m',
	l: 'size-l',
	xl: 'size-xl',
	xxl: 'size-xxl',
};

const textAlign: Record<TypographyAlign, string> = {
	auto: 'align-auto',
	start: 'align-start',
	center: 'align-center',
	end: 'align-end',
	justify: 'align-justify',
};

const textWeight: Record<TypographyWeight, string> = {
	100: 'weight-100',
	300: 'weight-300',
	500: 'weight-500',
};

const classes = computed(() => ({
	[textWeight[props.weight]]: true,
	[textAlign[props.align]]: true,
	[textSize[props.size]]: true,
	uppercase: props.uppercase,
	'full-width': props.fullWidth,
	[props.color]: true,
	[props.styleText]: true,
}));
</script>

<template>
	<as class="typography" :class="classes">
		<slot></slot>
	</as>
</template>

<style lang="css" scoped>
@import url('./Typography.css');
</style>
