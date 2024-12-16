<template>
  <section class="best-recipe">
    <h2>Лучший рецепт</h2>
    <div v-if="recipe" class="best-recipe-card">
      <div class="recipe-slider">
        <div class="main-slide">
          <img :src="slides[currentSlide] || placeholderImage" alt="Текущий слайд" />
          <button class="slider-btn prev" @click="prevSlide">‹</button>
          <button class="slider-btn next" @click="nextSlide">›</button>
        </div>
        <div class="thumbnails">
          <img
            v-for="(slide, index) in slides"
            :key="index"
            :src="slide || placeholderImage"
            :class="{ active: currentSlide === index }"
            @click="goToSlide(index)"
            alt="Миниатюра"
          />
        </div>
      </div>
      <div class="recipe-info">
        <h3>{{ recipe.title }}</h3>
        <p>{{ recipe.cookingTime }}</p>
        <p>{{ recipe.description }}</p>
        <router-link :to="`/recipe/${recipe.id}`" class="btn-small">Попробовать рецепт</router-link>
      </div>
    </div>
    <div v-else>Загрузка...</div>
  </section>
</template>

<script>
export default {
  name: "BestRecipeComponent",
  data() {
    return {
      recipe: null, // Данные рецепта
      slides: [],
      currentSlide: 0,
      placeholderImage: "@/assets/images/placeholder.png", // Заглушка для изображений
    };
  },
  mounted() {
    this.loadRecipeData();
  },
  methods: {
    async loadRecipeData() {
      try {
        const response = await import("@/assets/testing-data/test.json");
        this.recipe = response.recipe;
        this.slides = this.recipe.images || [];
        this.startSlider();
      } catch (error) {
        console.error("Ошибка загрузки данных рецепта:", error);
      }
    },
    nextSlide() {
      this.currentSlide = (this.currentSlide + 1) % this.slides.length;
    },
    prevSlide() {
      this.currentSlide = (this.currentSlide - 1 + this.slides.length) % this.slides.length;
    },
    goToSlide(index) {
      this.currentSlide = index;
    },
    startSlider() {
      if (this.slides.length > 1) {
        this.sliderInterval = setInterval(this.nextSlide, 5000); // Автоматическая смена каждые 5 секунд
      }
    },
  },
  beforeUnmount() {
    clearInterval(this.sliderInterval);
  },
};
</script>

<style scoped>
.best-recipe {
  margin-top: 120px;
  padding: 50px;
  background-color: white;
  animation: slideFromLeft 0.8s ease-out;
}

.best-recipe h2 {
  font-size: 2em;
  margin-bottom: 30px;
  text-align: left; /* Заголовок слева */
}

.best-recipe-card {
  display: flex;
  justify-content: space-between; /* Слайдер слева, текст справа */
  background-color: white;
  border-radius: 25px;
  box-shadow: 0 16px 16px rgba(0, 0, 0, 0.08);
  padding: 50px; /* Увеличены внутренние отступы карточки */
  text-align: left; /* Выравнивание текста */
  gap: 40px; /* Добавлен отступ между слайдером и текстом */
}

.recipe-slider {
  position: relative;
  flex: 1 1 25%; /* Уменьшили слайдер до четверти ширины карточки */
  margin-right: 20px;
}

.main-slide {
  position: relative;
  width: 100%;
  aspect-ratio: 4 / 3; /* Соотношение сторон 4:3 */
  border-radius: 25px; /* Равномерное закругление углов */
  overflow: hidden;
}

.main-slide img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 25px;
}

.slider-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(0, 0, 0, 0.5);
  border: none;
  color: white;
  font-size: 1.5em;
  padding: 5px 10px;
  border-radius: 50%;
  cursor: pointer;
  z-index: 2;
}

.slider-btn.prev {
  left: 10%;
}

.slider-btn.next {
  right: 3%;
}

.thumbnails {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-top: 10px;
}

.thumbnails img {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 10px;
  cursor: pointer;
  opacity: 0.7;
  transition: opacity 0.3s ease;
}

.thumbnails img.active {
  opacity: 1;
  transform: scale(1.1); /* Увеличиваем активный слайд */
  border: 3px solid #12a370; /* Толщина рамки */
  box-sizing: border-box; /* Учитываем размеры рамки */
}

.recipe-info {
  flex: 1 1 70%; /* Текст занимает больше места */
  max-width: 70%;
  margin-top: 20px; /* Отступ сверху */
}

.recipe-info h3 {
  font-size: 2em;
  margin-bottom: 20px; /* Увеличен отступ между заголовком и текстом */
}

.recipe-info p {
  font-size: 1em;
  color: #555;
  line-height: 1.5;
  margin: 10px 0;
}

.recipe-info .btn-small {
  display: inline-block;
  padding: 10px 20px;
  background-color: #12a370;
  border: 1px solid #12a370;
  color: white;
  text-decoration: none;
  border-radius: 25px;
  font-weight: bold;
  margin-top: 20px;
  transition: background-color 0.3s ease, color 0.3s ease;
  width: 50%; /* Адаптивная ширина относительно карточки */
  max-width: 200px; /* Максимальная ширина кнопки */
  text-align: center;
}

.recipe-info .btn-small:hover {
  color: #12a370;
  background-color: #ffffff;
}

/* Адаптивность */
@media (max-width: 768px) {
  .best-recipe-card {
    flex-direction: column;
    align-items: center; /* Центрируем содержимое */
    gap: 20px; /* Отступы между элементами на маленьких экранах */
  }

  .recipe-slider {
    flex: none;
    width: 80%; /* Слайдер занимает 80% ширины карточки */
    margin-right: 0;
    margin-bottom: 20px;
  }

  .recipe-info {
    max-width: 100%;
    text-align: center;
  }

  .recipe-info h3,
  .recipe-info p {
    text-align: center;
  }
  .recipe-info .btn-small {
    width: 100%; /* Увеличиваем ширину кнопки на маленьких экранах */
    padding: 10px 15px; /* Уменьшаем внутренние отступы */
    font-size: 0.9em; /* Уменьшаем шрифт */
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
