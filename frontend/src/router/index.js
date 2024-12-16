import { createRouter, createWebHistory } from 'vue-router';

// Импорт компонентов для каждой страницы
import MainView from '@/views/MainView.vue';
import AboutView from '@/views/AboutView.vue';
import AuthView from '@/views/AuthView.vue';
import FavoriteView from '@/views/FavoriteView.vue';
import HistoryView from '@/views/HistoryView.vue';
import NearView from '@/views/NearView.vue';
import PlanView from '@/views/PlanView.vue';
import PlansView from '@/views/PlansView.vue';
import ProfileView from '@/views/ProfileView.vue';
import RecipeView from '@/views/RecipeView.vue';
import RecipesView from '@/views/RecipesView.vue';
import RegisterComponent from '../components/RegisterComponent.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainView,
    meta: { requiresAuth: true }, // Главная требует авторизации
  },
  {
    path: '/about',
    name: 'About',
    component: AboutView,
    meta: { requiresAuth: true },
  },
  {
    path: "/auth",
    name: "Auth",
    component: () => import("../views/AuthView.vue"),
    meta: { isAuthPage: "auth" }, // Указываем, что это страница авторизации
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/RegisterView.vue"),
    meta: { isAuthPage: "register" }, // Указываем, что это страница регистрации
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: MainView, // Отображаем главную страницу, если маршрут не найден
  },
  {
    path: '/favorite',
    name: 'Favorite',
    component: FavoriteView,
    meta: { requiresAuth: true },
  },
  {
    path: '/history',
    name: 'History',
    component: HistoryView,
    meta: { requiresAuth: true },
  },
  {
    path: '/near',
    name: 'Near',
    component: NearView,
    meta: { requiresAuth: true },
  },
  {
    path: '/plan',
    name: 'Plan',
    component: PlanView,
    meta: { requiresAuth: true },
  },
  {
    path: '/plans',
    name: 'Plans',
    component: PlansView,
    meta: { requiresAuth: true },
  },
  {
    path: '/profile',
    name: 'Profile',
    component: ProfileView,
    meta: { requiresAuth: true },
  },
  {
    path: '/recipe',
    name: 'Recipe',
    component: RecipeView,
    meta: { requiresAuth: true },
  },
  {
    path: '/recipes',
    name: 'Recipes',
    component: RecipesView,
    meta: { requiresAuth: true },
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: MainView, // Для необработанных маршрутов
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Проверка авторизации
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); // Проверьте токен в localStorage
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/auth'); // Перенаправление на страницу авторизации
  } else {
    next();
  }
});

export default router;
