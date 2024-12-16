<template>
    <div>
      <HistoryComponent @toggle-sidebar="toggleSidebar"/>
      <SidebarComponent :is-visible="isSidebarVisible" @close-sidebar="closeSidebar" />
    </div>
</template>

<script>

import HistoryComponent from '../components/HistoryComponent.vue';
import SidebarComponent from '../components/SidebarComponent.vue';
export default{
  components:{
    HistoryComponent,
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