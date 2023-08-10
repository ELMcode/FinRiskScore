<template>
  <div class="risk-score">
    <h1>Risk Score for {{ symbol }}</h1>
    <h2>{{ riskScore }}</h2>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const riskScore = ref(0)
const symbol = ref('')

const route = useRoute()

onMounted(async () => {
  symbol.value = route.query.symbol as string

  const response = await fetch(`http://localhost:8080/calculate-risk-score?symbol=${symbol.value}`)

  const data: { riskScore: number } = await response.json()

  console.log('Données reçues de mon APIrest Go :', data);

  riskScore.value = data.riskScore
})
</script>
