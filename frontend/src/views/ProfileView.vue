<template>
    <div>
      <ProfileComponent/>
      <SubscriptionComponent/>
    </div>
</template>

<script>

import ProfileComponent from '../components/ProfileComponent.vue';
import SubscriptionComponent from '../components/SubscriptionComponent.vue';

export default{
  components:{
    ProfileComponent,
    SubscriptionComponent,
  },
};
</script>