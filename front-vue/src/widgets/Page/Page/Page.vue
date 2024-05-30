<script setup lang="ts">
import Header from '../Header/Header.vue';
import Content from '../Content/Content.vue';
import Footer from '../Footer/Footer.vue';
import { type HeaderVariants } from '../Header/Header.vue';
import { type BackgroundColorVariants } from '../Content/Content.vue';
import type { PropType } from 'vue';

const props = defineProps({
	header: {
		type: String as PropType<HeaderVariants>,
		default: 'unauth',
	},
	background: {
		type: String as PropType<BackgroundColorVariants>,
		default: 'content-bckgr',
	},
});
</script>
<template>
	<div>
		<Header v-if="props.header !== 'none'" :header="props.header"></Header>
		<Content :background="props.background">
			<slot></slot>
		</Content>
		<Footer />
	</div>
</template>
<style lang="css" scoped>
@import url('./Page.css');
</style>
