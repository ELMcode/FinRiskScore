<template>
  <div>
    <header>
      <div class="wrapper">
        <nav>
          <router-link to="/">Home</router-link>
          <router-link to="/calculate-risk-score">Calculate Risk Score</router-link>
        </nav>
      </div>
    </header>

    <router-view />
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>

</style>
