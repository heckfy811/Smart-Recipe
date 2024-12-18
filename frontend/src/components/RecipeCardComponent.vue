<template>
  <section class="best-recipe" v-if="recipe">
    <div class="best-recipe-card">
      <div class="recipe-slider">
        <div class="main-slide">
          <img :src="recipe.image" :alt="recipe.title" />
        </div>
      </div>
      <div class="recipe-info">
        <h3>{{ recipe.title }}</h3>
        <p>Время готовки: {{ recipe.cookingTime }}</p>
        <p>Ориентировочная стоимость: {{ recipe.price }}₽</p>
        <div class="action-buttons">
          <router-link to="/near" class="btn-small">Рассчитать стоимость</router-link>
          <button class="btn-small" @click="markAsTried">Рецепт опробован</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно для оценки рецепта -->
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <h2>Оцените рецепт</h2>
        <div class="sstars">
          <span
            v-for="n in 5"
            :key="n"
            @click="rateRecipe(n)"
            :class="{ 'filled': n <= rating }"
            class="sstar"
          >&#9733;</span>
        </div>
        <textarea v-model="reviewText" placeholder="Напишите ваш отзыв..."></textarea>
        <div class="modal-actions">
          <button @click="submitReview">Отправить</button>
          <button @click="closeModal">Отмена</button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import recipesData from "@/assets/testing-data/testRecipes.json";

export default {
  name: "RecipeCardComponent",
  props: ["recipe"],
  data() {
    return {
      showModal: false,
      rating: 0,
      reviewText: "",
    };
  },
  methods: {
    markAsTried() {
      this.showModal = true;
      let triedRecipes = JSON.parse(localStorage.getItem("triedRecipes")) || [];
      if (!triedRecipes.includes(this.recipe.id)) {
        triedRecipes.push(this.recipe.id);
        localStorage.setItem("triedRecipes", JSON.stringify(triedRecipes));
      }
    },
    closeModal() {
      this.showModal = false;
      this.resetReview();
    },
    rateRecipe(n) {
      this.rating = n;
    },
    submitReview() {
      if (this.rating === 0) {
        alert("Пожалуйста, выберите количество звёзд.");
        return;
      }

      const recipeIndex = recipesData.findIndex(
        (r) => r.id === this.recipe.id
      );
      if (recipeIndex !== -1) {
        const currentRating = recipesData[recipeIndex].rating || 0;
        recipesData[recipeIndex].rating = parseFloat(
          ((currentRating + this.rating) / 2).toFixed(1)
        );
      }

      alert(`Ваш отзыв: ${this.reviewText}\nРейтинг: ${this.rating} ★`);
      this.closeModal();
    },
    resetReview() {
      this.rating = 0;
      this.reviewText = "";
    },
  },
};
</script>

<style scoped>
.best-recipe {
  margin-top: 20px;
  padding: 20px;
  background-color: white;
  animation: slideFromLeft 0.8s ease-out;
}

.best-recipe-card {
  display: flex;
  justify-content: space-between;
  background-color: white;
  border-radius: 25px;
  box-shadow: 0 16px 16px rgba(0, 0, 0, 0.08);
  padding: 40px;
  text-align: left;
  gap: 30px;
}

.recipe-slider {
  flex: 1 1 25%;
  margin-right: 20px;
}

.main-slide img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 25px;
}

.recipe-info {
  flex: 1 1 70%;
  max-width: 70%;
  margin-top: 20px;
}

.recipe-info h3 {
  font-size: 2em;
  margin-bottom: 20px;
}

.recipe-info p {
  font-size: 1em;
  color: #555;
  line-height: 1.5;
  margin: 10px 0;
}

.action-buttons {
  display: flex;
  gap: 15px;
}

.recipe-info .btn-small {
  display: inline-block;
  padding: 10px 20px;
  background-color: #12a370;
  border: 1px solid #12a370;
  color: white;
  text-decoration: none;
  border-radius: 25px;
  margin-top: 20px;
  transition: background-color 0.3s ease, color 0.3s ease;
  width: auto;
  max-width: 200px;
  text-align: center;
  white-space: nowrap;
  font-size: 1em;
}

.recipe-info .btn-small:hover {
  color: #12a370;
  background-color: #ffffff;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
  width: 400px;
}

.sstars {
  font-size: 2em;
  color: lightgray;
  cursor: pointer;
}

.sstar {
  display: inline-block;
  transition: color 0.3s;
}

.sstar.filled {
  color: gold;
}

textarea {
  width: 100%;
  margin-top: 15px;
  height: 80px;
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  resize: none;
}

.modal-actions button {
  margin: 10px;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.modal-actions button:first-child {
  background-color: #12a370;
  color: white;
}

.modal-actions button:last-child {
  background-color: #f44336;
  color: white;
}

/* Адаптивность */
@media (max-width: 768px) {
  .best-recipe-card {
    flex-direction: column;
    align-items: center;
    gap: 20px;
  }

  .recipe-slider {
    width: 80%;
    margin-right: 0;
  }

  .recipe-info {
    max-width: 100%;
    text-align: center;
  }

  .recipe-info .btn-small {
    width: 100%;
    padding: 10px 15px;
    font-size: 0.9em;
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
