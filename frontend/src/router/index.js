import { createRouter, createWebHistory } from 'vue-router';

// Импортируйте компоненты для каждой страницы
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

// Определение маршрутов
const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainView,
  },
  {
    path: '/about',
    name: 'About',
    component: AboutView,
  },
  {
    path: '/auth',
    name: 'Auth',
    component: AuthView,
  },
  {
    path: '/favorite',
    name: 'Favorite',
    component: FavoriteView,
  },
  {
    path: '/history',
    name: 'History',
    component: HistoryView,
  },
  {
    path: '/near',
    name: 'Near',
    component: NearView,
  },
  {
    path: '/plan',
    name: 'Plan',
    component: PlanView,
  },
  {
    path: '/plans',
    name: 'Plans',
    component: PlansView,
  },
  {
    path: '/profile',
    name: 'Profile',
    component: ProfileView,
  },
  {
    path: '/recipe',
    name: 'Recipe',
    component: RecipeView,
  },
  {
    path: '/recipes',
    name: 'Recipes',
    component: RecipesView,
  },
  // Добавьте этот маршрут в конец массива маршрутов
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: MainView, // Или другой компонент, который будет отображаться по умолчанию
  },
];

// Создание роутера
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
