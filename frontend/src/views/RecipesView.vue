<template>
    <div>
    <RecipeList @toggle-sidebar="toggleSidebar" />
    <SidebarComponent :is-visible="isSidebarVisible" @close-sidebar="closeSidebar" />
    <BestRecipeComponent/>
  </div>
</template>  

<script>
import RecipeList from "../components/RecipeList.vue"
import BestRecipeComponent from "../components/BestRecipeComponent.vue";
import SidebarComponent from "../components/SidebarComponent.vue";


export default{
  components:{
    RecipeList,
    BestRecipeComponent,
    SidebarComponent,
  },
  data() {
    return {
      isSidebarVisible: false, // Состояние видимости боковой панели
    };
  },
  methods: {
    toggleSidebar() {
      this.isSidebarVisible = !this.isSidebarVisible;
      const overlay = document.getElementById('overlay');
      overlay.classList.add('active');
      const sidebar = document.getElementById('filter-sidebar');
      sidebar.classList.add('active');
    },
    closeSidebar() {
      this.isSidebarVisible = false;
      const overlay = document.getElementById('overlay');
      overlay.classList.remove('active');

    },
  },
};
</script>