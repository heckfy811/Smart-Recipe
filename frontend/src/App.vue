<template>
  <div id="app">
    <HeaderComponent v-if="!isAuthPage" />
    <router-view />
    <FooterComponent v-if="!isAuthPage" />
  </div>
</template>

<script>
import HeaderComponent from "@/components/HeaderComponent.vue";
import FooterComponent from "@/components/FooterComponent.vue";
import { useRoute } from "vue-router";
import { computed } from "vue";

export default {
  components: {
    HeaderComponent,
    FooterComponent,
  },
  setup() {
    const route = useRoute();
    const isAuthPage = computed(() => {
      return route.path === "/auth" || route.path === "/register";
    });

    return { isAuthPage };
  },
};
</script>

<style>
@import '/general-style.css';
</style>