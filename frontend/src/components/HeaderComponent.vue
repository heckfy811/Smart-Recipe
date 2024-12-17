<template>
  <header>
    <nav class="navbar">
      <!-- Логотип -->
      <router-link to="/" class="logo">
        <div class="smart">Smart</div>.Recipe
      </router-link>

      <!-- Навбар -->
      <ul class="nav-links" v-if="!isCollapsed">
        <li>
          <router-link to="/" class="head" :class="{ click: currentRoute === '/' }">Главная</router-link>
        </li>
        <li>
          <router-link to="/recipes" class="head" :class="{ click: currentRoute === '/recipes' }">Рецепты</router-link>
        </li>
        <li>
          <router-link to="/plans" class="head" :class="{ click: currentRoute === '/plans' }">Планы питания</router-link>
        </li>
        <li>
          <router-link to="/about" class="head" :class="{ click: currentRoute === '/about' }">О нас</router-link>
        </li>
        <li>
          <router-link to="/profile" class="profile-btn" :class="{ click: currentRoute === '/profile' }">Профиль</router-link>
        </li>
      </ul>

      <!-- Выпадающий список -->
      <div 
        class="dropdown" 
        v-else
        @mouseenter="openDropdown" 
        @mouseleave="closeDropdown"
      >
        <button class="dropdown-button">
          {{ currentPageName }}
        </button>
        <ul class="dropdown-menu" v-show="isDropdownVisible">
          <li v-for="(link, index) in navLinks" :key="index" :style="dropdownItemStyle(index)">
            <router-link :to="link.path" @click="closeDropdown">{{ link.name }}</router-link>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>



<script>
import { ref, computed, onMounted, onUnmounted } from "vue";
import { useRoute } from "vue-router";

export default {
  name: "HeaderComponent",
  setup() {
    const route = useRoute();
    const isCollapsed = ref(false);
    const isDropdownVisible = ref(false);

    const navLinks = [
      { name: "Главная", path: "/" },
      { name: "Рецепты", path: "/recipes" },
      { name: "Планы питания", path: "/plans" },
      { name: "О нас", path: "/about" },
      { name: "Профиль", path: "/profile" },
    ];

    const currentRoute = computed(() => route.path);

    const currentPageName = computed(() => {
      const currentLink = navLinks.find((link) => link.path === route.path);
      return currentLink ? currentLink.name : "Меню";
    });

    const openDropdown = () => {
      isDropdownVisible.value = true;
    };

    const closeDropdown = () => {
      isDropdownVisible.value = false;
    };

    const updateNavbarState = () => {
  const navbar = document.querySelector(".navbar");
  const navLinksEl = document.querySelector(".nav-links");

  if (navbar && navLinksEl) {
    // Проверяем ширину окна и вычисляем, нужно ли скрывать навбар
    const navLinksWidth = navLinksEl.offsetWidth || 0;
    isCollapsed.value = navbar.offsetWidth < navLinksWidth + 150;
  } else {
    isCollapsed.value = window.innerWidth <= 768; // Восстанавливаем isCollapsed при изменении ширины
  }
};


    onMounted(() => {
      window.addEventListener("resize", updateNavbarState);
      updateNavbarState();
    });

    onUnmounted(() => {
      window.removeEventListener("resize", updateNavbarState);
    });

    const dropdownItemStyle = (index) => {
      const delay = index * 0.1;
      return {
        animation: isDropdownVisible.value
          ? `fadeIn 0.3s ease ${delay}s forwards`
          : `fadeOut 0.3s ease ${delay}s forwards`,
      };
    };

    return {
      currentRoute,
      currentPageName,
      isCollapsed,
      isDropdownVisible,
      navLinks,
      openDropdown,
      closeDropdown,
      dropdownItemStyle,
    };
  },
};
</script>


<style>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: white;
  width: 100%;
  box-sizing: border-box;
}

/* Логотип */
.logo {
  font-size: 1.5em;
  font-weight: 600;
  padding-left: 120px;
  text-decoration: none;
  color: #000000;
  transition: transform 0.3s ease;
  display: flex;
  align-items: center;
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
  color: #000; /* Основной цвет текста */
  transition: color 0.3s ease;
}

.nav-links .click {
  color: #12a370; /* Зелёный цвет для активной страницы */
}

.nav-links a:hover {
  color: #12a370; /* Зелёный цвет при наведении */
}

.nav-links .profile-btn {
  color: #12a370 !important; /* Постоянный зелёный цвет */
  border: 1px solid #12a370;
  border-radius: 20px;
  padding: 8px 15px;
  text-decoration: none;
  transition: color 0.3s ease, background-color 0.3s ease;
}

/* Подсветка кнопки Профиль при наведении */
.nav-links .profile-btn:hover {
  color: white !important; /* Текст становится белым */
  background-color: #12a370; /* Фон становится зелёным */
}

/* Выпадающий список */
.dropdown {
  position: relative;
  margin-top: 10px;
  width: 90%;
  text-align: center;
}

.dropdown-button {
  background-color: #12a370;
  color: white;
  padding: 12px 15px;
  border: none;
  border-radius: 25px;
  cursor: pointer;
  width: 100%;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  background-color: white;
  list-style: none;
  padding: 10px 0;
  margin: 0;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  z-index: 1000;
  display: flex;
  flex-direction: column;
}

/* Убираем нижнее подчеркивание у текста в выпадающем списке */
.dropdown-menu a {
  text-decoration: none; /* Убираем подчеркивание */
  color: #000; /* Основной цвет текста */
  transition: color 0.3s ease; /* Плавный переход цвета */
}

/* Подсветка текста зелёным при наведении */
.dropdown-menu a:hover {
  color: #12a370; /* Зелёный цвет при наведении */
}

/* Анимации */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeOut {
  from {
    opacity: 1;
    transform: translateY(0);
  }
  to {
    opacity: 0;
    transform: translateY(-10px);
  }
}

/* Мобилка */
@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    align-items: center;
  }

  .logo {
    margin: 0 auto;
    transform: scale(1.5);
    text-align: center;
    display: block;
    padding-left: 0px !important;
  }

  .nav-links {
    width: 100%;
    justify-content: center;
    margin-top: 15px;
  }

  .dropdown {
    width: 90%;
    margin-top: 10px;
    text-align: center;
  }

  .dropdown-button {
    width: 90%;
  }
}

</style>
