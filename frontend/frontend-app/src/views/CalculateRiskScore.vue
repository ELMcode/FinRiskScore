<template>
  <v-container class="wallboard" fluid>
    <v-row>
      <v-col v-for="(data, index) in companyData" :key="data.title" :cols="getCols(data.title)">
        <v-card :color="cardColors[data.title]" class="pa-4 white--text" style="height: 100%;">
          <v-card-title
              class="text-uppercase text-subtitle-2 font-weight-bold"
              v-if="data.title !== 'Image' && data.title !== 'Description'">
            {{ data.title }}
          </v-card-title>
          <v-card-text class="d-flex justify-center text-h6">
            <template v-if="data.title === 'Website'">
              <a :href="data.value" target="_blank">{{ data.value }}</a>
            </template>
            <template v-else-if="data.title === 'Image'">
              <img :src="data.value" alt="Company Image" style="max-width: 30%;" />
            </template>
            <template v-else>
              {{ data.value }}
            </template>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { VCard, VCardTitle, VCardText, VContainer, VRow, VCol } from 'vuetify/components';
import { useRoute } from 'vue-router';
import { onMounted, ref } from "vue";

const companyData = ref([
  { title: 'Company Name', value: '' },
  { title: 'Risk Score', value: 0 },
  { title: 'Industry', value: '' },
  { title: 'Description', value: '' },
  { title: 'Price', value: 0 },
  { title: 'Changes', value: 0 },
  { title: 'Website', value: '' },
  { title: 'Image', value: '' },
  { title: 'Exchange', value: '' },
  { title: 'Is Actively Trading', value: false }
]);

const cardColors = {
  'Company Name': 'blue',
  'Risk Score': 'red',
  'Industry': 'green',
  'Description': 'brown',
  'Price': 'orange',
  'Changes': 'pink',
  'Website': 'teal',
  'Image': 'grey',
  'Exchange': 'cyan',
  'Is Actively Trading': 'amber'
};

const route = useRoute();

onMounted(async () => {
  const symbol = route.query.symbol as string;
  const response = await fetch(`http://localhost:8080/calculate-risk-score?symbol=${symbol}`);
  const data = await response.json();

  console.log(data);

  companyData.value = [
    { title: 'Company Name', value: data.companyName },
    { title: 'Image', value: data.image },
    { title: 'Industry', value: data.industry },
    { title: 'Description', value: data.description },
    { title: 'Price', value: `${data.price} $` },
    { title: 'Changes', value: `${data.changes} %` },
    { title: 'Exchange', value: data.exchange },
    { title: 'Is Actively Trading', value: data.isActivelyTrading },
    { title: 'Website', value: data.website },
    { title: 'Risk Score', value: data.riskScore },
  ];
});

const getCols = (title: string) => {
  return title === 'Description' ? 12 : 4;
};
</script>

<style>
.wallboard {
  padding: 30px;
  background-color:#424242;
}
</style>
