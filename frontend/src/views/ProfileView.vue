<template>
    <div>
      <ProfileComponent/>
      <ProfileFavouriteComponent/>
      <ProfileHistoryComponent/>
      <SubscriptionComponent/>
    </div>
</template>

<script>

import ProfileComponent from '../components/ProfileComponent.vue';
import ProfileFavouriteComponent from '../components/ProfileFavouriteComponent.vue';
import ProfileHistoryComponent from '../components/ProfileHistoryComponent.vue';
import SubscriptionComponent from '../components/SubscriptionComponent.vue';

export default{
  components:{
    ProfileComponent,
    ProfileFavouriteComponent,
    ProfileHistoryComponent,
    SubscriptionComponent,
  },
};
</script>