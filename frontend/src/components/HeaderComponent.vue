<template>
  <header>
    <nav class="navbar">
      <router-link to="/" class="logo">
        <div class="smart">Smart</div>.Recipe
      </router-link>
      <!-- Меню для широких экранов -->
      <ul class="nav-links" v-if="!isMobile">
        <li><router-link to="/" class="head" :class="{ click: currentRoute === '/' }">Главная</router-link></li>
        <li><router-link to="/recipes" class="head" :class="{ click: currentRoute === '/recipes' }">Рецепты</router-link></li>
        <li><router-link to="/plans" class="head" :class="{ click: currentRoute === '/plans' }">Планы питания</router-link></li>
        <li><router-link to="/about" class="head" :class="{ click: currentRoute === '/about' }">О нас</router-link></li>
        <li><router-link to="/profile" class="profile-btn" :class="{ click: currentRoute === '/profile' }">Профиль</router-link></li>
      </ul>
      <!-- Выпадающее меню для мобильных устройств -->
      <div class="dropdown" v-if="isMobile">
        <button class="dropdown-button">
          {{ currentPageName }}
        </button>
        <ul class="dropdown-menu">
          <li><router-link to="/" @click="closeDropdown">Главная</router-link></li>
          <li><router-link to="/recipes" @click="closeDropdown">Рецепты</router-link></li>
          <li><router-link to="/plans" @click="closeDropdown">Планы питания</router-link></li>
          <li><router-link to="/about" @click="closeDropdown">О нас</router-link></li>
          <li><router-link to="/profile" @click="closeDropdown">Профиль</router-link></li>
        </ul>
      </div>
    </nav>
  </header>
</template>

<script>
import { ref, computed } from "vue";
import { useRoute } from "vue-router";

export default {
  name: "HeaderComponent",
  setup() {
    const route = useRoute();
    const isMobile = ref(false);

    // Определяем текущий маршрут
    const currentRoute = computed(() => route.path);

    // Название текущей страницы
    const currentPageName = computed(() => {
      switch (route.path) {
        case "/":
          return "Главная";
        case "/recipes":
          return "Рецепты";
        case "/plans":
          return "Планы питания";
        case "/about":
          return "О нас";
        case "/profile":
          return "Профиль";
        default:
          return "Меню";
      }
    });

    // Обновляем состояние мобильного вида при изменении ширины экрана
    const updateViewMode = () => {
      isMobile.value = window.innerWidth <= 768;
    };

    window.addEventListener("resize", updateViewMode);
    updateViewMode(); // Инициализация

    return {
      currentRoute,
      currentPageName,
      isMobile,
    };
  },
};
</script>

<style scoped>
/* Основной стиль навбара */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: white;
}

.logo {
  font-size: 1.5em;
  font-weight: 600;
  text-decoration: none;
  color: #000000;
  display: flex;
  justify-content: center;
}

.smart {
  color: #12a370;
  display: inline;
}

.nav-links {
  display: flex;
  list-style: none;
}

.nav-links li {
  margin-left: 20px;
  display: flex;
  align-items: center;
}

.nav-links a {
  text-decoration: none;
  padding: 8px 15px;
}

.profile-btn {
  padding: 8px 15px;
  text-decoration: none;
  border: 1px solid #12a370;
  border-radius: 20px;
  color: #12a370;
}

/* Выпадающее меню */
.dropdown {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.dropdown-button {
  background-color: #12a370;
  color: white;
  padding: 8px 15px;
  border: none;
  border-radius: 20px;
  cursor: pointer;
}

.dropdown-menu {
  display: none;
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  background-color: white;
  list-style: none;
  padding: 10px 0;
  margin: 0;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  z-index: 1000;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.dropdown:hover .dropdown-menu {
  display: block;
  opacity: 1;
  visibility: visible;
  transform: translateX(-50%) translateY(10px);
}

.dropdown-menu li {
  padding: 10px 20px;
  text-align: center;
}

.dropdown-menu a {
  text-decoration: none;
  color: #000;
}

.dropdown-menu a:hover {
  color: #12a370;
}

/* Мобильная версия */
@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .logo {
    margin-bottom: 10px;
  }

  .nav-links {
    display: none;
  }

  .dropdown {
    width: 100%;
    text-align: center;
  }
}
</style>
