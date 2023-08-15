<template>
  <v-container>
    <v-row>
      <v-col
          v-for="data in companyData"
          :key="data.title"
          cols="12"
          md="5"
      >
        <v-card class="card">
          <v-card-title>{{ data.title }}</v-card-title>
          <v-card-text>{{ data.value }}</v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
// Importer les composants Vuetify
import { VCard, VCardTitle, VCardText, VContainer, VRow, VCol } from 'vuetify/components'
import { useRoute } from 'vue-router'
import {onMounted, ref} from "vue";

const companyData = ref([
  { title: 'Company Name', value: '' },
  { title: 'Risk Score', value: 0 },
  { title: 'Industry', value: '' },
  { title: 'Full Time Employees', value: '' },
  { title: 'Description', value: '' },
  { title: 'Price', value: 0 },
  { title: 'Changes', value: 0 },
  { title: 'Website', value: '' },
  { title: 'Image', value: '' },
  { title: 'Exchange', value: '' },
  { title: 'Is Actively Trading', value: false }
])

const route = useRoute()

onMounted(async () => {
  const symbol = route.query.symbol as string
  const response = await fetch(`http://localhost:8080/calculate-risk-score?symbol=${symbol}`)
  const data = await response.json()

  console.log('Données reçues de mon APIrest Go :', data);

  companyData.value = [
    { title: 'Company Name', value: data.companyName },
    { title: 'Risk Score', value: data.riskScore },
    { title: 'Industry', value: data.industry },
    { title: 'Full Time Employees', value: data.fullTimeEmployees },
    { title: 'Description', value: data.description },
    { title: 'Price', value: `$${data.price}` },
    { title: 'Changes', value: data.changes },
    { title: 'Website', value: data.website },
    { title: 'Image', value: data.image },
    { title: 'Exchange', value: data.exchange },
    { title: 'Is Actively Trading', value: data.isActivelyTrading }
  ]
})
</script>

<style>


</style>
