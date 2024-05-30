<script setup lang="ts">
import { type PropType, computed } from "vue";
import LeftArrow from "@/shared/assets/icons/LeftArrow.vue";
import RightArrow from "@/shared/assets/icons/RightArrow.vue";

export type ButtonVariant = "round-btn" | "content-btn";
export type ButtonColor = "white-btn" | "green-btn" | "purple-btn";

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false,
  },
  variant: {
    type: String as PropType<ButtonVariant>,
    default: "content-btn",
  },
  variantColor: {
    type: String as PropType<ButtonColor>,
    default: "white-btn",
  },
  fullWidth: {
    type: Boolean,
  },
  isLeftArrow: {
    type: Boolean,
    default: false,
  },
  isRightArrow: {
    type: Boolean,
    default: false,
  },
});

const classes = computed(() => ({
  "full-width": props.fullWidth,
  disabled: props.disabled,
  [props.variant]: true,
  [props.variantColor]: true,
}));
</script>

<template>
  <button :disabled="disabled" type="button" :class="classes">
    <template v-if="isLeftArrow">
        <LeftArrow />
    </template>
    <slot></slot>
    <template v-if="isRightArrow">
        <RightArrow />
    </template>
  </button>
</template>

<style lang="css" scoped>
@import url("./Button.css");
</style>
