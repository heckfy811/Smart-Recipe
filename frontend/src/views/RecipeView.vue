<template>
  <div>
    <RecipeComponent :recipe="recipe" />
    <RecipeIngredientsComponent :recipe="recipe" />
    <RecipeCardComponent :recipe="recipe" />
  </div>
</template>

<script>
import RecipeComponent from "@/components/RecipeComponent.vue";
import RecipeIngredientsComponent from "@/components/RecipeIngredientsComponent.vue";
import RecipeCardComponent from "@/components/RecipeCardComponent.vue";
import recipeData from "@/assets/testing-data/testRecipes.json";
import { useRoute } from "vue-router";

export default {
  components: {
    RecipeComponent,
    RecipeIngredientsComponent,
    RecipeCardComponent,
  },
  data() {
    return {
      recipe: null, // Текущий рецепт
    };
  },
  computed: {
    recipes() {
      return recipeData; // Все рецепты
    },
  },
  methods: {
    loadRecipe(id) {
      this.recipe = this.recipes.find((recipe) => recipe.id === parseInt(id, 10));
    },
  },
  mounted() {
    const route = useRoute();
    this.loadRecipe(route.params.id);
  },
  watch: {
    "$route.params.id"(newId) {
      this.loadRecipe(newId);
    },
  },
};
</script>