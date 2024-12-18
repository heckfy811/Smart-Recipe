<template>
  <div v-if="recipe" class="recipe-wrapper">
    <!-- Картинка рецепта -->
    <div class="recipe-image">
      <img :src="recipe.image" :alt="recipe.title" />
    </div>

    <!-- Секция "О рецепте" -->
    <div class="about-recipe">
      <h2>О рецепте</h2>
      <p>{{ recipe.about }}</p>
    </div>
  </div>
  <div v-else>
    <p>Рецепт не найден.</p>
  </div>
</template>

<script>
import recipeData from "@/assets/testing-data/testRecipes.json";
import { useRoute } from "vue-router";

export default {
  name: "RecipeComponent",
  data() {
    return {
      recipe: null, // Данные рецепта
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

<style scoped>
/* Контейнер для изображения и описания */
.recipe-wrapper {
  display: flex;
  justify-content: flex-start; /* Располагаем элементы слева направо */
  align-items: center; /* Центрируем элементы по вертикали */
  gap: 20px; /* Расстояние между изображением и описанием */
  padding: 20px;
  animation: slideFromLeft 1s ease-out; /* Анимация выезжания */
}

/* Стили для изображения */
.recipe-image img {
  max-width: 400px; /* Ограничиваем ширину изображения */
  height: auto;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

/* Секция "О рецепте" */
.about-recipe {
  max-width: 50%; /* Ограничиваем ширину описания */
  background: #f9f9f9;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.about-recipe h2 {
  margin-bottom: 10px;
  font-size: 1.5em;
  color: #333;
}

.about-recipe p {
  font-size: 1em;
  color: #555;
  line-height: 1.6;
}

/* Адаптивность */
@media (max-width: 768px) {
  .recipe-wrapper {
    flex-direction: column; /* Элементы располагаются друг под другом */
    align-items: center; /* Центрируем элементы по горизонтали */
  }

  .recipe-image img {
    max-width: 90%; /* Увеличиваем ширину изображения */
  }

  .about-recipe {
    max-width: 90%; /* Ширина описания адаптируется */
    text-align: center;
  }
}

@keyframes slideFromLeft {
  from {
    transform: translateX(-100px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}
</style>
