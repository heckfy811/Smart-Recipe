<template>
  <div>
    <HeroComponent />
    <PopularRecipesComponent />
    <SubscriptionComponent />
    <BestRecipeComponent />
  </div>
</template>

<script>
import HeroComponent from "../components/HeroComponent.vue";
import PopularRecipesComponent from "../components/PopularRecipesComponent.vue";
import SubscriptionComponent from "../components/SubscriptionComponent.vue";
import BestRecipeComponent from "../components/BestRecipeComponent.vue";

export default {
  components: {
    HeroComponent,
    PopularRecipesComponent,
    SubscriptionComponent,
    BestRecipeComponent,
  },
};
</script>

<style scoped>
/* При необходимости добавьте уникальные стили для главной страницы */
</style>