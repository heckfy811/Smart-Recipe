<template>
    <div>
      <HeroComponent />
      <AboutComponent/>
      <SubscriptionComponent/>
    </div>
</template>  

<script>
import HeroComponent from "../components/HeroComponent.vue";
import AboutComponent from "../components/AboutComponent.vue";
import SubscriptionComponent from "../components/SubscriptionComponent.vue";

export default{
  components:{
    HeroComponent,
    AboutComponent,
    SubscriptionComponent,
  },
  
};
</script>


