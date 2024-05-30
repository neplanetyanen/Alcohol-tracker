<script setup lang="ts">
import { Column, Input, Button } from '@/shared/ui';
import { computed, onMounted, ref, type ComputedRef, type PropType } from 'vue';
import type { AuthContentType } from '../../constans/types/authContentType';
import { loginContent, signupContent } from '../../constans/authContent';
import Typography from '@/shared/ui/Typography/Typography.vue';
import { useAuthFormStore } from '../../model/authFormStore';
import {
	validationAuthFields,
	errorsFields,
	errorsServer,
	validationAuthErrorClear,
} from '../../lib/validators/validationHandler';

const props = defineProps({
	authType: {
		type: String as PropType<'signup' | 'login'>,
		default: loginContent,
	},
});

const authFormStore = useAuthFormStore();

const authContent: ComputedRef<AuthContentType> = computed(() =>
	props.authType === 'signup' ? signupContent : loginContent
);

const countPages = ref(0);
const valueBtnSelected = ref(new Set<string>());

const emits = defineEmits();

onMounted(() => {
	emits('authContentChanged', authContent.value);
});

const handlerForm = () => {
	if (props.authType === 'signup') {
		if (countPages.value < 4) {
			countPages.value += 1;
			emits('pageChanged', countPages.value);
			valueBtnSelected.value = new Set<string>();
		} else {
			authFormStore.signup();
			countPages.value = 0;
		}
	} else {
		authFormStore.login();
		countPages.value = 0;
	}
};

const selectedBtn = (value: string, model: string) => {
	if (model === 'nationality') {
		authFormStore.nationality = value;
		valueBtnSelected.value.clear();
		valueBtnSelected.value.add(value);
	} else if (model === 'body_type') {
		authFormStore.body_type = value;
		valueBtnSelected.value.clear();
		valueBtnSelected.value.add(value);
	} else if (model === 'allergies') {
		if (valueBtnSelected.value.has(value)) {
			valueBtnSelected.value.delete(value);
		} else if (
			!valueBtnSelected.value.has(value) &&
			value !== 'none' &&
			!valueBtnSelected.value.has('none')
		) {
			valueBtnSelected.value.add(value);
		} else if (value === 'none') {
			valueBtnSelected.value.clear();
			valueBtnSelected.value.add(value);
		} else if (value !== 'none' && valueBtnSelected.value.has('none')) {
			valueBtnSelected.value.delete('none');
			valueBtnSelected.value.add(value);
		}
		authFormStore.allergies = valueBtnSelected.value;
	} else if (model === 'goal') {
		authFormStore.goal = value;
		valueBtnSelected.value.clear();
		valueBtnSelected.value.add(value);
	}
};

const isFormValid = computed(() => {
	if (
		authContent.value.content[countPages.value]?.inputs &&
		authContent.value.content[countPages.value]?.btns
	) {
		return (
			authContent.value.content[countPages.value].inputs?.every(
				(input) => authFormStore[input.model] !== ''
			) && valueBtnSelected.value.size !== 0
		);
	} else if (authContent.value.content[countPages.value]?.inputs) {
		return authContent.value.content[countPages.value]?.inputs?.every(
			(input) => authFormStore[input.model] !== ''
		);
	} else {
		return valueBtnSelected.value.size !== 0;
	}
});
</script>
<template>
	<Column :gap="'32'" :align="'start'" full-width class="auth-form">
		<template v-if="authContent.content[countPages]?.inputs">
			<Column full-width :gap="'16'" :align="'start'" class="auth-inputs">
				<template v-for="input in authContent.content[countPages].inputs">
					<Column full-width :gap="'8'" :align="'start'">
						<Typography :size="'s'">{{ input.name }}</Typography>
						<Input
							:type="input.model === 'password' ? 'password' : 'text'"
							v-model="authFormStore[input.model]"
							full-width
							@input="
								validationAuthFields(input.model, authFormStore[input.model])
							"
							:borderError="errorsFields[input.model].value.size !== 0"
						/>
						<template
							v-if="errorsFields[input.model]"
							v-for="error in errorsFields[input.model].value"
						>
							<Typography :size="'xs'" :color="'error-msg'">{{
								error
							}}</Typography>
						</template>
					</Column>
				</template>
			</Column>
		</template>
		<template v-if="authType === 'login'">
			<Typography
				@click="
					$router.push('/signup');
					validationAuthErrorClear('all');
				"
				class="link-signup"
				:size="'xs'"
			>
				Нет аккаунта? <u>Регистрация</u>
			</Typography>
		</template>
		<template v-if="authContent.content[countPages]?.btns">
			<Typography :size="'s'">{{
				authContent.content[countPages].title
			}}</Typography>
			<Column full-width :gap="'16'" :align="'start'">
				<template v-for="btn in authContent.content[countPages].btns">
					<Button
						@click="selectedBtn(btn.value, btn.model)"
						class="auth-btns"
						:class="{ 'selected-btn': valueBtnSelected.has(btn.value) }"
						>{{ btn.name }}
					</Button>
				</template>
			</Column>
		</template>
		<template v-if="errorsServer.size !== 0" v-for="error in errorsServer">
			<Typography :size="'xs'" :color="'error-msg'">{{ error }}</Typography>
		</template>
		<Button
			:is-right-arrow="authType === 'signup'"
			:variant-color="'purple-btn'"
			@click="handlerForm"
			:disabled="
				!isFormValid ||
				Object.values(errorsFields).some(
					(errorsSet) => errorsSet.value.size > 0
				)
			"
		>
			{{ authContent.content[countPages]?.btnText }}
		</Button>
	</Column>
</template>
<style scoped lang="css">
@import url('./AuthForm.css');
</style>
