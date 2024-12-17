<template>
  <div>
    <!-- Секция "Немного о нас" -->
    <div class="about-us">
      <h2>Немного о нас</h2>
      <p>
        На нашем сайте мы стремимся сделать ваш кулинарный опыт простым, выгодным и удобным. Мы понимаем,
        как важно не только выбрать вкусный рецепт, но и найти, где приобрести ингредиенты по наиболее
        выгодной цене и с минимальными усилиями.
      </p>
    </div>

    <!-- Секция "Наши разработчики" -->
    <section class="developers">
      <h2>Наши разработчики</h2>
      <div class="dev-container">
        <div class="dev-grid">
          <div v-for="develop in develops" :key="develop.id" class="dev-card" @mouseover="hoveredDev = develop" @mouseleave="hoveredDev = null">
            <div class="image-container">
              <img :src="develop.image" :alt="develop.title" />
            </div>
            <h3>{{ develop.title }}</h3>
            <ul class="contact-list" :class="{ 'visible': hoveredDev === develop }">
              <li v-for="(contact, index) in develop.contacts" :key="index">
                <a :href="contact.url" target="_blank">{{ contact.platform }}</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

import img1 from "@/assets/images/goga.png"
import img2 from "@/assets/images/yana.png"
import img3 from "@/assets/images/artem.png"
import img4 from "@/assets/images/ruslan.png"
import img5 from "@/assets/images/stepa.png"

export default {
  
  name: "About.vue",
  data() {
    return {
      hoveredDev: null, // Хранит информацию о текущей карточке, на которую наведён курсор
      develops: [
        { 
          id: 1, 
          title: "Рудницкий-Антипов Георгий", 
          image: img1, 
          contacts: [
            { platform: "Telegram", url: "https://t.me/jojiikk" },
            { platform: "WhatsApp", url: "https://wa.me/123456789" },
            { platform: "Instagram", url: "https://instagram.com/goga" }
          ]
        },
        { 
          id: 2, 
          title: "Сопина Яна", 
          image: img2, 
          contacts: [
            { platform: "Telegram", url: "https://t.me/sp_yanaa" },
            { platform: "WhatsApp", url: "https://wa.me/987654321" },
            { platform: "Instagram", url: "https://instagram.com/yana" }
          ]
        },
        { 
          id: 3, 
          title: "Коротких Артем", 
          image: img3, 
          contacts: [
            { platform: "Telegram", url: "https://t.me/krtkkhrtm" },
            { platform: "WhatsApp", url: "https://wa.me/112233445" },
            { platform: "Instagram", url: "https://instagram.com/artem" }
          ]
        },
        { 
          id: 4, 
          title: "Муртазин Руслан", 
          image: img4, 
          contacts: [
            { platform: "Telegram", url: "https://t.me/i_hate_gingers" },
            { platform: "WhatsApp", url: "https://wa.me/556677889" },
            { platform: "Instagram", url: "https://instagram.com/ruslan" }
          ]
        },
        { 
          id: 5, 
          title: "Соколов Степан", 
          image: img5, 
          contacts: [
            { platform: "Telegram", url: "https://t.me/ssksokolov" },
            { platform: "WhatsApp", url: "https://wa.me/998877665" },
            { platform: "Instagram", url: "https://instagram.com/stepa" }
          ]
        },
      ],
    };
  },
};
</script>


<style scoped>
.dev-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-right: 50px;
}

.dev-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-auto-rows: auto;
  gap: 30px;
  max-width: 100%;
}

.dev-card {
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 16px 16px rgba(0, 0, 0, 0.08);
  text-align: center;
  width: 390px; /* Фиксированная ширина карточки */
  height: 400px; /* Фиксированная высота карточки */
  margin-bottom: 50px;
  overflow: hidden; /* Обрезаем содержимое, выходящее за пределы карточки */
  transition: all 0.3s ease; /* Плавный переход для всех изменений */
  box-sizing: border-box; /* Учитываем отступы и границы в ширине */
  position: relative;
}
.image-container {
  width: 100%;
  height: 300px;
  overflow: hidden;
  position: relative;
}

.dev-card img {
  width: 100%;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}

.dev-card h3 {
  font-weight: 500;
  padding: 10px 0; /* Плавный переход для увеличения области с подписью */
  
}

.contact-list {
  list-style-type: none; /* Убираем маркеры списка */
  padding-left: 30px; /* Убираем отступы */
  padding-bottom: 30px;
  margin: 0; /* Убираем внешние отступы */
  opacity: 0; /* Скрываем список по умолчанию */
  transform: translateY(20px); /* Сдвигаем список вниз */
  transition: opacity 0.3s ease, transform 0.3s ease; /* Плавный переход для opacity и transform */
  position: absolute; /* Позиционируем список относительно карточки */
  bottom: 10px; /* Располагаем список внизу карточки */
  transform: translateY(50%); /* Центрируем по горизонтали */
}

.contact-list.visible {
  opacity: 1; /* Делаем список видимым */
  transform: translateY(0); /* Возвращаем список на место */
}

.contact-list li {
  margin: 5px 0; /* Отступы между элементами списка */
}

.contact-list a {
  text-decoration: none; /* Убираем подчеркивание */
  color: black;
  transition: color 0.3s ease; /* Плавный переход для цвета */
}

.contact-list a:hover {
  color: #0c6445; /* Цвет ссылок при наведении */
}

.dev-card:hover img {
  opacity: 0.7; /* Делаем изображение прозрачнее при наведении */
  transform: translateY(-50px);
  transition: transform 0.3s ease;

}

.dev-card:hover h3{
  opacity: 0;
  transition: opacity 0.3s ease;
}

.dev-card:hover .contact-list{
  opacity: 1;
  font-weight: bold;
}
</style>