<template>
  <div id="app">
    <HeaderComponent v-if="!isAuthPage" />
    <router-view />
    <FooterComponent v-if="!isAuthPage" />
  </div>
</template>

<script>
import HeaderComponent from '@/components/HeaderComponent.vue';
import FooterComponent from '@/components/FooterComponent.vue';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

export default {
  components: {
    HeaderComponent,
    FooterComponent,
  },
  setup() {
    const route = useRoute();

    const isAuthPage = computed(() => route.meta.isAuthPage === "auth" || route.meta.isAuthPage === "register");

    return { isAuthPage };
  },
};
</script>

<style>
@import '/general-style.css';
</style>