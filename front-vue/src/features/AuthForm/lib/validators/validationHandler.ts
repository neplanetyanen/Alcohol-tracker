import { ref, type Ref } from 'vue';

import { authValidationMessages } from './messages';

interface ErrorsFields {
	[key: string]: Ref<Set<string>>;
}

export const errorsFields: ErrorsFields = {
	username: ref(new Set<string>()),
	password: ref(new Set<string>()),
	email: ref(new Set<string>()),
	height: ref(new Set<string>()),
	weight: ref(new Set<string>()),
};
export const errorsServer = ref<Set<string>>(new Set());

const validationUsername = (username: string) => {
	const regexp = /^[a-zA-Z0-9._-]+$/;
	if ((username.length > 30 || username.length < 5) && username !== '') {
		errorsFields.username.value.add(
			authValidationMessages.INVALID_LENGTH.USERNAME
		);
	} else {
		errorsFields.username.value.delete(
			authValidationMessages.INVALID_LENGTH.USERNAME
		);
	}
	if (!regexp.test(username) && username !== '') {
		errorsFields.username.value.add(
			authValidationMessages.REGEXP_MISMATCH.USERNAME
		);
	} else {
		errorsFields.username.value.delete(
			authValidationMessages.REGEXP_MISMATCH.USERNAME
		);
	}
};

const validationPassword = (password: string) => {
	const regexp = /^[a-zA-Z0-9.@_-]+$/;
	if (password.length < 8 && password !== '') {
		errorsFields.password.value.add(
			authValidationMessages.INVALID_LENGTH.PASSWORD
		);
	} else {
		errorsFields.password.value.delete(
			authValidationMessages.INVALID_LENGTH.PASSWORD
		);
	}
	if (!regexp.test(password) && password !== '') {
		errorsFields.password.value.add(
			authValidationMessages.REGEXP_MISMATCH.PASSWORD
		);
	} else {
		errorsFields.password.value.delete(
			authValidationMessages.REGEXP_MISMATCH.PASSWORD
		);
	}
};

const validationEmail = (email: string) => {
	const regexp = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/;
	if (!regexp.test(email)) {
		errorsFields.email.value.add(authValidationMessages.REGEXP_MISMATCH.EMAIL);
	} else {
		errorsFields.email.value.delete(
			authValidationMessages.REGEXP_MISMATCH.EMAIL
		);
	}
};

const validationHeight = (height: string) => {
	const regexp = /^[0-9]+$/;
	if (Number(height) <= 0 || !regexp.test(height)) {
		errorsFields.height.value.add(authValidationMessages.INVALID_LENGTH.HEIGHT);
	} else {
		errorsFields.height.value.delete(
			authValidationMessages.INVALID_LENGTH.HEIGHT
		);
	}
};

const validationWeight = (weight: string) => {
	const regexp = /^[0-9]+$/;
	if (Number(weight) <= 0 || !regexp.test(weight)) {
		errorsFields.weight.value.add(authValidationMessages.INVALID_LENGTH.WEIGHT);
	} else {
		errorsFields.weight.value.delete(
			authValidationMessages.INVALID_LENGTH.WEIGHT
		);
	}
};

export function validationAuthFields(fieldName: string, value: string): void {
	switch (fieldName) {
		case 'username':
			validationUsername(value);
			break;
		case 'password':
			validationPassword(value);
			break;
		case 'email':
			validationEmail(value);
			break;
		case 'height':
			validationHeight(value);
			break;
		case 'weight':
			validationWeight(value);
			break;
		default:
			errorsServer.value.add(`Validation error in the field ${fieldName}`);
	}
}

export function validationAuth(error: string | undefined) {
	switch (error) {
		case 'USERNAME_ALREADY_EXISTS':
			errorsServer.value.add(authValidationMessages.USERNAME_ALREADY_EXISTS);
			break;
		case 'EMAIL_ALREADY_EXISTS':
			errorsServer.value.add(authValidationMessages.EMAIL_ALREADY_EXISTS);
			break;
		case 'USERNAME_NOT_FOUND':
			errorsServer.value.add(authValidationMessages.USERNAME_NOT_FOUND);
			break;
		case 'INCORRECT_PASSWORD':
			errorsServer.value.add(authValidationMessages.INCORRECT_PASSWORD);
			break;
		default:
			errorsServer.value.add(String(error));
	}
}

export function validationAuthErrorClear(field: string) {
	switch (field) {
		case 'username':
			errorsServer.value.delete(authValidationMessages.USERNAME_ALREADY_EXISTS);
			errorsServer.value.delete(authValidationMessages.USERNAME_NOT_FOUND);
			break;
		case 'email':
			errorsServer.value.delete(authValidationMessages.EMAIL_ALREADY_EXISTS);
			break;
		case 'password':
			errorsServer.value.delete(authValidationMessages.INCORRECT_PASSWORD);
			break;
		case 'all':
			errorsServer.value.clear();
			break;
		default:
			break;
	}
}
