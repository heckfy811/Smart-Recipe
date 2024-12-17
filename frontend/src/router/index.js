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
import RegisterView from '@/views/RegisterView.vue';
import CheapView from '../views/CheapView.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainView,
    meta: { requiresAuth: true },
  },
  {
    path: '/about',
    name: 'About',
    component: AboutView,
    meta: { requiresAuth: true },
  },
  {
    path: '/auth',
    name: 'Auth',
    component: AuthView,
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterView,
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
    path: '/cheap',
    name: 'Cheap',
    component: CheapView,
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
    component: MainView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// auth-guard
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem("isAuthenticated");

  if (to.meta.requiresAuth && !isAuthenticated) {
    next("/auth");
  } 
  else if ((to.path === "/auth" || to.path === "/register") && isAuthenticated) {
    next("/");
  } 
  else {
    next();
  }
});

export default router;
