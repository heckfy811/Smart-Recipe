<template>
    <div>
      <ProfileComponent/>
      <ProfileFavouriteComponent/>
      <SubscriptionComponent/>
    </div>
</template>

<script>

import ProfileComponent from '../components/ProfileComponent.vue';
import ProfileFavouriteComponent from '../components/ProfileFavouriteComponent.vue';
import SubscriptionComponent from '../components/SubscriptionComponent.vue';

export default{
  components:{
    ProfileComponent,
    ProfileFavouriteComponent,
    SubscriptionComponent,
  },
};
</script>