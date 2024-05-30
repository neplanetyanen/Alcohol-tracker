<script setup lang="ts">
import { Page } from '@/widgets/Page';
import { Row, Column, Typography } from '@/shared/ui';
import { ref } from 'vue';
import { type AuthContentType, AuthForm } from '@/features/AuthForm';
import Table from '@/shared/assets/figures/Table.vue';

const authContent = ref<AuthContentType>();
const countPages = ref<number>(0);

const handleAuthContentChanged = (newAuthContent: AuthContentType) => {
	authContent.value = newAuthContent;
};

const handlePageChanged = (newCountPages: number) => {
	countPages.value = newCountPages;
};
</script>
<template>
	<Page :header="'none'" :background="'auth-bckgr'">
		<Row :justify="'center'" class="signup-container">
			<Column :justify="'center'" :align="'start'" class="left-side">
				<AuthForm
					:authType="'signup'"
					@authContentChanged="handleAuthContentChanged"
					@pageChanged="handlePageChanged"
				/>
			</Column>
			<Column
				:gap="'16'"
				:align="'center'"
				:justify="'center'"
				class="right-side"
			>
				<Typography
					:align="'center'"
					:size="'l'"
					:weight="500"
					:styleText="'italic'"
					>{{ authContent?.content[countPages]?.factTitle }}</Typography
				>
				<Typography
					:align="'center'"
					:size="'s'"
					:weight="100"
					v-html="authContent?.content[countPages]?.factSubtitle"
				></Typography>
				<template
					v-if="
						authContent?.content[countPages]?.btns &&
						authContent?.content[countPages]?.inputs
					"
				>
					<Table />
				</template>
			</Column>
		</Row>
	</Page>
</template>
<style scoped lang="css">
@import url('./SignUpPage.css');
</style>
