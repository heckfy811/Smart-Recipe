<template>
  <div>
    <div class="recipe">
      <h1>Тушёная говядина</h1>
      <br>
      <img src="../assets/images/beef-background.png">
    </div>

    <!-- Секция "О рецепте" -->
    <div class="about-recipe">
        <h2>О рецепте</h2>
        <br>
        <p>Вьетнамская тушёная говядина (Bo Kho) — это ароматное блюдо, приготовленное с использованием специй, таких как имбирь, лимонник, звёздчатый анис, и овощей (морковь, лук). Оно отличается низким содержанием жиров, благодаря тушению, и высоким уровнем белка, что делает его идеальным для сбалансированного питания (ПП). Говядина содержит важные витамины и минералы, такие как железо и цинк. Специи и травы добавляют блюду уникальный вкус и ускоряют метаболизм. Тушение позволяет сохранить полезные вещества продуктов, делая это блюдо одновременно вкусным и питательным.</p>
    </div>

    <!-- Секция "Рецепт и ингридиенты" -->
    <div class="ingridients">
      <h2>Рецепт и ингридиенты</h2>
      <br>
      <p>Ингредиенты для вьетнамской тушёной говядины (Bo Kho):</p>
      <ul>
        <li>500 г говядины (лопатка или грудинка)</li>
        <li>1 луковица (нарезанная крупно)</li>
        <li>2 моркови (нарезанные на кусочки)</li>
        <li>3 зубчика чеснока (измельчённые)</li>
        <li>1 стебель лимонника (размятый)</li>
        <li>1 кусочек имбиря (2-3 см, нарезанный на тонкие ломтики)</li>
        <li>2-3 звёздочки бадьяна (звёздчатый анис)</li>
        <li>1 палочка корицы</li>
        <li>2 столовые ложки рыбного соуса</li>
        <li>1 столовая ложка соевого соуса</li>
        <li>1 столовая ложка томатной пасты</li>
        <li>1 чайная ложка сахара</li>
        <li>500 мл говяжего бульона или воды</li>
        <li>2 столовые ложки растительного масла</li>
        <li>Соль и перец по вкусу</li>
        <li>Свежий кориандр для подачи (по желанию)</li>
      </ul>
      <br>
      <p>Рецепт:</p>
      <ol>
        <li>Нарежьте говядину на крупные кубики (около 3-4 см). Посолите и поперчите по вкусу.  </li>
        <li>В глубокой кастрюле или сковороде разогрейте масло на среднем огне и обжарьте кусочки говядины до коричневого цвета, примерно 5-7 минут. Переложите говядину на тарелку.</li>
        <li>В той же кастрюле обжарьте лук, чеснок, имбирь и лимонник в течение 2-3 минут, пока они не станут ароматными.</li>
        <li>Добавьте томатную пасту, рыбный соус, соевый соус, сахар, бадьян и корицу. Перемешайте и готовьте ещё 1 минуту.</li>
        <li>Верните говядину в кастрюлю, добавьте бульон (или воду) и доведите до кипения.</li>
        <li>Уменьшите огонь, накройте крышкой и тушите на слабом огне 1,5–2 часа, пока говядина не станет мягкой.</li>
        <li>За 30 минут до готовности добавьте морковь и продолжайте тушить до её мягкости.</li>
        <li>Перед подачей украсьте свежим кориандром.</li>
      </ol>
    </div>
    <section class="best-recipe">
        <div class="best-recipe-card">
            <img src="../assets/images/beef.png" alt="Тушеная говядина" class="best-recipe-card-img">
            <div class="additional-images">
              <img src="../assets/images/beef.png" alt="Шаг 1">
              <img src="../assets/images/beef.png" alt="Шаг 2">
              <img src="../assets/images/beef.png" alt="Шаг 3">
            </div>
            <div class="recipe-info">
                <h3>Тушеная говядина</h3>
                <p>Время готовки 180 минут</p>
                <p>Ориентировочная стоимость блюда составляет: 2431</p>
                <router-link to="/near" class="btn-small">Рассчитать полную корзину</router-link>
                <div class="stars">
                  <h5>4.5</h5>
                  <span>★★★★☆</span>
                </div>
                
                <button class="feedback" @click="openModal">Оставить отзыв</button>
            </div>
        </div>
    </section>
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <h2>Оставить отзыв</h2>
        <div class="sstars">
          <h
              v-for="n in 5"
              :key="n"
              @click="rate(n)"
              :class="{ 'filled': n <= rating }"
              class="sstar"
            >&#9733;
          </h>
        </div>
        <textarea v-model="reviewText" placeholder="Напишите ваш отзыв..."></textarea>
        <div class="modal-actions">
          <button @click="submitReview">Отправить</button>
          <button @click="closeModal">Отмена</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "RecipeComponent.vue",
  data() {
    return {
      showModal: false, // Управляет показом модального окна
      rating: 0, // Количество выбранных звёзд
      reviewText: "", // Текст отзыва
    };
  },
  methods: {
    openModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
      this.resetReview();
    },
    rate(n) {
      this.rating = n; // Устанавливает рейтинг при клике на звезду
    },
    submitReview() {
      if (this.rating === 0) {
        alert("Пожалуйста, выберите количество звёзд.");
        return;
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
  .feedback {
    color: #12A370;
    cursor: pointer;
    text-decoration: underline;
    background: none;
    border: none;
    font-size: 15px;
    text-align: left;
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
  color: lightgray; /* Цвет невыбранных звёзд */
  cursor: pointer;
}

.sstar {
  display: inline-block;
  transition: color 0.3s;
}

.sstar.filled {
  color: gold; /* Цвет выбранных звёзд */
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
    background-color: #12A370;
    color: white;
  }

  .modal-actions button:last-child {
    background-color: #f44336;
    color: white;
  }
</style>
