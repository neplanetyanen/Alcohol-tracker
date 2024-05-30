<script setup lang="ts">
import { Column, Row, Typography, Input, Button, Select } from '@/shared/ui';
import { userContent } from '../../constans/userContent';
import {
	computed,
	onMounted,
	ref,
	type ComputedRef,
	type PropType,
	type Ref,
} from 'vue';
import { signupContent } from '../../constans/authContent';
import { useAuthFormStore } from '../../model/authFormStore';
import {
	validationAuthFields,
	errorsFields,
} from '../../lib/validators/validationHandler';
import { USER_LOCAL_STORAGE_KEY } from '@/shared/consts/localStorage';
import { useUserStore, type UserInformation } from '@/entities/User';

const currentUser: Ref<UserInformation | null> = ref(null);

const authFormStore = useAuthFormStore();

onMounted(async () => {
	const user = localStorage.getItem(USER_LOCAL_STORAGE_KEY);
	if (user) {
		currentUser.value = JSON.parse(user);
		if (currentUser.value) {
			authFormStore.id = currentUser.value?.id;
			authFormStore.username = currentUser.value?.username;
			authFormStore.email = currentUser.value?.email;
			authFormStore.nationality = currentUser.value?.nationality;
			authFormStore.height = currentUser.value?.height;
			authFormStore.weight = currentUser.value?.weight;
			authFormStore.body_type = currentUser.value?.body_type;
			authFormStore.allergies = new Set<string>(currentUser.value?.allergies);
			authFormStore.goal = currentUser.value?.goal;
		}
	}
});

const isFormValid = computed(() => {
	return (
		authFormStore.username !== '' &&
		authFormStore.email !== '' &&
		authFormStore.height !== '' &&
		authFormStore.weight !== ''
	);
});

const updateBodyType = (option: string) => {
	const selectedOption = Object.keys(userContent.body_type).find(
		(key) => userContent.body_type[key] === option
	);
	if (selectedOption) {
		authFormStore.body_type = selectedOption;
	}
};

const updateAllergies = (option: Set<string>) => {
	const selectedOption = Array.from(option).map((item) =>
		Object.keys(userContent.allergies).find(
			(key) => userContent.allergies[key] === item
		)
	);
	const definedOptions = selectedOption.filter(
		(option): option is string => option !== undefined
	);

	authFormStore.allergies = new Set<string>(definedOptions);
};

const updateGoal = (option: string) => {
	const selectedOption = Object.keys(userContent.goal).find(
		(key) => userContent.goal[key] === option
	);
	if (selectedOption) {
		authFormStore.goal = selectedOption;
	}
};

const updateNationality = (option: string) => {
	const selectedOption = Object.keys(userContent.nationality).find(
		(key) => userContent.nationality[key] === option
	);
	if (selectedOption) {
		authFormStore.nationality = selectedOption;
	}
};
</script>
<template>
	<Column :gap="'16'" class="information" :align="'center'">
		<Typography :size="'m'">Мой профиль</Typography>
		<Column :gap="'8'" full-width :align="'start'">
			<Row full-width :justify="'between'">
				<Typography :size="'s'">Имя</Typography>
				<Input
					class="information-input"
					v-model="authFormStore.username"
					@input="validationAuthFields('username', authFormStore.username)"
					:borderError="errorsFields.username.value.size !== 0"
				/>
			</Row>
			<template
				v-if="errorsFields.username"
				v-for="error in errorsFields.username.value"
			>
				<Typography
					class="error-msg-update"
					:size="'xs'"
					:color="'error-msg'"
					>{{ error }}</Typography
				>
			</template>
			<Row full-width :justify="'between'">
				<Typography :size="'s'">Почта</Typography>
				<Input
					class="information-input"
					v-model="authFormStore.email"
					@input="validationAuthFields('email', authFormStore.email)"
					:borderError="errorsFields.email.value.size !== 0"
				/>
			</Row>
			<template
				v-if="errorsFields.email"
				v-for="error in errorsFields.email.value"
			>
				<Typography
					class="error-msg-update"
					:size="'xs'"
					:color="'error-msg'"
					>{{ error }}</Typography
				>
			</template>
			<Row v-if="currentUser?.nationality" full-width :justify="'between'">
				<Typography :size="'s'">Национальность</Typography>
				<Select
					:options="
						signupContent.content[1].btns?.map((item) => item.name) || []
					"
					:selectedOption="[userContent.nationality[currentUser.nationality]]"
					class="information-select-nationality"
					@updateSelectedOption="updateNationality"
				/>
			</Row>
			<Row full-width :justify="'between'">
				<Typography :size="'s'">Рост(см)</Typography>
				<Input
					:full-width="false"
					class="numerical-input"
					v-model="authFormStore.height"
					@input="validationAuthFields('height', authFormStore.height)"
					:borderError="errorsFields.height.value.size !== 0"
				/>
				<Typography :size="'s'">Вес(кг)</Typography>
				<Input
					:full-width="false"
					class="numerical-input"
					v-model="authFormStore.weight"
					@input="validationAuthFields('weight', authFormStore.weight)"
					:borderError="errorsFields.weight.value.size !== 0"
				/>
			</Row>

			<template
				v-if="errorsFields.height"
				v-for="error in errorsFields.height.value"
			>
				<Typography :size="'xs'" :color="'error-msg'">{{ error }}</Typography>
			</template>
			<template
				v-if="errorsFields.weight"
				v-for="error in errorsFields.weight.value"
			>
				<Typography :size="'xs'" :color="'error-msg'">{{ error }}</Typography>
			</template>
			<Row v-if="currentUser?.body_type" full-width :justify="'between'">
				<Typography :size="'s'">Телосложение</Typography>
				<Select
					class="information-select-body-type"
					:selectedOption="[userContent.body_type[currentUser?.body_type]]"
					:options="
						signupContent.content[2].btns?.map((item) => item.name) || []
					"
					@updateSelectedOption="updateBodyType"
				/>
			</Row>
			<Row full-width :justify="'between'" v-if="currentUser?.allergies">
				<Typography :size="'s'">Аллергии</Typography>
				<Select
					class="information-select-allergies"
					:selectedOption="
						currentUser?.allergies.map((item) => userContent.allergies[item])
					"
					:options="
						signupContent.content[3].btns?.map((item) => item.name) || []
					"
					multipleChoice
					@updateSelectedOption="updateAllergies"
				/>
			</Row>
			<Row v-if="currentUser?.goal" full-width :justify="'between'">
				<Typography :size="'s'">Цель</Typography>
				<Select
					class="information-select-goal"
					:options="
						signupContent.content[4].btns?.map((item) => item.name) || []
					"
					:selectedOption="[userContent.goal[currentUser?.goal]]"
					@updateSelectedOption="updateGoal"
				/>
			</Row>
		</Column>
		<Button
			@click="authFormStore.update"
			:variant-color="'purple-btn'"
			full-width
			:disabled="
				!isFormValid ||
				Object.values(errorsFields).some(
					(errorsSet) => errorsSet.value.size > 0
				)
			"
			><Typography :color="'light'" :size="'s'"
				>Сохранить изменения</Typography
			></Button
		>
	</Column>
</template>
<style scoped lang="css">
@import url('./UpdateForm.css');
</style>
