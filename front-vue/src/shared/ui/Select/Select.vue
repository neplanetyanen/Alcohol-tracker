<script lang="ts" setup>
import { ref, watch, type PropType } from 'vue';
import Polygon from '@/shared/assets/figures/Polygon.vue';

const props = defineProps({
	options: {
		type: Array<string>,
		required: true,
	},
	placeholder: {
		type: String,
		default: '',
	},
	fullWidth: {
		type: Boolean,
		default: false,
	},
	selectedOption: {
		type: Array as PropType<string[]>,
		default: [''],
	},
	multipleChoice: {
		type: Boolean,
		default: false,
	},
});

const emit = defineEmits();

const isOpen = ref(false);
const selectedMultiplyOptions = ref(new Set<string>(props.selectedOption));
const selectedOption = ref(props.selectedOption?.join(', '));

const selectOption = (option: string) => {
	selectedOption.value = option;
	isOpen.value = true;
	emit('updateSelectedOption', option);
};

const selectMultipleOptions = (option: string) => {
	if (
		selectedMultiplyOptions.value.has(option) &&
		selectedMultiplyOptions.value.size > 1
	) {
		selectedMultiplyOptions.value.delete(option);
	} else if (
		!selectedMultiplyOptions.value.has(option) &&
		option !== 'Ничего из перечисленного' &&
		!selectedMultiplyOptions.value.has('Ничего из перечисленного')
	) {
		selectedMultiplyOptions.value.add(option);
	} else if (option === 'Ничего из перечисленного') {
		selectedMultiplyOptions.value.clear();
		selectedMultiplyOptions.value.add(option);
	} else if (
		option !== 'Ничего из перечисленного' &&
		selectedMultiplyOptions.value.has('Ничего из перечисленного')
	) {
		selectedMultiplyOptions.value.delete('Ничего из перечисленного');
		selectedMultiplyOptions.value.add(option);
	}
	selectedOption.value = Array.from(selectedMultiplyOptions.value).join(', ');
	emit('updateSelectedOption', selectedMultiplyOptions.value);
};
</script>
<template>
	<div
		:class="{ 'full-width': props.fullWidth }"
		class="custom-select"
		@click="isOpen = !isOpen"
	>
		<div class="default-option">
			<div class="selected-option-container">
				{{ selectedOption || props.placeholder }}
			</div>

			<div class="polygon-container">
				<Polygon class="select-polygon" />
			</div>
		</div>
		<div v-if="isOpen" class="dropdown">
			<div
				v-for="option in props.options"
				:key="option"
				class="dropdown-option"
				@click="
					multipleChoice ? selectMultipleOptions(option) : selectOption(option)
				"
			>
				{{ option }}
			</div>
		</div>
	</div>
</template>
<style lang="css" scoped>
@import url('./Select.css');
</style>
