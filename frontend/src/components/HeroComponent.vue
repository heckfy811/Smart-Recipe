<template>
  <section class="hero">
    <div class="hero-content">
      <h1>Вкусные рецепты & Выгодные покупки</h1>
      <p>Выбирайте, чем себя побаловать, а мы поможем сделать с наибольшей выгодой и удобством</p>
      <div class="hero-buttons">
        <router-link to="/recipes" class="btn">Рецепты</router-link>
        <router-link to="/subscription" class="btn-outline">Подписка</router-link>
      </div>
    </div>
    <!-- Отображаем изображения только если они помещаются -->
    <div class="hero-images" v-if="imagesVisible">
      <img
        v-for="(image, index) in images"
        :key="index"
        :src="image.src"
        :alt="image.alt"
        class="image"
      />
    </div>
  </section>
</template>

<script>
import img1 from "@/assets/images/img1.png";
import img2 from "@/assets/images/img2.png";
import img3 from "@/assets/images/img3.png";

export default {
  name: "HeroComponent",
  data() {
    return {
      images: [
        { src: img1, alt: "Изображение 1" },
        { src: img2, alt: "Изображение 2" },
        { src: img3, alt: "Изображение 3" },
      ],
      imagesVisible: true,
    };
  },
  mounted() {
    this.updateImagesVisibility();
    window.addEventListener("resize", this.updateImagesVisibility);
  },
  beforeUnmount() {
    window.removeEventListener("resize", this.updateImagesVisibility);
  },
  methods: {
    updateImagesVisibility() {
      const hero = document.querySelector(".hero");
      const content = document.querySelector(".hero-content");
      const heroWidth = hero.offsetWidth;
      const contentWidth = content.offsetWidth;
      this.imagesVisible = heroWidth - contentWidth > 300; // Убираем изображения, если они не помещаются
    },
  },
};
</script>

<style scoped>
.hero {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 50px;
  padding-right: 0px;
  background-color: #fff;
  position: relative;
}

.hero-content {
  max-width: 600px;
  padding-left: 100px;
}

.hero-content h1 {
  font-size: 2.5em;
  margin-bottom: 20px;
}

.hero-content p {
  font-size: 1.2em;
  margin-bottom: 20px;
}

.hero-buttons .btn,
.hero-buttons .btn-outline {
  padding: 12px 20px;
  border: none;
  border-radius: 25px;
  text-decoration: none;
  font-size: 1em;
  font-weight: bold;
  cursor: pointer;
  margin-right: 15px;
}

.hero-buttons .btn {
  background-color: #12a370;
  color: #fff;
  transition: background-color 0.3s ease, transform 0.3s ease;
}

.hero-buttons .btn:hover {
  background-color: #0e8a5d;
  transform: scale(1.05);
}

.hero-buttons .btn-outline {
  background-color: transparent;
  border: 2px solid #12a370;
  color: #12a370;
  transition: color 0.3s ease, background-color 0.3s ease, transform 0.3s ease;
}

.hero-buttons .btn-outline:hover {
  background-color: #12a370;
  color: #fff;
  transform: scale(1.05);
}

/* Анимация изображений */
.hero-images {
  display: flex;
  flex-wrap: nowrap; /* Запрещаем перенос строк */
  overflow: hidden; /* Прячем содержимое за краем */
}

.hero-images .image {
  max-width: 100%;
  border-radius: 15px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  margin-right: 10px; /* Расстояние между изображениями */
  transform: translateX(100%); /* Исходное положение - за правым краем */
  opacity: 0; /* Прозрачность для анимации */
  animation: slideIn 1s ease-out forwards; /* Анимация "выезда" */
  animation-delay: calc(var(--index, 0) * 0.2s); /* Задержка для последовательного появления */
}

/* Анимация */
@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

/* Адаптивность */
@media (max-width: 1024px) {
  .hero {
    flex-direction: column;
    align-items: center;
    text-align: center;
    padding: 35px !important;
  }

  .hero-content {
    padding-left: 0px;
    padding-right: 0px;
    max-width: 100%;
    margin-bottom: 20px;
    align-items: center;
    justify-content: center;
  }

  .hero-content h1,
  .hero-content p {
    text-align: center;
    margin-left: 0;
    margin-right: 0;
  }

  .hero-content h1 {
    font-size: 2em;
  }

  .hero-content p {
    font-size: 1em;
  }

  .hero-buttons .btn,
  .hero-buttons .btn-outline {
    font-size: 0.9em;
    padding: 10px 15px;
  }
}

@media (max-width: 768px) {
  .hero-content h1 {
    font-size: 1.8em;
  }

  .hero-content p {
    font-size: 0.9em;
  }

  .hero-buttons .btn,
  .hero-buttons .btn-outline {
    font-size: 0.8em;
    padding: 8px 12px;
  }
}
</style>
