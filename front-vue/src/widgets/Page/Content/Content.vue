<script setup lang="ts">
import type { PropType } from 'vue';

export type BackgroundColorVariants = 'home-bckgr' | 'auth-bckgr' | 'content-bckgr' | 'edit-bckgr';
const props = defineProps({
    background: {
        type: String as PropType<BackgroundColorVariants>,
        default: 'content-bckgr'
    }
});
</script>
<template>
    <div :class="{'content': true, [props.background]: true}">
        <slot></slot>
    </div>
</template>
<style lang="css" scoped>
@import url('./Content.css');
</style>