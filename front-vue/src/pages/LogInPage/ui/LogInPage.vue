<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Row, Column, Typography } from '@/shared/ui';
import { ref } from 'vue';
import { type AuthContentType, AuthForm } from '@/features/AuthForm';

const authContent = ref<AuthContentType>();
const countPages = ref<number>(0);

const handleAuthContentChanged = (newAuthContent: AuthContentType) => {
	authContent.value = newAuthContent;
};
</script>
<template>
	<Page :header="'none'" :background="'auth-bckgr'">
		<Row :justify="'center'" class="login-container">
			<Column :justify="'center'" :align="'start'" class="left-side">
				<AuthForm
					:authType="'login'"
					@authContentChanged="handleAuthContentChanged"
				/>
			</Column>
			<Column
				:gap="'32'"
				:align="'center'"
				:justify="'center'"
				class="right-side"
			>
				<Typography
					:align="'center'"
					:size="'l'"
					:weight="500"
					:styleText="'italic'"
					>{{ authContent?.content[countPages].factTitle }}</Typography
				>
				<Typography :align="'center'" :size="'m'" :weight="100">{{
					authContent?.content[countPages].factSubtitle
				}}</Typography>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./LogInPage.css');
</style>
